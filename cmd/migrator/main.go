package main

import (
	"biophilia/internal/config"
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	cfg := config.MustLoad()

	var migrationsTable string
	flag.StringVar(&migrationsTable, "migrationsTable", "schema_migrations", "Name of the migrations table")
	flag.Parse()

	databaseURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDBName,
	)

	if cfg.MigrationsPath == "" {
		log.Fatal("Migrations path is required")
	}

	migrator, err := migrate.New(
		"file://"+cfg.MigrationsPath,
		databaseURL,
	)
	if err != nil {
		log.Fatalf("Failed to initialize migrator: %v", err)
	}

	if err := migrator.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("No new migrations to apply.")
		} else {
			log.Fatalf("Migration failed: %v", err)
		}
	} else {
		fmt.Println("Migrations applied successfully.")
	}
}
