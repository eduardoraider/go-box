package users

import (
	"github.com/DATA-DOG/go-sqlmock"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"
)

func TestList(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	h := handler{db}

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	rows := sqlmock.NewRows([]string{"id", "name", "login", "password", "created_at", "modified_at", "deleted", "last_login"}).
		AddRow(1, "Eduardo", "wookye.dev@gmail.com", "12345678", time.Now(), time.Now(), false, time.Now()).
		AddRow(2, "John", "john-doe@example.com", "12345678", time.Now(), time.Now(), false, time.Now())

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM users WHERE deleted = false`)).
		WillReturnRows(rows)

	h.List(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("error: %v", rr)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}

}

func TestSelectAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "login", "password", "created_at", "modified_at", "deleted", "last_login"}).
		AddRow(1, "Eduardo", "wookye.dev@gmail.com", "12345678", time.Now(), time.Now(), false, time.Now()).
		AddRow(2, "John", "john-doe@example.com", "12345678", time.Now(), time.Now(), false, time.Now())

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM users WHERE deleted = false`)).
		WillReturnRows(rows)

	_, err = SelectAll(db)
	if err != nil {
		t.Error(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}
}
