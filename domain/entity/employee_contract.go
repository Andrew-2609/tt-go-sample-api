package entity

import "context"

// GetEmployeesRepository defines the base methods
// that any repository built to retrieve employees
// from a data store should use.
type GetEmployeesRepository interface {
	// ListEmployees list employees from the data store
	// based on a page and a limit, used for pagination.
	ListEmployees(ctx context.Context, page, limit int32) ([]*Employee, error)
}

// WriteEmployeeRepository defines the base methods
// that any repository built to write or update employees
// to/in a data store should use.
type WriteEmployeeRepository interface {
	// Register persists an employee in the data store.
	Register(ctx context.Context, employee *Employee) (*Employee, error)
}
