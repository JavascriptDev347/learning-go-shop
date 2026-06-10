// Package main
package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/JavascriptDev347/learning-go-shop/internal/config"
	"github.com/JavascriptDev347/learning-go-shop/internal/database"
	"github.com/JavascriptDev347/learning-go-shop/internal/providers"
	"github.com/JavascriptDev347/learning-go-shop/internal/server"
	"github.com/JavascriptDev347/learning-go-shop/internal/services"
	"github.com/gin-gonic/gin"

	"github.com/JavascriptDev347/learning-go-shop/internal/logger"
)

func main() {

	// Initialize logger
	log := logger.New()
	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load configuration")
	}

	// connect to db
	db, err := database.New(&cfg.Database)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	mainDb, err := db.DB()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get database")
	}
	defer func(mainDb *sql.DB) {
		err := mainDb.Close()
		if err != nil {
			log.Error().Err(err).Msg("Failed to close database connection")
		}
	}(mainDb)

	gin.SetMode(cfg.Server.GinMode)

	// services
	authService := services.NewAuthService(db, cfg)
	productService := services.NewProductService(db)
	userService := services.NewUserService(db)
	uploadService := services.NewUploadService(providers.NewLocalUploadProvider(cfg.Upload.Path))

	srv := server.New(cfg, db, &log, authService, productService, userService, uploadService)
	router := srv.SetupRoutes()

	httpServer := http.Server{
		Addr:        fmt.Sprintf(":%s", cfg.Server.Port),
		Handler:     router,
		ReadTimeout: 10 * time.Second,
	}

	go func() {
		log.Info().Str("port", cfg.Server.Port).Msg("starting http server")
		if err := httpServer.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("Failed to start server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("failed to shutdown http server")
		return
	}

	log.Info().Msg("Shutting down database")
}
