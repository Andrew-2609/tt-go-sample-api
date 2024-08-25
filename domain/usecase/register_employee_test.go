package usecase

import (
	"context"
	"testing"
	"time"
	"tt-go-sample-api/domain/apierr"
	"tt-go-sample-api/domain/entity"
	"tt-go-sample-api/domain/usecase/dto"
	"tt-go-sample-api/mock"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func makeRegisterEmployeeUseCaseSut(t *testing.T, ctrl *gomock.Controller) (
	sut *RegisterEmployeeUseCase,
	writeEmployeeRepository *mock.MockWriteEmployeeRepository,
) {
	writeEmployeeRepository = mock.NewMockWriteEmployeeRepository(ctrl)

	sut = NewRegisterEmployeeUseCase(writeEmployeeRepository)

	return
}

func TestNewRegisterEmployeeUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sut, _ := makeRegisterEmployeeUseCaseSut(t, ctrl)

	require.NotEmpty(t, sut)
	require.IsType(t, &RegisterEmployeeUseCase{}, sut)
}

func TestRegisterEmployeeUseCase_Execute(t *testing.T) {
	testCases := map[string]struct {
		setupAndValidate func(t *testing.T, ctrl *gomock.Controller)
	}{
		"entity.WriteEmployeeRepository.Register Fails": {
			setupAndValidate: func(t *testing.T, ctrl *gomock.Controller) {
				sut, writeEmployeeRepository := makeRegisterEmployeeUseCaseSut(t, ctrl)

				ctx := context.Background()

				input := dto.RegisterEmployeeInputDTO{
					Name: "andrew",
				}

				fakeError := apierr.NewInternalServerError(apierr.CodeUnknownServerError)

				employee := entity.NewEmployee(input.Name)

				writeEmployeeRepository.EXPECT().
					Register(ctx, employee).
					Return(nil, fakeError)

				output, err := sut.Execute(ctx, input)

				require.Exactly(t, dto.RegisterEmployeeOutputDTO{}, output)
				require.EqualError(t, err, fakeError.Error())
			},
		},
		"Success Without Andrew": {
			setupAndValidate: func(t *testing.T, ctrl *gomock.Controller) {
				sut, writeEmployeeRepository := makeRegisterEmployeeUseCaseSut(t, ctrl)

				ctx := context.Background()

				input := dto.RegisterEmployeeInputDTO{
					Name: "John Stockton",
				}

				employee := entity.NewEmployee(input.Name)

				newEmployee := &entity.Employee{
					ID:        21,
					PublicID:  uuid.NewString(),
					Name:      employee.Name,
					CreatedAt: time.Now().Add(-time.Hour),
					UpdatedAt: time.Now().Add(-time.Hour),
				}

				expectedOutput := dto.RegisterEmployeeOutputDTO{
					Message: "Employee successfuly registered!",
					NewEmployee: dto.NewRegisteredEmployeeOutputDTO{
						ID:        newEmployee.ID,
						PublicID:  newEmployee.PublicID,
						Name:      newEmployee.Name,
						CreatedAt: newEmployee.CreatedAt,
						UpdatedAt: newEmployee.UpdatedAt,
					},
					Suggestion: "Great choice! But would you consider hiring Andrew Silva? He really wants to join TT, and you can find him at 'https://www.linkedin.com/in/andrew-2609/?locale=en_US'",
				}

				writeEmployeeRepository.EXPECT().
					Register(ctx, employee).
					Return(newEmployee, nil)

				output, err := sut.Execute(ctx, input)

				require.NoError(t, err)
				require.Exactly(t, expectedOutput, output)
			},
		},
		"Success With Andrew": {
			setupAndValidate: func(t *testing.T, ctrl *gomock.Controller) {
				sut, writeEmployeeRepository := makeRegisterEmployeeUseCaseSut(t, ctrl)

				ctx := context.Background()

				input := dto.RegisterEmployeeInputDTO{
					Name: "Andrew Silva",
				}

				employee := entity.NewEmployee(input.Name)

				newEmployee := &entity.Employee{
					ID:        21,
					PublicID:  uuid.NewString(),
					Name:      employee.Name,
					CreatedAt: time.Now().Add(-time.Hour),
					UpdatedAt: time.Now().Add(-time.Hour),
				}

				expectedOutput := dto.RegisterEmployeeOutputDTO{
					Message: "Employee successfuly registered!",
					NewEmployee: dto.NewRegisteredEmployeeOutputDTO{
						ID:        newEmployee.ID,
						PublicID:  newEmployee.PublicID,
						Name:      newEmployee.Name,
						CreatedAt: newEmployee.CreatedAt,
						UpdatedAt: newEmployee.UpdatedAt,
					},
					Suggestion: "Great choice! I can't express how happy and honored I am! I promise not to disappoint any of you, TT!",
				}

				writeEmployeeRepository.EXPECT().
					Register(ctx, employee).
					Return(newEmployee, nil)

				output, err := sut.Execute(ctx, input)

				require.NoError(t, err)
				require.Exactly(t, expectedOutput, output)
			},
		},
	}

	for scenario, testCase := range testCases {
		t.Run(scenario, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			testCase.setupAndValidate(t, ctrl)
		})
	}
}
