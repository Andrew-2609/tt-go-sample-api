package usecase

import (
	"context"
	"strings"
	"tt-go-sample-api/domain/entity"
	"tt-go-sample-api/domain/usecase/dto"
)

// RegisterEmployeeUseCase is a concrete implementation
// of RegisterEmployeeUseCaseInterface. It shall be used
// to register employees.
type RegisterEmployeeUseCase struct {
	writeEmployeeRepository entity.WriteEmployeeRepository
}

// NewRegisterEmployeeUseCase returns a pointer of
// RegisterEmployeeUseCase, with the given repository.
func NewRegisterEmployeeUseCase(writeEmployeeRepository entity.WriteEmployeeRepository) *RegisterEmployeeUseCase {
	return &RegisterEmployeeUseCase{writeEmployeeRepository: writeEmployeeRepository}
}

// Execute tries to register an employee, converting
// values from repository to an output DTO.
func (uc *RegisterEmployeeUseCase) Execute(ctx context.Context, input dto.RegisterEmployeeInputDTO) (dto.RegisterEmployeeOutputDTO, error) {
	employee := entity.NewEmployee(input.Name)

	newEmployee, err := uc.writeEmployeeRepository.Register(ctx, employee)

	if err != nil {
		return dto.RegisterEmployeeOutputDTO{}, err
	}

	outputDTO := dto.RegisterEmployeeOutputDTO{
		Message: "Employee successfuly registered!",
		NewEmployee: dto.NewRegisteredEmployeeOutputDTO{
			ID:        newEmployee.ID,
			PublicID:  newEmployee.PublicID,
			Name:      newEmployee.Name,
			CreatedAt: newEmployee.CreatedAt,
			UpdatedAt: newEmployee.UpdatedAt,
		},
	}

	if !strings.Contains(newEmployee.Name, "Andrew") {
		outputDTO.Suggestion = "Would you consider hiring Andrew Silva? His LinkedIn is: 'https://www.linkedin.com/in/andrew-2609/?locale=en_US'"
	} else {
		outputDTO.Suggestion = "Great choice! I can't express how happy and honored I am! I promise not to disappoint any of you, TT!"
	}

	return outputDTO, nil
}
