package rdb

import (
	"context"
	"database/sql"
	"fmt"
	"tt-go-sample-api/pkg/logger"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Migrator is the inner engine of the
// MigrationRunner wrapper. It must be
// able to run relational database migrations.
type Migrator interface {
	// Up shall run all _up_ migrations.
	Up() error
}

// MigrationRunner is a wrapper to a migration
// runner engine (e.g. migrate).
type MigrationRunner struct {
	Migrator Migrator
}

// NewMigrationRunner returns a pointer of MigrationRunner.
func NewMigrationRunner() *MigrationRunner {
	return &MigrationRunner{}
}

// WithMigrate configures the MigrationRunner wrapper
// to run its migrations with the `migrate` engine.
func (r *MigrationRunner) WithMigrate(driver, databaseName string, db *sql.DB) (*MigrationRunner, error) {
	var dbDriver database.Driver
	var err error

	switch driver {
	case "postgres":
		{
			dbDriver, err = postgres.WithInstance(db, &postgres.Config{})

			if err != nil {
				return r, fmt.Errorf("coult not initialize migrate with postgres driver: %v", err)
			}

			break
		}
	default:
		return r, fmt.Errorf("unsupported driver given to migrate: '%s'", driver)
	}

	mig, err := migrate.NewWithDatabaseInstance("file://external/rdb/migration", databaseName, dbDriver)

	if err != nil {
		return r, fmt.Errorf("could not setup migrations with migrate: %v", err)
	}

	r.Migrator = mig

	return r, nil
}

// Run is responsible for running the database migrations.
func (r *MigrationRunner) Run(ctx context.Context) error {
	logger.APILoggerSingleton.Info(ctx, logger.LogInput{
		Message: "Running migrations...",
	})

	if err := r.Migrator.Up(); err != nil {
		if err.Error() != "no change" {
			return fmt.Errorf("could not run migrations: %v", err)
		}

		logger.APILoggerSingleton.Info(ctx, logger.LogInput{
			Message: "No changes detected",
		})
	} else {
		logger.APILoggerSingleton.Info(ctx, logger.LogInput{
			Message: "Migrations ran successfully!",
		})
	}

	return nil
}
