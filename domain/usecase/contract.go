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
