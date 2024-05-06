package users

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
