package config

import (
	"fmt"
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	AppName    string `env:"APP_NAME"`
	AppEnv     string `env:"APP_ENV"`
	AppHost    string `env:"APP_HOST"`
	AppPort    int    `env:"APP_PORT"`
	AppLogPath string `env:"APP_LOG_PATH"`

	DBConnection string `env:"DB_CONNECTION"`
	DBHost       string `env:"DB_HOST"`
	DBPort       int    `env:"DB_PORT"`
	DBDatabase   string `env:"DB_DATABASE"`
	DBUsername   string `env:"DB_USERNAME"`
	DBPassword   string `env:"DB_PASSWORD"`
}

var Cfg *Config

func Init() error {
	err := godotenv.Load(".env")

	if err != nil {
		return fmt.Errorf("godotenv.Load: %w", err)
	}

	errCfg := env.Parse(&Cfg)

	if errCfg != nil {
		return fmt.Errorf("env.Parse: %w", errCfg)
	}
}
