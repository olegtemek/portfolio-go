package logger

import (
	"log/slog"
	"os"

	"github.com/olegtemek/portfolio-go/internal/config"
)

func New(cfg *config.Config) *slog.Logger {

	var opts slog.HandlerOptions
	var handler slog.Handler

	switch cfg.LoggerLevel {
	case "DEBUG":
		opts.Level = slog.LevelDebug
	case "INFO":
		opts.Level = slog.LevelInfo
	case "WARN":
		opts.Level = slog.LevelWarn
	case "ERROR":
		opts.Level = slog.LevelWarn
	default:
		opts.Level = slog.LevelDebug
	}

	handler = slog.NewJSONHandler(os.Stdout, &opts)

	logger := slog.New(handler)

	logger.With(slog.String("Logger level:", cfg.LoggerLevel))

	return logger
}
