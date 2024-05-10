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

func TestGet(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	h := handler{db}

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/{id}", nil)

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

	rows := sqlmock.NewRows([]string{"id", "parent_id", "name", "created_at", "modified_at", "deleted"}).
		AddRow(1, 2, "Docs", time.Now(), time.Now(), false)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM folders WHERE id = $1`)).
		WithArgs(1).
		WillReturnRows(rows)

	contentRows := sqlmock.NewRows([]string{"id", "parent_id", "name", "created_at", "modified_at", "deleted"}).
		AddRow(2, 3, "Projects", time.Now(), time.Now(), false).
		AddRow(4, 5, "Videos", time.Now(), time.Now(), false)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM folders WHERE parent_id = $1 AND deleted = false`)).
		WithArgs(1).
		WillReturnRows(contentRows)

	folderRows := sqlmock.NewRows([]string{"id", "folder_id", "owner_id", "name", "type", "path", "created_at", "modified_at", "deleted"}).
		AddRow(1, 1, 1, "Gopher.png", "image/png", "/", time.Now(), time.Now(), false).
		AddRow(2, 1, 1, "Golang.jpg", "image/jpg", "/", time.Now(), time.Now(), false)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM files WHERE folder_id=$1 AND deleted=false`)).
		WillReturnRows(folderRows)

	h.Get(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("error: %v", rr)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}

}

func TestGetFolder(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "parent_id", "name", "created_at", "modified_at", "deleted"}).
		AddRow(1, 2, "Docs", time.Now(), time.Now(), false)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM folders WHERE id = $1`)).
		WithArgs(1).
		WillReturnRows(rows)

	_, err = GetFolder(db, 1)
	if err != nil {
		t.Error(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}
}

func TestGetSubFolder(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "parent_id", "name", "created_at", "modified_at", "deleted"}).
		AddRow(2, 3, "Projects", time.Now(), time.Now(), false).
		AddRow(4, 5, "Videos", time.Now(), time.Now(), false)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM folders WHERE parent_id = $1 AND deleted = false`)).
		WithArgs(1).
		WillReturnRows(rows)

	_, err = getSubFolders(db, 1)
	if err != nil {
		t.Error(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}
}
