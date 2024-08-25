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

type ListEmployeesWebHandlerTestSuite struct {
	suite.Suite
	Method   string
	BasePath string
}

func (suite *ListEmployeesWebHandlerTestSuite) SetupSuite() {
	suite.Method = http.MethodGet
	suite.BasePath = test.BuildURL("employees")

	test.SetupTestsPostgreSQL(suite.T(), apiConfigTestSingleton.GetPostgreSQLSource(), apiConfigTestSingleton.DBName)
	test.TruncateTables(suite.T(), db.SQLStoreSingleton.GetDB())
}

func (suite *ListEmployeesWebHandlerTestSuite) TearDownSubTest() {
	test.TruncateTables(suite.T(), db.SQLStoreSingleton.GetDB())
}

func (suite *ListEmployeesWebHandlerTestSuite) TearDownSuite() {
	require.NoError(suite.T(), db.SQLStoreSingleton.CloseDB())
}

func (suite *ListEmployeesWebHandlerTestSuite) TestHandle() {
	testCases := map[string]struct {
		page     int32
		limit    int32
		setup    func(t *testing.T)
		validate func(t *testing.T, resp *http.Response)
	}{
		"400 - Bad Request": {
			page:  -3,
			setup: func(t *testing.T) {},
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
		"200 - Success Without Items": {
			page:  1,
			setup: func(t *testing.T) {},
			validate: func(t *testing.T, resp *http.Response) {
				require.Equal(t, http.StatusOK, resp.StatusCode)

				var outputDTO dto.ListEmployeesOutputDTO

				err := json.NewDecoder(resp.Body).Decode(&outputDTO)
				require.NoError(t, err)

				require.Len(t, outputDTO.Items, 0)
			},
		},
		"200 - Success With Items": {
			page:  1,
			limit: 4,
			setup: func(t *testing.T) {
				test.PersistFakeEmployee(t, db.SQLStoreSingleton.GetDB(), "Andrew Silva")
			},
			validate: func(t *testing.T, resp *http.Response) {
				require.Equal(t, http.StatusOK, resp.StatusCode)

				var outputDTO dto.ListEmployeesOutputDTO

				err := json.NewDecoder(resp.Body).Decode(&outputDTO)
				require.NoError(t, err)

				require.Len(t, outputDTO.Items, 1)
				require.NoError(t, uuid.Validate(outputDTO.Items[0].PublicID))
				require.Equal(t, "Andrew Silva", outputDTO.Items[0].Name)
				require.NotZero(t, outputDTO.Items[0].CreatedAt)
				require.NotZero(t, outputDTO.Items[0].UpdatedAt)
			},
		},
	}

	for scenario, testCase := range testCases {
		suite.Run(scenario, func() {
			testCase.setup(suite.T())

			url := test.AddQueryParameters(suite.BasePath, map[string]any{
				"page":  testCase.page,
				"limit": testCase.limit,
			})

			req, err := test.NewRequest(suite.Method, url, nil)
			require.NoError(suite.T(), err)

			resp, err := apiApp.Test(req)
			require.NoError(suite.T(), err)

			testCase.validate(suite.T(), resp)
		})
	}
}

func TestListEmployeesWebHandlerTestSuiteIntegration(t *testing.T) {
	test.SkipTestIfShortFlagWasPassed(t)
	suite.Run(t, new(ListEmployeesWebHandlerTestSuite))
}
