package main

import (
	"authjwt/internal/config"
	logs "authjwt/internal/logger"
	"authjwt/internal/store"
	"fmt"
	"log"
	"log/slog"
	"os"
)

func main() {
	cfgErr := config.Init()

	if cfgErr != nil {
		log.Fatalf("Error init cfg. Err: %e", cfgErr)
		os.Exit(1)
	}

	logErr := logs.Init(config.Cfg)

	if logErr != nil {
		log.Fatalf("Error init logger. Err: %e", logErr)
		os.Exit(1)
	}

	logs.Logger.Info("starting auth-jwt", slog.String("env", config.Cfg.AppEnv))

	dsn := createDsn(config.Cfg)

	errDb := store.InitDB(dsn, config.Cfg.DBConnection)

	if errDb != nil {
		log.Fatalf("store.InitDB: %e", errDb)
		os.Exit(1)
	}

	errMigrate := store.RunMigrate(dsn)

	if errMigrate != nil {
		log.Fatalf("store.RunMigrate: %e", errMigrate)
		os.Exit(1)
	}

	// Router - github.com/julienschmidt/httprouter

	// Orm - gorm.io/gorm - router and orm too big for my project, but I want to try it

	// Middleware - github.com/go-chi/jwtauth
}

func createDsn(cfg *config.Config) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.DBUsername,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBDatabase,
	)
}
