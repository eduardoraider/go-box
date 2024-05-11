package folders

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"regexp"
	"time"
)

func (ts *TransactionSuite) TestList() {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	setMockList(ts.mock)
	setMockListRootFiles(ts.mock)

	ts.handler.List(rr, req)
	assert.Equal(ts.T(), http.StatusOK, rr.Code)
}

func (ts *TransactionSuite) TestGetRootSubfolders() {
	setMockList(ts.mock)

	_, err := getRootSubFolders(ts.conn)
	assert.NoError(ts.T(), err)
}

func setMockList(mock sqlmock.Sqlmock) {
	rows := sqlmock.NewRows([]string{"id", "parent_id", "name", "created_at", "modified_at", "deleted"}).
		AddRow(1, nil, "Docs", time.Now(), time.Now(), false).
		AddRow(5, nil, "Contracts", time.Now(), time.Now(), false)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM folders WHERE parent_id IS NULL AND deleted=false`)).
		WillReturnRows(rows)
}

func setMockListRootFiles(mock sqlmock.Sqlmock) {
	rows := sqlmock.NewRows([]string{"id", "folder_id", "owner_id", "name", "type", "path", "created_at", "modified_at", "deleted"}).
		AddRow(1, nil, 1, "Gopher.png", "image/png", "/", time.Now(), time.Now(), false).
		AddRow(2, nil, 1, "Golang.jpg", "image/jpg", "/", time.Now(), time.Now(), false)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM files WHERE folder_id IS NULL AND deleted=false`)).
		WillReturnRows(rows)
}
