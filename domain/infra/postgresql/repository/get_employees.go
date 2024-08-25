package repository

import (
	"context"
	"tt-go-sample-api/domain/apierr"
	"tt-go-sample-api/domain/entity"
	db "tt-go-sample-api/external/rdb/sqlc"
)

// GetEmployeesPostgreSQLRepository is a repository
// used to retrieve many employees from the database.
//
// It implements the domain.GetEmployeesRepository
// interface, and use PostgreSQL implementation
// underneath.
type GetEmployeesPostgreSQLRepository struct {
}

// NewGetEmployeesPostgreSQLRepository returns a pointer
// to GetEmployeesPostgreSQLRepository.
func NewGetEmployeesPostgreSQLRepository() *GetEmployeesPostgreSQLRepository {
	return &GetEmployeesPostgreSQLRepository{}
}

// ListEmployees returns a list of employees from the
// PostgreSQL database, paginated with page and limit
// parameters.
func (r *GetEmployeesPostgreSQLRepository) ListEmployees(ctx context.Context, page, limit int32) ([]*entity.Employee, error) {
	employees, err := db.SQLStoreSingleton.ListEmployees(ctx, db.ListEmployeesParams{
		Limit:  limit,
		Offset: (page - 1) * limit,
	})

	if err != nil {
		return nil, apierr.NewInternalServerError(apierr.CodeSQLListEmployeesFailedErrorCode)
	}

	employeeEntities := make([]*entity.Employee, 0, len(employees))

	for _, employee := range employees {
		employeeEntities = append(employeeEntities, &entity.Employee{
			ID:        employee.ID,
			PublicID:  employee.PublicId.String(),
			Name:      employee.Name,
			CreatedAt: employee.CreatedAt,
			UpdatedAt: employee.UpdatedAt,
		})
	}

	return employeeEntities, nil
}
