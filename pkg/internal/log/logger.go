package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/silaselisha/go-daraja/pkg/internal/config"
)

func DarajaLogger(cfg *config.Configs) zerolog.Logger {
	env := cfg.MpesaEnvironment

	switch env {
	case "sandbox":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "production":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	}

	return zerolog.New(os.Stdout).With().Timestamp().Logger()
}
