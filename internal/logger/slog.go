package logger

import (
	"authjwt/internal/config"
	"io"
	"log/slog"
	"os"
)

const (
	EnvLocal = "local"
	EnvDev   = "dev"
	EnvProd  = "prod"
)

var Logger *slog.Logger

func Init(cfg *config.Config) error {
	logFile, err := os.OpenFile(cfg.AppLogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		return err
	}

	logWriter := io.MultiWriter(os.Stdout, logFile)
	var slogHandler slog.Handler

	switch cfg.AppEnv {
	case EnvLocal:
		slogHandler = slog.NewTextHandler(logWriter, &slog.HandlerOptions{Level: slog.LevelDebug})
	case EnvDev:
		slogHandler = slog.NewJSONHandler(logWriter, &slog.HandlerOptions{Level: slog.LevelDebug})
	case EnvProd:
		slogHandler = slog.NewJSONHandler(logWriter, &slog.HandlerOptions{Level: slog.LevelInfo})
	}

	Logger = slog.New(slogHandler)

	return nil
}
