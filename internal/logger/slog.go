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

func Init(cfg *config.Config) (*slog.Logger, error) {
	logFile, err := os.OpenFile(cfg.AppLogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		return nil, err
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

	return slog.New(slogHandler), nil
}
