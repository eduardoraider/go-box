package files

import (
	"github.com/DATA-DOG/go-sqlmock"
	"regexp"
	"testing"
)

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta(`UPDATE file SET name=$1, modified_at=$2, delete=$3 WHERE id=$4`)).
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
