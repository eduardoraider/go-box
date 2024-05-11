package users

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"regexp"
)

func (ts *TransactionSuite) TestModify() {
	u := User{
		ID:   1,
		Name: "Eduardo",
	}

	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(&u)
	assert.NoError(ts.T(), err)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/{id}", &b)

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

	setMockUpdate(ts.mock)
	setMockGet(ts.mock)

	ts.handler.Modify(rr, req)
	assert.Equal(ts.T(), http.StatusOK, rr.Code)
}

func (ts *TransactionSuite) TestUpdate() {
	setMockUpdate(ts.mock)

	err := Update(ts.conn, 1, &User{Name: "Eduardo"})
	assert.NoError(ts.T(), err)
}

func setMockUpdate(mock sqlmock.Sqlmock) {
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE users SET name=$1, modified_at=$2 WHERE id=$3`)).
		WithArgs("Eduardo", AnyTime{}, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
}
