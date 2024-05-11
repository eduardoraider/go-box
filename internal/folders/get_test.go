package folders

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

func (ts *TransactionSuite) TestGet() {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/{id}", nil)

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

	setMockGet(ts.mock)
	setMockGetSubFolder(ts.mock)
	setMockListFiles(ts.mock)

	ts.handler.Get(rr, req)
	assert.Equal(ts.T(), http.StatusOK, rr.Code)
}

func (ts *TransactionSuite) TestGetFolder() {
	setMockGet(ts.mock)

	_, err := GetFolder(ts.conn, 1)
	assert.NoError(ts.T(), err)
}

func (ts *TransactionSuite) TestGetSubFolder() {
	setMockGetSubFolder(ts.mock)

	_, err := getSubFolders(ts.conn, 1)
	assert.NoError(ts.T(), err)
}

func setMockGet(mock sqlmock.Sqlmock) {
	rows := sqlmock.NewRows([]string{"id", "parent_id", "name", "created_at", "modified_at", "deleted"}).
		AddRow(1, 1, "Docs", time.Now(), time.Now(), false)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM folders WHERE id=$1`)).
		WithArgs(1).
		WillReturnRows(rows)
}

func setMockGetSubFolder(mock sqlmock.Sqlmock) {
	rows := sqlmock.NewRows([]string{"id", "parent_id", "name", "created_at", "modified_at", "deleted"}).
		AddRow(1, 1, "Projects", time.Now(), time.Now(), false).
		AddRow(2, 1, "Videos", time.Now(), time.Now(), false)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM folders WHERE parent_id=$1 AND deleted=false`)).
		WithArgs(1).
		WillReturnRows(rows)

}
