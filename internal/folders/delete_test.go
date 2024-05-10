package folders

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-chi/chi"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"
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

	filesRows := sqlmock.NewRows([]string{"id", "folder_id", "owner_id", "name", "type", "path", "created_at", "modified_at", "deleted"}).
		AddRow(1, 1, 1, "Gopher.png", "image/png", "/", time.Now(), time.Now(), false).
		AddRow(2, 1, 1, "Golang.jpg", "image/jpg", "/", time.Now(), time.Now(), false)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM files WHERE folder_id=$1 AND deleted=false`)).
		WillReturnRows(filesRows)

	mock.ExpectExec(regexp.QuoteMeta(`UPDATE files SET name=$1, modified_at=$2, deleted=$3 WHERE id=$4`)).
		WithArgs("Gopher.png", AnyTime{}, true, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec(regexp.QuoteMeta(`UPDATE files SET name=$1, modified_at=$2, deleted=$3 WHERE id=$4`)).
		WithArgs("Golang.jpg", AnyTime{}, true, 2).
		WillReturnResult(sqlmock.NewResult(1, 1))

	foldersRows := sqlmock.NewRows([]string{"id", "parent_id", "name", "created_at", "modified_at", "deleted"}).
		AddRow(1, 1, "Docs", time.Now(), time.Now(), false).
		AddRow(2, 1, "Contracts", time.Now(), time.Now(), false)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM folders WHERE parent_id = $1 AND deleted = false`)).
		WithArgs(1).
		WillReturnRows(foldersRows)

	mock.ExpectExec(regexp.QuoteMeta(`UPDATE folders SET modified_at=$1, deleted=true WHERE id=$2`)).
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

	mock.ExpectExec(regexp.QuoteMeta(`UPDATE folders SET modified_at=$1, deleted=true WHERE id=$2`)).
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
