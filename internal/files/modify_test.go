package files

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

	h := handler{db, nil, nil}

	f := File{
		ID:   1,
		Name: "learning-golang.png",
	}

	var b bytes.Buffer
	err = json.NewEncoder(&b).Encode(&f)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when encoding user", err)
	}

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/{id}", &b)

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

	rows := sqlmock.NewRows([]string{"id", "folder_id", "owner_id", "name", "type", "path", "created_at", "modified_at", "deleted"}).
		AddRow(1, 1, 1, "learning-golang.png", "image/png", "/", time.Now(), time.Now(), false)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM files WHERE id = $1;`)).
		WithArgs(1).
		WillReturnRows(rows)

	mock.ExpectExec(regexp.QuoteMeta(`UPDATE files SET name=$1, modified_at=$2, deleted=$3 WHERE id=$4`)).
		WithArgs(f.Name, AnyTime{}, false, f.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

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

	mock.ExpectExec(regexp.QuoteMeta(`UPDATE files SET name=$1, modified_at=$2, deleted=$3 WHERE id=$4`)).
		WithArgs("Gopher-SP", AnyTime{}, false, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = Update(db, 1, &File{Name: "Gopher-SP"})
	if err != nil {
		t.Error(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}
}
