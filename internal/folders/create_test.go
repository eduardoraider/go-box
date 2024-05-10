package folders

import (
	"bytes"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"net/http"
	"net/http/httptest"
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

	h := handler{db}

	f := Folder{
		Name: "Photos",
	}

	var b bytes.Buffer
	err = json.NewEncoder(&b).Encode(&f)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when encoding user", err)
	}

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/", &b)

	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO folders (parent_id, name, modified_at) VALUES ($1, $2, $3)`)).
		WithArgs(0, "Photos", f.ModifiedAt).
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

	f, err := New("Photos", 0)
	if err != nil {
		t.Error(err)
	}

	f.ModifiedAt = time.Now()

	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO folders (parent_id, name, modified_at) VALUES ($1, $2, $3)`)).
		WithArgs(0, "Photos", f.ModifiedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = Insert(db, f)
	if err != nil {
		t.Error(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}
}
