-- name: ListEmployees :many
SELECT * FROM employees
ORDER BY id
LIMIT $1 OFFSET $2;