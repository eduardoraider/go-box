package users

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"regexp"
	"time"
)

func (ts *TransactionSuite) TestGetByID() {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/{id}", nil)

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

	setMockGet(ts.mock)

	ts.handler.GetByID(rr, req)
	assert.Equal(ts.T(), http.StatusOK, rr.Code)
}

func setMockGet(mock sqlmock.Sqlmock) {
	rows := sqlmock.NewRows([]string{"id", "name", "login", "password", "created_at", "modified_at", "deleted", "last_login"}).
		AddRow(1, "Eduardo", "wookye.dev@gmail.com", "12345678", time.Now(), time.Now(), false, time.Now())

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM users WHERE id=$1`)).
		WithArgs(1).
		WillReturnRows(rows)
}
