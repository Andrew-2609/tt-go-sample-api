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

func makeListEmployeesUseCaseSut(t *testing.T, ctrl *gomock.Controller) (
	sut *ListEmployeesUseCase,
	getEmployeesRepository *mock.MockGetEmployeesRepository,
) {
	getEmployeesRepository = mock.NewMockGetEmployeesRepository(ctrl)

	sut = NewListEmployeesUseCase(getEmployeesRepository)

	return
}

func TestNewListEmployeesUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sut, _ := makeListEmployeesUseCaseSut(t, ctrl)

	require.NotEmpty(t, sut)
	require.IsType(t, &ListEmployeesUseCase{}, sut)
}

func TestListEmployeesUseCase_Execute(t *testing.T) {
	testCases := map[string]struct {
		setupAndValidate func(t *testing.T, ctrl *gomock.Controller)
	}{
		"entity.GetEmployeesRepository.ListEmployees Fails": {
			setupAndValidate: func(t *testing.T, ctrl *gomock.Controller) {
				sut, getEmployeesRepository := makeListEmployeesUseCaseSut(t, ctrl)

				ctx := context.Background()
				page := int32(1)

				input := dto.ListEmployeesInputDTO{
					Page: &page,
				}

				fakeError := apierr.NewInternalServerError(apierr.CodeUnknownServerError)

				getEmployeesRepository.EXPECT().
					ListEmployees(ctx, page, int32(40)).
					Return(nil, fakeError)

				output, err := sut.Execute(ctx, input)

				require.Exactly(t, dto.ListEmployeesOutputDTO{}, output)
				require.EqualError(t, err, fakeError.Error())
			},
		},
		"Success": {
			setupAndValidate: func(t *testing.T, ctrl *gomock.Controller) {
				sut, getEmployeesRepository := makeListEmployeesUseCaseSut(t, ctrl)

				ctx := context.Background()
				page := int32(1)
				limit := int32(20)

				input := dto.ListEmployeesInputDTO{
					Page:  &page,
					Limit: &limit,
				}

				entitiesFromDB := []*entity.Employee{{
					ID:        21,
					PublicID:  uuid.NewString(),
					Name:      "Andrew Silva",
					CreatedAt: time.Now().Add(-time.Hour),
					UpdatedAt: time.Now().Add(-time.Hour),
				}}

				expectedOutput := dto.ListEmployeesOutputDTO{
					Items: []dto.EmployeeItemOutputDTO{{
						Name:      entitiesFromDB[0].Name,
						PublicID:  entitiesFromDB[0].PublicID,
						CreatedAt: entitiesFromDB[0].CreatedAt,
						UpdatedAt: entitiesFromDB[0].UpdatedAt,
					}},
				}

				getEmployeesRepository.EXPECT().
					ListEmployees(ctx, page, limit).
					Return(entitiesFromDB, nil)

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
