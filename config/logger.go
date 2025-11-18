package config

import (
	"os"

	"github.com/rs/zerolog"
)

func NewLogger() zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	return zerolog.New(os.Stderr).With().Str("service", "post management").Timestamp().Logger()
}
