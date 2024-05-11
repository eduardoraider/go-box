package files

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"regexp"
	"time"
)

func (ts *TransactionSuite) TestList() {
	setMockList(ts.mock)

	_, err := List(ts.conn, 1)
	assert.NoError(ts.T(), err)
}

func (ts *TransactionSuite) TestListRoot() {
	setMockListRoot(ts.mock)

	_, err := ListRoot(ts.conn)
	assert.NoError(ts.T(), err)

}

func setMockList(mock sqlmock.Sqlmock) {
	rows := sqlmock.NewRows([]string{"id", "folder_id", "owner_id", "name", "type", "path", "created_at", "modified_at", "deleted"}).
		AddRow(1, 1, 1, "Gopher.png", "image/png", "/", time.Now(), time.Now(), false).
		AddRow(2, 1, 1, "Golang.jpg", "image/jpg", "/", time.Now(), time.Now(), false)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM files WHERE folder_id=$1 AND deleted=false`)).
		WithArgs(1).
		WillReturnRows(rows)
}

func setMockListRoot(mock sqlmock.Sqlmock) {
	rows := sqlmock.NewRows([]string{"id", "folder_id", "owner_id", "name", "type", "path", "created_at", "modified_at", "deleted"}).
		AddRow(1, nil, 1, "Gopher.png", "image/png", "/", time.Now(), time.Now(), false).
		AddRow(2, nil, 1, "Golang.jpg", "image/jpg", "/", time.Now(), time.Now(), false)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM files WHERE folder_id IS NULL AND deleted=false`)).
		WillReturnRows(rows)
}
