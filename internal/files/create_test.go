package files

import (
	"bytes"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/eduardoraider/go-box/internal/bucket"
	"github.com/eduardoraider/go-box/internal/queue"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	b, err := bucket.New(bucket.MockProvider, nil)
	if err != nil {
		t.Error(err)
	}

	q, err := queue.New(queue.Mock, nil)
	if err != nil {
		t.Error(err)
	}

	h := handler{db, b, q}

	// Start Upload
	body := new(bytes.Buffer)

	mw := multipart.NewWriter(body)

	file, err := os.Open("./testdata/test_gopher.png")
	if err != nil {
		t.Error(err)
	}

	defer file.Close()

	w, err := mw.CreateFormFile("file", "test_gopher.png")
	if err != nil {
		t.Error(err)
	}

	_, err = io.Copy(w, file)
	if err != nil {
		t.Error(err)
	}

	mw.Close()
	// End upload

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/", body)
	req.Header.Add("Content-Type", mw.FormDataContentType())

	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO files (folder_id, owner_id, name, type, path, modified_at) VALUES ($1, $2, $3, $4, $5, $6)`)).
		WithArgs(0, 1, "test_gopher.png", "application/octet-stream", "/test_gopher.png", AnyTime{}).
		WillReturnResult(sqlmock.NewResult(1, 1))

	h.Create(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("error: %v", rr)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}

}

func TestInsert(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	f, err := New(1, "Gopher.png", "image/png", "/")
	if err != nil {
		t.Error(err)
	}

	f.FolderId = 1
	f.ModifiedAt = time.Now()

	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO files (folder_id, owner_id, name, type, path, modified_at) VALUES ($1, $2, $3, $4, $5, $6)`)).
		WithArgs(1, 1, "Gopher.png", "image/png", "/", f.ModifiedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = Insert(db, *f)
	if err != nil {
		t.Error(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}
}
