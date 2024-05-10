package users

import (
	"bytes"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	h := handler{db}

	u := User{
		Name:     "Eduardo",
		Login:    "wookye.dev@gmail.com",
		Password: "12345678",
	}

	var b bytes.Buffer
	err = json.NewEncoder(&b).Encode(&u)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when encoding user", err)
	}

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/", &b)

	mock.ExpectExec(`INSERT INTO users(name, login, password, modified_at)*`).
		WithArgs(u.Name, u.Login, sqlmock.AnyArg(), u.ModifiedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	h.Create(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("error: %v", rr)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}

}

func TestPasswordHashing(t *testing.T) {
	u := User{
		Password: "12345678",
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		t.Errorf("Error generating password hash: %v", err)
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(u.Password))
	if err != nil {
		t.Errorf("The provided password does not match the generated hash: %v", err)
	}
}

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
