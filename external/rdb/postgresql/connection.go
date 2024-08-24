package postgresql

import (
	"database/sql"
	"fmt"
	db "tt-go-sample-api/external/rdb/sqlc"

	_ "github.com/lib/pq"
)

// NewPostgreSQLConnection tries to connect to a
// PostgreSQL database with the given database
// source string.
//
// It returns a db.Store containing the PostgreSQL
// connection as its inner engine.
func NewPostgreSQLConnection(source string) (db.Store, error) {
	conn, err := sql.Open("postgres", source)

	if err != nil {
		return nil, fmt.Errorf("could not establish postgresql connection: %v", err)
	}

	return db.NewSQLStore(conn), nil
}
