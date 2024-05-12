package users

import (
	"errors"
	"time"
)

func (h *handler) authenticate(login, password string) (*User, error) {
	stmt := `SELECT * FROM users WHERE login=$1`
	row := h.db.QueryRow(stmt, login)

	var u User
	err := row.Scan(&u.ID, &u.Name, &u.Login, &u.Password,
		&u.CreatedAt, &u.ModifiedAt, &u.Deleted, &u.LastLogin)
	if err != nil {
		return nil, err
	}

	if !VerifyPassword(u.Password, password) {
		return nil, errors.New("invalid password")
	}

	return &u, nil
}

func (h *handler) updateLogin(u *User) error {
	u.LastLogin = time.Now()
	return Update(h.db, u.ID, u)
}

func Authenticate(login, password string) (u *User, err error) {
	u, err = gh.authenticate(login, password)
	if err != nil {
		return
	}

	err = gh.updateLogin(u)
	if err != nil {
		return nil, err
	}

	return
}
