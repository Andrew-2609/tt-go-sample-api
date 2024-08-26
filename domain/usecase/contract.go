package usecase

import (
	"context"
	"tt-go-sample-api/domain/usecase/dto"
)

// ListEmployeesUseCaseInterface is responsible for
// retrieving a paginated list of employees.
type ListEmployeesUseCaseInterface interface {
	// Execute tries to retrieve a paginated list of employees
	// based on an input DTO.
	Execute(ctx context.Context, input dto.ListEmployeesInputDTO) (dto.ListEmployeesOutputDTO, error)
}

// RegisterEmployeeUseCaseInterface is responsible for
// registering employees.
type RegisterEmployeeUseCaseInterface interface {
	// Execute tries to register an employee based on an input
	// DTO.
	Execute(ctx context.Context, input dto.RegisterEmployeeInputDTO) (dto.RegisterEmployeeOutputDTO, error)
}

// RequireEmployeeFromHRUseCaseInterface is responsible
// for requiring employees from the HR team.
type RequireEmployeeFromHRUseCaseInterface interface {
	// Execute tries to require an employee based on an input
	// DTO.
	Execute(ctx context.Context, input dto.RequireEmployeeFromHRInputDTO) (dto.RequireEmployeeFromHROutputDTO, error)
}
