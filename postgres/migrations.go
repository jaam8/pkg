package postgres

import (
	"context"
	"errors"
	"log"

	"github.com/pressly/goose/v3"
)

func Migrate(ctx context.Context, config Config, migrationsPath string) error {
	dsn := config.GetDsn()

	db, err := goose.OpenDBWithDriver("pgx", dsn)
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	defer db.Close()

	if err = goose.SetDialect("postgres"); err != nil {
		log.Fatalf("failed to set dialect: %v", err)
	}

	if err = goose.Up(db, migrationsPath); err != nil && !errors.Is(err, goose.ErrNoMigrations) {
		log.Fatalf("failed to run migrations: %v", err)
	}
	log.Println("migrated successfully")
	return nil
}
