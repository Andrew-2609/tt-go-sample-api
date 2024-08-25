package test

import (
	"database/sql"
	"fmt"
	"sync"
	"testing"
	"tt-go-sample-api/util"

	"github.com/stretchr/testify/require"
)

// TruncateTables truncates all database tables.
//
// MUST RUN ONLY DURING TESTS!!
func TruncateTables(t *testing.T, db *sql.DB) {
	if !util.IsTestEnv() {
		panic("truncate tables function CANNOT run outside tests environment")
	}

	rows, err := db.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public' AND table_name <> 'schema_migrations'")
	require.NoError(t, err)

	var tables []string

	for rows.Next() {
		var table string

		require.NoError(t, rows.Scan(&table))

		tables = append(tables, table)
	}

	require.NoError(t, rows.Err())

	var wg sync.WaitGroup

	for _, table := range tables {
		wg.Add(1)

		go func(table string) {
			defer wg.Done()

			_, err := db.Exec(fmt.Sprintf("TRUNCATE TABLE \"%s\" CASCADE", table))
			require.NoError(t, err)
		}(table)
	}

	wg.Wait()
}
