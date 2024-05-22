package repositories

import (
	"database/sql"
	"database/sql/driver"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/eduardoraider/go-box/internal/users"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"regexp"
	"testing"
	"time"
)

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

type TransactionUserSuite struct {
	suite.Suite
	conn   *sql.DB
	mock   sqlmock.Sqlmock
	repo   UserRepository
	entity *users.User
}

func (ts *TransactionUserSuite) SetupTest() {
	var err error
	ts.conn, ts.mock, err = sqlmock.New()
	assert.NoError(ts.T(), err)

	ts.repo = UserRepository{ts.conn}

	ts.entity = &users.User{
		Name:     "Eduardo",
		Login:    "wookye.dev@gmail.com",
		Password: "12345678",
	}
	assert.NoError(ts.T(), err)
}

func (ts *TransactionUserSuite) TearDownSuite(_, _ string) {
	assert.NoError(ts.T(), ts.mock.ExpectationsWereMet())
}

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(TransactionUserSuite))
}

func (ts *TransactionUserSuite) TestUserCreate() {
	setMockUserCreate(ts.mock, ts.entity, false)
	_, err := ts.repo.Create(ts.entity)
	assert.NoError(ts.T(), err)
}

func setMockUserCreate(mock sqlmock.Sqlmock, entity *users.User, err bool) {
	exp := mock.ExpectQuery(`INSERT INTO users(name, login, password, modified_at)* `).
		WithArgs(entity.Name, entity.Login, sqlmock.AnyArg(), entity.ModifiedAt)

	if err {
		exp.WillReturnError(sql.ErrConnDone)
	} else {
		exp.WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	}

}

func (ts *TransactionUserSuite) TestUserUpdate() {
	setMockUserUpdate(ts.mock)

	err := ts.repo.Update(1, &users.User{Name: "Eduardo"})
	assert.NoError(ts.T(), err)
}

func setMockUserUpdate(mock sqlmock.Sqlmock) {
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE users SET name=$1, modified_at=$2, last_login=$3  WHERE id=$4`)).
		WithArgs("Eduardo", AnyTime{}, AnyTime{}, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
}

func (ts *TransactionUserSuite) TestUserDelete() {
	setMockUserDelete(ts.mock, 1, false)

	err := ts.repo.Delete(1)
	assert.NoError(ts.T(), err)
}

func setMockUserDelete(mock sqlmock.Sqlmock, id int64, err bool) {
	exp := mock.ExpectExec(`UPDATE users SET *`).
		WithArgs(AnyTime{}, 1)

	if err {
		exp.WillReturnError(sql.ErrNoRows)
	} else {
		exp.WillReturnResult(sqlmock.NewResult(1, 1))
	}

}

func (ts *TransactionUserSuite) TestUserList() {
	setMockUserList(ts.mock)

	_, err := ts.repo.List()
	assert.NoError(ts.T(), err)
}

func setMockUserList(mock sqlmock.Sqlmock) {
	rows := sqlmock.NewRows([]string{"id", "name", "login", "password", "created_at", "modified_at", "deleted", "last_login"}).
		AddRow(1, "Eduardo", "wookye.dev@gmail.com", "12345678", time.Now(), time.Now(), false, time.Now()).
		AddRow(2, "John", "john-doe@example.com", "12345678", time.Now(), time.Now(), false, time.Now())

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM users WHERE deleted=false`)).
		WillReturnRows(rows)
}
