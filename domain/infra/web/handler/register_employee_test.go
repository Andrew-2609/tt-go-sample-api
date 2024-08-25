package handler_test

import (
	"encoding/json"
	"net/http"
	"testing"
	"tt-go-sample-api/domain/apierr"
	"tt-go-sample-api/domain/usecase/dto"
	db "tt-go-sample-api/external/rdb/sqlc"
	"tt-go-sample-api/test"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type RegisterEmployeeWebHandlerTestSuite struct {
	suite.Suite
	Method   string
	BasePath string
}

func (suite *RegisterEmployeeWebHandlerTestSuite) SetupSuite() {
	suite.Method = http.MethodPost
	suite.BasePath = test.BuildURL("employees")

	test.SetupTestsPostgreSQL(suite.T(), apiConfigTestSingleton.GetPostgreSQLSource(), apiConfigTestSingleton.DBName)
	test.TruncateTables(suite.T(), db.SQLStoreSingleton.GetDB())
}

func (suite *RegisterEmployeeWebHandlerTestSuite) TearDownSubTest() {
	test.TruncateTables(suite.T(), db.SQLStoreSingleton.GetDB())
}

func (suite *RegisterEmployeeWebHandlerTestSuite) TearDownSuite() {
	require.NoError(suite.T(), db.SQLStoreSingleton.CloseDB())
}

func (suite *RegisterEmployeeWebHandlerTestSuite) TestHandle() {
	testCases := map[string]struct {
		inputDTO dto.RegisterEmployeeInputDTO
		validate func(t *testing.T, resp *http.Response)
	}{
		"400 - Bad Request": {
			inputDTO: dto.RegisterEmployeeInputDTO{
				Name: "A",
			},
			validate: func(t *testing.T, resp *http.Response) {
				require.Equal(t, http.StatusBadRequest, resp.StatusCode)

				var apiErr struct {
					Error apierr.APIError `json:"error"`
				}

				err := json.NewDecoder(resp.Body).Decode(&apiErr)
				require.NoError(t, err)

				requestValidationErr := apierr.NewRequestValidationError()

				require.Equal(t, requestValidationErr.Code, apiErr.Error.Code)
				require.Equal(t, requestValidationErr.Message, apiErr.Error.Message)
			},
		},
		"201 - Success Without Andrew": {
			inputDTO: dto.RegisterEmployeeInputDTO{
				Name: "kobe Bryant",
			},
			validate: func(t *testing.T, resp *http.Response) {
				require.Equal(t, http.StatusCreated, resp.StatusCode)

				var outputDTO dto.RegisterEmployeeOutputDTO

				err := json.NewDecoder(resp.Body).Decode(&outputDTO)
				require.NoError(t, err)

				require.Equal(t, "Employee successfuly registered!", outputDTO.Message)
				require.True(t, outputDTO.NewEmployee.ID > 0)
				require.NoError(t, uuid.Validate(outputDTO.NewEmployee.PublicID))
				require.Equal(t, "Kobe Bryant", outputDTO.NewEmployee.Name)
				require.NotZero(t, outputDTO.NewEmployee.CreatedAt)
				require.NotZero(t, outputDTO.NewEmployee.UpdatedAt)
				require.Equal(t, "Great choice! But would you consider hiring Andrew Silva? He really wants to join TT, and you can find him at 'https://www.linkedin.com/in/andrew-2609/?locale=en_US'", outputDTO.Suggestion)
			},
		},
		"201 - Success With Andrew": {
			inputDTO: dto.RegisterEmployeeInputDTO{
				Name: "Andrew Silva",
			},
			validate: func(t *testing.T, resp *http.Response) {
				require.Equal(t, http.StatusCreated, resp.StatusCode)

				var outputDTO dto.RegisterEmployeeOutputDTO

				err := json.NewDecoder(resp.Body).Decode(&outputDTO)
				require.NoError(t, err)

				require.Equal(t, "Employee successfuly registered!", outputDTO.Message)
				require.True(t, outputDTO.NewEmployee.ID > 0)
				require.NoError(t, uuid.Validate(outputDTO.NewEmployee.PublicID))
				require.Equal(t, "Andrew Silva", outputDTO.NewEmployee.Name)
				require.NotZero(t, outputDTO.NewEmployee.CreatedAt)
				require.NotZero(t, outputDTO.NewEmployee.UpdatedAt)
				require.Equal(t, "Great choice! I can't express how happy and honored I am! I promise not to disappoint any of you, TT!", outputDTO.Suggestion)
			},
		},
	}

	for scenario, testCase := range testCases {
		suite.Run(scenario, func() {
			body := test.ParseRequestBody(suite.T(), testCase.inputDTO)

			req, err := test.NewRequest(suite.Method, suite.BasePath, body)
			require.NoError(suite.T(), err)

			resp, err := apiApp.Test(req)
			require.NoError(suite.T(), err)

			testCase.validate(suite.T(), resp)
		})
	}
}

func TestRegisterEmployeeWebHandlerTestSuiteIntegration(t *testing.T) {
	test.SkipTestIfShortFlagWasPassed(t)
	suite.Run(t, new(RegisterEmployeeWebHandlerTestSuite))
}
