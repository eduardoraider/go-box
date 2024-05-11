package files

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"regexp"
)

func (ts *TransactionSuite) TestDeleteHTTP() {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/{id}", nil)

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

	setMockDelete(ts.mock)

	ts.handler.Delete(rr, req)
	assert.Equal(ts.T(), http.StatusNoContent, rr.Code)
}

func (ts *TransactionSuite) TestDelete() {
	setMockDelete(ts.mock)

	err := Delete(ts.conn, 1)
	assert.NoError(ts.T(), err)
}

func setMockDelete(mock sqlmock.Sqlmock) {
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE files SET modified_at=$1, deleted=true WHERE id=$2`)).
		WithArgs(AnyTime{}, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
}
