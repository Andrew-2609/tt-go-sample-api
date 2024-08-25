// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: employees.sql

package db

import (
	"context"
)

const listEmployees = `-- name: ListEmployees :many
SELECT id, "publicId", name, "createdAt", "updatedAt" FROM employees
ORDER BY id
LIMIT $1 OFFSET $2
`

type ListEmployeesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListEmployees(ctx context.Context, arg ListEmployeesParams) ([]Employee, error) {
	rows, err := q.db.QueryContext(ctx, listEmployees, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Employee{}
	for rows.Next() {
		var i Employee
		if err := rows.Scan(
			&i.ID,
			&i.PublicId,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const registerEmployee = `-- name: RegisterEmployee :one
INSERT INTO employees ("name")
VALUES ($1)
RETURNING id, "publicId", name, "createdAt", "updatedAt"
`

func (q *Queries) RegisterEmployee(ctx context.Context, name string) (Employee, error) {
	row := q.db.QueryRowContext(ctx, registerEmployee, name)
	var i Employee
	err := row.Scan(
		&i.ID,
		&i.PublicId,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
