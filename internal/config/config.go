// Package config manages application configuration states and environments.
package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// Config holds all configurations parsed for the application lifespan.
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	AWS      AWSConfig
	Upload   UploadConfig
}

// ServerConfig for server settings, including port and Gin mode.
type ServerConfig struct {
	Port    string
	GinMode string
}

// DatabaseConfig for db settings, including host,port,user,password name and sslmode
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

// JWTConfig for jwt, including secret, expires in and refresh token expires in
type JWTConfig struct {
	Secret                string
	ExpiresIn             time.Duration
	RefreshTokenExpiresIn time.Duration
}

// AWSConfig for aws, including accessKeyId, secretAccessKey, region, s3Bucket and s3Endpoint
type AWSConfig struct {
	AccessKeyID     string
	SecretAccessKey string
	Region          string
	S3Bucket        string
	S3Endpoint      string
}

// UploadConfig for file upload settings, including path and max file size
type UploadConfig struct {
	Path           string
	MaxFileSize    int64
	UploadProvider string
}

// Load initializes and returns the application configuration.
func Load() (*Config, error) {
	_ = godotenv.Load()

	jwtExpiresIn, _ := time.ParseDuration(getEnv("JWT_EXPIRES_IN", "15m"))
	refreshTokenExpiresIn, _ := time.ParseDuration(getEnv("REFRESH_TOKEN_EXPIRES_IN", "168h"))
	maxUploadSize, _ := strconv.ParseInt(getEnv("MAX_UPLOAD_SIZE", "10485760"), 10, 64) // 10 MB
	return &Config{
		Server: ServerConfig{
			Port:    getEnv("PORT", "8080"),
			GinMode: getEnv("GIN_MODE", "debug"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DATABASE_HOST", "localhost"),
			Port:     getEnv("DATABASE_PORT", "5432"),
			User:     getEnv("DATABASE_USER", "postgres"),
			Password: getEnv("DATABASE_PASSWORD", "postgres"),
			Name:     getEnv("DATABASE_NAME", "go_ecommerce_shop"),
			SSLMode:  getEnv("DATABASE_SSLMODE", "disable"),
		},
		JWT: JWTConfig{
			Secret:                getEnv("JWT_SECRET", "secret"),
			ExpiresIn:             jwtExpiresIn,
			RefreshTokenExpiresIn: refreshTokenExpiresIn,
		},
		AWS: AWSConfig{
			AccessKeyID:     getEnv("AWS_ACCESS_KEY_ID", "test"),
			SecretAccessKey: getEnv("AWS_SECRET_ACCESS_KEY", "test"),
			Region:          getEnv("AWS_REGION", "us-east-1"),
			S3Bucket:        getEnv("AWS_S3_BUCKET", "ecommerce-uploads"),
			S3Endpoint:      getEnv("AWS_S3_ENDPOINT", "http://localhost:9000"),
		},
		Upload: UploadConfig{
			Path:           getEnv("UPLOAD_PATH", "./uploads"),
			MaxFileSize:    maxUploadSize,
			UploadProvider: getEnv("UPLOAD_PROVIDER", "local"),
		},
	}, nil
}

// getEnv retrieves environment variables with a fallback default value.
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
