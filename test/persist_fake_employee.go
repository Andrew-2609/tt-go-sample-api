package test

import (
	"database/sql"
	"testing"
	"tt-go-sample-api/domain/entity"
	db "tt-go-sample-api/external/rdb/sqlc"

	"github.com/stretchr/testify/require"
)

// PersistFakeEmployee persists an employee directly
// into the SQL database, and returns an *entiy.Employee.
func PersistFakeEmployee(t *testing.T, sqlDb *sql.DB, employeeName string) *entity.Employee {
	var dbEmployee db.Employee

	err := sqlDb.
		QueryRow(`INSERT INTO employees (name) VALUES ($1) RETURNING id, "publicId", "name", "createdAt", "updatedAt"`, employeeName).
		Scan(
			&dbEmployee.ID,
			&dbEmployee.PublicId,
			&dbEmployee.Name,
			&dbEmployee.CreatedAt,
			&dbEmployee.UpdatedAt,
		)

	require.NoError(t, err)

	return &entity.Employee{
		ID:        dbEmployee.ID,
		PublicID:  dbEmployee.PublicId.String(),
		Name:      dbEmployee.Name,
		CreatedAt: dbEmployee.CreatedAt,
		UpdatedAt: dbEmployee.UpdatedAt,
	}
}
