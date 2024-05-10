package users

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-chi/chi"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeleteHTTP(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	h := handler{db}

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/{id}", nil)

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

	mock.ExpectExec(`UPDATE users SET *`).
		WithArgs(AnyTime{}, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	h.Delete(rr, req)

	if rr.Code != http.StatusNoContent {
		t.Errorf("error: %v", rr)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}

}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec(`UPDATE users SET *`).
		WithArgs(AnyTime{}, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = Delete(db, 1)
	if err != nil {
		t.Error(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}
}
