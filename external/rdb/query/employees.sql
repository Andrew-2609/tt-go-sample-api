-- name: ListEmployees :many
SELECT * FROM employees
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: RegisterEmployee :one
INSERT INTO employees ("name")
VALUES ($1)
RETURNING *;