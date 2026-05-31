// Package main
package main

import (
	"database/sql"

	"github.com/JavascriptDev347/learning-go-shop/internal/config"
	"github.com/JavascriptDev347/learning-go-shop/internal/database"
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

	log.Info().Msg("Starting server")
}
