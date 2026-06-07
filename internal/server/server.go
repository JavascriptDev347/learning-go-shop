// Package server for server information
package server

import (
	"net/http"

	"github.com/JavascriptDev347/learning-go-shop/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

// Server struct includes config, db and logger for the server
type Server struct {
	config *config.Config
	db     *gorm.DB
	logger *zerolog.Logger
}

// New func includes cfg, db and logger for the server
func New(cfg *config.Config, db *gorm.DB, logger *zerolog.Logger) *Server {
	return &Server{
		config: cfg,
		db:     db,
		logger: logger,
	}
}

// SetupRoutes func belongs to server and includes router setup such as gin Logger, gin Recovery (for panic) and custom corsMiddleware. It also includes a health check route for testing the server status.
func (s *Server) SetupRoutes() *gin.Engine {
	router := gin.New()

	// add middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(s.corsMiddleware())

	// add routes
	router.GET("/health", s.healthCheck)

	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", s.register)
			auth.POST("/login", s.login)
			auth.POST("/refresh", s.refreshToken)
			auth.POST("/logout", s.logout)
		}
	}
	return router
}

// healthCheck func for check the server is working or no.
func (s *Server) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// corsMiddleware custom middleware for setting cors and some methods and headers for the server. It also handles OPTIONS method for preflight requests.
func (s *Server) corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
