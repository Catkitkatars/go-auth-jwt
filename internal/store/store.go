package store

import (
	"authjwt/internal/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB
var Migrate *migrate.Migrate

func InitDB() error {
	dsn := createDsn(config.Cfg)

	db, err := sql.Open(config.Cfg.DBConnection, dsn)
	if err != nil {
		return fmt.Errorf("sql.Open failed: %w", err)
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("db.Ping failed: %w", err)
	}

	DB = db

	Migrate, err = migrate.New(
		"file://migrations",
		dsn,
	)
	if err != nil {
		return fmt.Errorf("failed to init migrate: %w", err)
	}

	return nil
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
