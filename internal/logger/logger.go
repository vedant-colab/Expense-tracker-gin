package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Init(env string) {
	// Timestamp format
	zerolog.TimeFieldFormat = time.RFC3339

	if env == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		})
	} else {
		log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	}

	log.Info().Msg("Logger initialized")
}

func L() *zerolog.Logger {
	return &log.Logger
}
