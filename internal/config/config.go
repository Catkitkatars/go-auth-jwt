package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"log"
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

func Init() (Config, error) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Ошибка при загрузке .env файла: %v", err)
		return Config{}, err
	}

	var cfg Config
	errCfg := env.Parse(&cfg)

	if errCfg != nil {
		log.Fatalf("Не удалось проинициализировать .env, err: %e", errCfg)
		return cfg, errCfg
	}

	return cfg, nil
}
