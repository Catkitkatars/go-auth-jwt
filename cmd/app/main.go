package main

import (
	"authjwt/internal/config"
	"authjwt/internal/http"
	logs "authjwt/internal/logger"
	"authjwt/internal/store"
	"log"
	"log/slog"
	"os"
)

func main() {
	cfgErr := config.Init()

	if cfgErr != nil {
		log.Fatalf("Error init cfg. Err: %v", cfgErr)
		os.Exit(1)
	}

	logErr := logs.Init(config.Cfg)

	if logErr != nil {
		log.Fatalf("Error init logger. Err: %v", logErr)
		os.Exit(1)
	}

	logs.Logger.Info("starting auth-jwt", slog.String("env", config.Cfg.AppEnv))

	errDb := store.InitDB()

	if errDb != nil {
		log.Fatalf("store.InitDB: %v", errDb)
		os.Exit(1)
	}

	errMigrate := store.Migrate.Up()

	if errMigrate != nil {
		if errMigrate.Error() != "no change" {
			log.Fatalf("store.Migrate.Up: %v", errMigrate)
			os.Exit(1)
		}
	}

	srvErr := http.ServerStart()

	if srvErr != nil {
		log.Fatalf("http.ServerStart: %v", srvErr)
		os.Exit(1)
	}
}
