package repository_test

import (
	"context"
	"testing"
	"tt-go-sample-api/domain/entity"
	"tt-go-sample-api/domain/infra/postgresql/repository"
	db "tt-go-sample-api/external/rdb/sqlc"
	"tt-go-sample-api/test"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type GetEmployeesPostgreSQLRepositoryTestSuite struct {
	suite.Suite
}

func (suite *GetEmployeesPostgreSQLRepositoryTestSuite) SetupSuite() {
	test.SetupTestsPostgreSQL(suite.T(), apiConfigTestSingleton.DBName)
	test.TruncateTables(suite.T(), db.SQLStoreSingleton.GetDB())
}

func (suite *GetEmployeesPostgreSQLRepositoryTestSuite) TearDownSubTest() {
	test.TruncateTables(suite.T(), db.SQLStoreSingleton.GetDB())
}

func (suite *GetEmployeesPostgreSQLRepositoryTestSuite) TearDownSuite() {
	require.NoError(suite.T(), db.SQLStoreSingleton.GetDB().Close())
}

func (suite *GetEmployeesPostgreSQLRepositoryTestSuite) TestListEmployees() {
	sut := repository.NewGetEmployeesPostgreSQLRepository()
	ctx := context.Background()

	testCases := map[string]struct {
		page     int32
		limit    int32
		setup    func(t *testing.T)
		validate func(t *testing.T, employees []*entity.Employee, err error)
	}{
		"Success with No Employees": {
			page:  1,
			limit: 20,
			setup: func(t *testing.T) {},
			validate: func(t *testing.T, employees []*entity.Employee, err error) {
				require.NoError(suite.T(), err)
				require.Len(suite.T(), employees, 0)
			},
		},
		"Success with Employees": {
			page:  1,
			limit: 20,
			setup: func(t *testing.T) {
				employeeName := "Andrew"
				createdEmployee := test.PersistFakeEmployee(t, db.SQLStoreSingleton.GetDB(), employeeName)
				require.NotZero(t, createdEmployee.ID)
				require.Equal(t, employeeName, createdEmployee.Name)
			},
			validate: func(t *testing.T, employees []*entity.Employee, err error) {
				require.NoError(suite.T(), err)
				require.Len(suite.T(), employees, 1)
				require.Equal(t, "Andrew", employees[0].Name)
			},
		},
	}

	for scenario, testCase := range testCases {
		suite.Run(scenario, func() {
			testCase.setup(suite.T())
			employees, err := sut.ListEmployees(ctx, testCase.page, testCase.limit)
			testCase.validate(suite.T(), employees, err)
		})
	}
}

func TestGetEmployeesPostgreSQLRepositoryTestSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test (-short flag provided)")
	}

	suite.Run(t, new(GetEmployeesPostgreSQLRepositoryTestSuite))
}
