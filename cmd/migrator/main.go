package main

import (
	"biophilia/internal/infrastructure/config"
	"errors"
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

func main() {
	cfg := config.MustLoad()

	var steps int
	var direction string
	flag.IntVar(&steps, "steps", 0, "Number of migration steps (0 to ignore)")
	flag.StringVar(&direction, "direction", "up", "Migration direction: 'up' or 'down'")
	flag.Parse()

	if cfg.MigrationsPath == "" {
		log.Fatal("Migrations path is required")
	}

	migrator, err := migrate.New(
		"file://"+cfg.MigrationsPath,
		cfg.DataBaseDSN(),
	)
	if err != nil {
		log.Fatalf("Failed to initialize migrator: %v", err)
	}
	defer closeMigrator(migrator)

	if steps != 0 {
		fmt.Printf("Applying %d migration steps in %s direction...\n", steps, direction)
		if err := applySteps(migrator, steps, direction); err != nil {
			log.Fatalf("Migration step failed: %v", err)
		}
		return
	}

	switch direction {
	case "up":
		fmt.Println("Applying migrations up...")
		err = migrator.Up()
	case "down":
		fmt.Println("Applying migrations down...")
		err = migrator.Down()
	default:
		log.Fatalf("Invalid migration direction: %s. Use 'up' or 'down'.", direction)
	}

	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("No new migrations to apply.")
		} else {
			log.Fatalf("Migration failed: %v", err)
		}
	} else {
		fmt.Println("Migrations applied successfully.")
	}
}

func applySteps(migrator *migrate.Migrate, steps int, direction string) error {
	switch direction {
	case "up":
		return migrator.Steps(steps)
	case "down":
		return migrator.Steps(-steps)
	default:
		return fmt.Errorf("invalid migration direction: %s", direction)
	}
}

func closeMigrator(migrator *migrate.Migrate) {
	srcErr, dbErr := migrator.Close()
	if srcErr != nil || dbErr != nil {
		log.Printf("Error closing migrator: source error: %v, database error: %v", srcErr, dbErr)
	}
}
