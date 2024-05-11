package files

import (
	"bytes"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
)

func (ts *TransactionSuite) TestCreate() {
	// Start Upload
	body := new(bytes.Buffer)

	mw := multipart.NewWriter(body)

	file, err := os.Open("./testdata/test_gopher.png")
	assert.NoError(ts.T(), err)

	w, err := mw.CreateFormFile("file", "test_gopher.png")
	assert.NoError(ts.T(), err)

	_, err = io.Copy(w, file)
	assert.NoError(ts.T(), err)

	mw.Close()
	// End upload

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/", body)
	req.Header.Add("Content-Type", mw.FormDataContentType())

	setMockInsert(ts.mock, ts.entity)

	ts.handler.Create(rr, req)
	assert.Equal(ts.T(), http.StatusCreated, rr.Code)
}

func (ts *TransactionSuite) TestInsert() {
	setMockInsert(ts.mock, ts.entity)

	_, err := Insert(ts.conn, ts.entity)
	assert.NoError(ts.T(), err)

}

func setMockInsert(mock sqlmock.Sqlmock, entity *File) {
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO files (folder_id, owner_id, name, type, path, modified_at) VALUES ($1, $2, $3, $4, $5, $6)`)).
		WithArgs(0, entity.OwnerId, entity.Name, entity.Type, entity.Path, AnyTime{}).
		WillReturnResult(sqlmock.NewResult(1, 1))
}
