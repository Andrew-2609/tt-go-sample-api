package usecase

import (
	"context"
	"tt-go-sample-api/domain/entity"
	"tt-go-sample-api/domain/usecase/dto"
)

// ListEmployeesUseCase is a concrete implementation
// of ListEmployeesUseCaseInterface. It shall be
// used to search for a paginated list of employees.
type ListEmployeesUseCase struct {
	getEmployeesRepository entity.GetEmployeesRepository
}

// NewListEmployeesUseCase returns a pointer to
// ListEmployeesUseCase, with the given repository.
func NewListEmployeesUseCase(getEmployeesRepository entity.GetEmployeesRepository) *ListEmployeesUseCase {
	return &ListEmployeesUseCase{getEmployeesRepository: getEmployeesRepository}
}

// Execute tries to retrieve a paginated list of
// employees, converting the values from repository
// to an output DTO.
func (uc *ListEmployeesUseCase) Execute(ctx context.Context, input dto.ListEmployeesInputDTO) (dto.ListEmployeesOutputDTO, error) {
	var page, limit int32

	page = *input.Page

	if input.Limit == nil {
		limit = 40
	} else {
		limit = *input.Limit
	}

	employees, err := uc.getEmployeesRepository.ListEmployees(ctx, page, limit)

	if err != nil {
		return dto.ListEmployeesOutputDTO{}, err
	}

	items := make([]dto.EmployeeItemOutputDTO, 0, len(employees))

	for _, employee := range employees {
		items = append(items, dto.EmployeeItemOutputDTO{
			PublicID:  employee.PublicID,
			Name:      employee.Name,
			CreatedAt: employee.CreatedAt,
			UpdatedAt: employee.UpdatedAt,
		})
	}

	return dto.ListEmployeesOutputDTO{
		Items: items,
	}, nil
}
