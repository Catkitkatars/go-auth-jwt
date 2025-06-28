package main

import (
	"authjwt/internal/config"
	logs "authjwt/internal/logger"
	"log"
	"log/slog"
	"os"
)

func main() {
	cfg, cfgErr := config.Init()

	if cfgErr != nil {
		log.Fatalf("Error init cfg. Err: %e", cfgErr)
		os.Exit(1)
	}

	logger, logErr := logs.Init(&cfg)

	if logErr != nil {
		log.Fatalf("Error init logger. Err: %e", logErr)
		os.Exit(1)
	}

	logger.Info("starting auth-jwt", slog.String("env", cfg.AppEnv))

	// Router - github.com/gin-gonic/gin

	// Orm - gorm.io/gorm - router and orm too big for my project, but I want to try it

	// Middleware - github.com/go-chi/jwtauth
}
