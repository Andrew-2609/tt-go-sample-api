package test

import (
	"context"
	"testing"
	"tt-go-sample-api/external/rdb"
	db "tt-go-sample-api/external/rdb/sqlc"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/stretchr/testify/require"
)

// SetupTestsPostgreSQL sets up the PostgreSQL connection
// and run the database migrations.
func SetupTestsPostgreSQL(t *testing.T, dbName string) {
	dbDriver, err := postgres.WithInstance(db.SQLStoreSingleton.GetDB(), &postgres.Config{})
	require.NoError(t, err)

	mig, err := migrate.NewWithDatabaseInstance("file://../../../../external/rdb/migration", dbName, dbDriver)
	require.NoError(t, err)

	migrationRunner := &rdb.MigrationRunner{Migrator: mig}

	require.NoError(t, migrationRunner.Run(context.Background()))
}
