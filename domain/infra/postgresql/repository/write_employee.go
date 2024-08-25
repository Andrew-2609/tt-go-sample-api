package repository

import (
	"context"
	"tt-go-sample-api/domain/entity"
	db "tt-go-sample-api/external/rdb/sqlc"
)

// WriteEmployeePostgreSQLRepository is a repository used
// to register an employee in a PostgreSQL database.
type WriteEmployeePostgreSQLRepository struct {
}

// NewWriteEmployeePostgreSQLRepository returns a pointer of
// WriteEmployeePostgreSQLRepository.
func NewWriteEmployeePostgreSQLRepository() *WriteEmployeePostgreSQLRepository {
	return &WriteEmployeePostgreSQLRepository{}
}

// Register registers an employee in the PostgreSQL database.
func (r *WriteEmployeePostgreSQLRepository) Register(ctx context.Context, employee *entity.Employee) (*entity.Employee, error) {
	newEmployee, err := db.SQLStoreSingleton.RegisterEmployee(ctx, employee.Name)

	if err != nil {
		return nil, err
	}

	return &entity.Employee{
		ID:        newEmployee.ID,
		PublicID:  newEmployee.PublicId.String(),
		Name:      newEmployee.Name,
		CreatedAt: newEmployee.CreatedAt,
		UpdatedAt: newEmployee.UpdatedAt,
	}, nil
}
