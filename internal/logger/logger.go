// Package logger provides structured logging functionalities using zerolog.
package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// New func for logger initialize
func New() zerolog.Logger {
	zerolog.TimeFieldFormat = time.RFC3339

	if os.Getenv("GIN_MODE") != "release" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	}

	return log.Logger
}
