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

func TestList(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	h := handler{db}

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

	foldersRows := sqlmock.NewRows([]string{"id", "parent_id", "name", "created_at", "modified_at", "deleted"}).
		AddRow(2, nil, "Projects", time.Now(), time.Now(), false).
		AddRow(4, nil, "Videos", time.Now(), time.Now(), false)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM folders WHERE parent_id IS NULL AND deleted = false`)).
		WillReturnRows(foldersRows)

	folderRows := sqlmock.NewRows([]string{"id", "folder_id", "owner_id", "name", "type", "path", "created_at", "modified_at", "deleted"}).
		AddRow(1, nil, 1, "Gopher.png", "image/png", "/", time.Now(), time.Now(), false).
		AddRow(2, nil, 1, "Golang.jpg", "image/jpg", "/", time.Now(), time.Now(), false)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM files WHERE folder_id IS NULL AND deleted=false`)).
		WillReturnRows(folderRows)

	h.List(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("error: %v", rr)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}

}

func TestGetRootSubfolder(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "parent_id", "name", "created_at", "modified_at", "deleted"}).
		AddRow(1, nil, "Docs", time.Now(), time.Now(), false).
		AddRow(2, nil, "Contracts", time.Now(), time.Now(), false)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM folders WHERE parent_id IS NULL AND deleted = false`)).
		WillReturnRows(rows)

	_, err = getRootSubFolders(db)
	if err != nil {
		t.Error(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}
}
