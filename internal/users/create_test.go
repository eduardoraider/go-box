package users

import (
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

func TestInsert(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	u, err := New("Eduardo", "wookye.dev@gmail.com", "12345678")
	if err != nil {
		t.Error(err)
	}

	mock.ExpectExec(`INSERT INTO users(name, login, password, modified_at)*`).
		WithArgs("Eduardo", "wookye.dev@gmail.com", u.Password, u.ModifiedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = Insert(db, u)
	if err != nil {
		t.Error(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}
}
