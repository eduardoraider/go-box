package files

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"regexp"
	"time"
)

func (ts *TransactionSuite) TestGet() {
	setMockGet(ts.mock)

	_, err := Get(ts.conn, 1)
	assert.NoError(ts.T(), err)
}

func setMockGet(mock sqlmock.Sqlmock) {
	rows := sqlmock.NewRows([]string{"id", "folder_id", "owner_id", "name", "type", "path", "created_at", "modified_at", "deleted"}).
		AddRow(1, 1, 1, "Gopher.png", "image/png", "/", time.Now(), time.Now(), false)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM files WHERE id=$1`)).
		WithArgs(1).
		WillReturnRows(rows)
}
