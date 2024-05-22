package users

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/http/httptest"

	domain "github.com/eduardoraider/go-box/internal/users"
)

func (ts *TransactionSuite) TestCreate() {
	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(ts.entity)
	assert.NoError(ts.T(), err)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/", &b)

	setMockInsert(ts.mock, ts.entity, false)

	ts.handler.Create(rr, req)
	assert.Equal(ts.T(), http.StatusCreated, rr.Code)
}

func (ts *TransactionSuite) TestPasswordHashing() {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(ts.entity.Password), bcrypt.DefaultCost)
	assert.NoError(ts.T(), err)

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(ts.entity.Password))
	assert.NoError(ts.T(), err)
}

func setMockInsert(mock sqlmock.Sqlmock, entity *domain.User, err bool) {
	exp := mock.ExpectQuery(`INSERT INTO users(name, login, password, modified_at)* `).
		WithArgs(entity.Name, entity.Login, sqlmock.AnyArg(), entity.ModifiedAt)

	if err {
		exp.WillReturnError(sql.ErrConnDone)
	} else {
		exp.WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	}

}
