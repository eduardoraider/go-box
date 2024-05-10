package users

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-chi/chi"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"
)

func TestModify(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	h := handler{db}

	u := User{
		ID:   1,
		Name: "Eduardo Raider",
	}

	var b bytes.Buffer
	err = json.NewEncoder(&b).Encode(&u)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when encoding user", err)
	}

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/{id}", &b)

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

	mock.ExpectExec(regexp.QuoteMeta(`UPDATE users SET name=$1, modified_at=$2 WHERE id=$3`)).
		WithArgs("Eduardo Raider", AnyTime{}, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	rows := sqlmock.NewRows([]string{"id", "name", "login", "password", "created_at", "modified_at", "deleted", "last_login"}).
		AddRow(1, "Eduardo", "wookye.dev@gmail.com", "12345678", time.Now(), time.Now(), false, time.Now())

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM users WHERE id=$1`)).
		WithArgs(1).
		WillReturnRows(rows)

	h.Modify(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("error: %v", rr)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}

}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta(`UPDATE users SET name=$1, modified_at=$2 WHERE id=$3`)).
		WithArgs("Eduardo", AnyTime{}, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = Update(db, 1, &User{Name: "Eduardo"})
	if err != nil {
		t.Error(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}
}
