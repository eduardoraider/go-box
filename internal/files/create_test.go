package files

import (
	"github.com/DATA-DOG/go-sqlmock"
	"regexp"
	"testing"
	"time"
)

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
