package users

import (
	"bytes"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/http/httptest"
)

func (ts *TransactionSuite) TestCreate() {
	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(ts.entity)
	assert.NoError(ts.T(), err)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/", &b)

	setMockInsert(ts.mock, ts.entity)

	ts.handler.Create(rr, req)
	assert.Equal(ts.T(), http.StatusOK, rr.Code)
}

func (ts *TransactionSuite) TestPasswordHashing() {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(ts.entity.Password), bcrypt.DefaultCost)
	assert.NoError(ts.T(), err)

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(ts.entity.Password))
	assert.NoError(ts.T(), err)
}

func (ts *TransactionSuite) TestInsert() {
	setMockInsert(ts.mock, ts.entity)
	_, err := Insert(ts.conn, ts.entity)
	assert.NoError(ts.T(), err)
}

func setMockInsert(mock sqlmock.Sqlmock, entity *User) {
	mock.ExpectExec(`INSERT INTO users(name, login, password, modified_at)*`).
		WithArgs(entity.Name, entity.Login, entity.Password, entity.ModifiedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))
}
