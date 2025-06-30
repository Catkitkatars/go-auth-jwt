package store

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func InitDB(dsn string, dbType string) error {
	db, err := sql.Open(dbType, dsn)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	DB = db
	return nil
}

func RunMigrate(dsn string) error {
	m, err := migrate.New(
		"file://migrations",
		dsn,
	)

	if err != nil {
		return fmt.Errorf("error creating migrate instance: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("error applying migrations: %w", err)
	}

	return nil
}
