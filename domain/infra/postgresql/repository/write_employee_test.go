package repository_test

import (
	"context"
	"testing"
	"tt-go-sample-api/domain/entity"
	"tt-go-sample-api/domain/infra/postgresql/repository"
	db "tt-go-sample-api/external/rdb/sqlc"
	"tt-go-sample-api/test"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type WriteEmployeePostgreSQLRepositoryTestSuite struct {
	suite.Suite
}

func (suite *WriteEmployeePostgreSQLRepositoryTestSuite) SetupSuite() {
	test.SetupTestsPostgreSQL(suite.T(), apiConfigTestSingleton.GetPostgreSQLSource(), apiConfigTestSingleton.DBName)
	test.TruncateTables(suite.T(), db.SQLStoreSingleton.GetDB())
}

func (suite *WriteEmployeePostgreSQLRepositoryTestSuite) TearDownSubTest() {
	test.TruncateTables(suite.T(), db.SQLStoreSingleton.GetDB())
}

func (suite *WriteEmployeePostgreSQLRepositoryTestSuite) TearDownSuite() {
	require.NoError(suite.T(), db.SQLStoreSingleton.CloseDB())
}

func (suite *WriteEmployeePostgreSQLRepositoryTestSuite) TestRegister() {
	sut := repository.NewWriteEmployeePostgreSQLRepository()
	ctx := context.Background()

	employeeName := "Andrew Silva"

	testCases := map[string]struct {
		setup    func(t *testing.T) *entity.Employee
		validate func(t *testing.T, newEmployee *entity.Employee, err error)
	}{
		"Success": {
			setup: func(t *testing.T) *entity.Employee {
				return entity.NewEmployee(employeeName)
			},
			validate: func(t *testing.T, newEmployee *entity.Employee, err error) {
				require.NoError(t, err)
				require.Equal(t, employeeName, newEmployee.Name)
				require.True(t, newEmployee.ID > int64(0))
				require.NoError(t, uuid.Validate(newEmployee.PublicID))
				require.NotZero(t, newEmployee.CreatedAt)
				require.NotZero(t, newEmployee.UpdatedAt)
			},
		},
	}

	for scenario, testCase := range testCases {
		suite.Run(scenario, func() {
			employee := testCase.setup(suite.T())
			newEmployee, err := sut.Register(ctx, employee)
			testCase.validate(suite.T(), newEmployee, err)
		})
	}
}

func TestWriteEmployeePostgreSQLRepositoryTestSuiteIntegration(t *testing.T) {
	test.SkipTestIfShortFlagWasPassed(t)
	suite.Run(t, new(WriteEmployeePostgreSQLRepositoryTestSuite))
}
