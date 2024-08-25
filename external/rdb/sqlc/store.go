package db

import (
	"database/sql"
)

// Store defines a database Store, capable of performing
// operations in a database.
type Store interface {
	Querier
	GetDB() *sql.DB
	CloseDB() error
}

// SQLStore is a database relational database Store.
type SQLStore struct {
	*Queries
	db *sql.DB
}

// NewSQLStore returns a pointer to an SQLStore.
func NewSQLStore(db *sql.DB) *SQLStore {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// GetDB returns SQLStore's inner relational
// database handle.
func (s *SQLStore) GetDB() *sql.DB {
	return s.db
}

// CloseDB closes the inner SQLStore's relational
// database connection.
func (s *SQLStore) CloseDB() error {
	return s.db.Close()
}
