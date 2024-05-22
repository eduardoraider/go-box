package users

import (
	"database/sql"
	"database/sql/driver"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/eduardoraider/go-box/factories"
	"github.com/eduardoraider/go-box/internal/users"
	"github.com/eduardoraider/go-box/repositories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

type TransactionSuite struct {
	suite.Suite
	conn    *sql.DB
	mock    sqlmock.Sqlmock
	handler handler
	entity  *users.User
}

func (ts *TransactionSuite) SetupTest() {
	var err error
	ts.conn, ts.mock, err = sqlmock.New()
	assert.NoError(ts.T(), err)

	repo := repositories.NewUserRepository(ts.conn)
	factory := factories.NewUserFactory(repo)

	ts.handler = handler{repo, factory}

	ts.entity = &users.User{
		Name:     "Eduardo",
		Login:    "wookye.dev@gmail.com",
		Password: "12345678",
	}
	assert.NoError(ts.T(), err)
}

func (ts *TransactionSuite) TearDownSuite(_, _ string) {
	assert.NoError(ts.T(), ts.mock.ExpectationsWereMet())
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TransactionSuite))
}
