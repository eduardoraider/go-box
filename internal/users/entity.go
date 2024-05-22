package users

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var (
	ErrNameRequired     = errors.New("username is required")
	ErrLoginRequired    = errors.New("username is required")
	ErrPasswordRequired = errors.New("password is required and cannot be empty")
	ErrPasswordLength   = errors.New("password must be at least 8 characters")
)

type User struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Login      string    `json:"login"`
	Password   string    `json:"password"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	Deleted    bool      `json:"-"`
	LastLogin  time.Time `json:"last_login"`
}

func (u *User) GetID() int64 {
	return u.ID
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) ChangeName(name string) error {
	if name == "" {
		return ErrNameRequired
	}

	u.Name = name

	return nil
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) GetLogin() string {
	return u.Login
}

func VerifyPassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func encryptPassword(u *User) error {
	if u.Password == "" {
		return ErrPasswordRequired
	}
	if len(u.Password) < 8 {
		return ErrPasswordLength
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)

	return nil
}

func (u *User) ChangePassword(password string) error {
	u.Password = password

	return encryptPassword(u)
}

func (u *User) Validate() error {
	if u.Name == "" {
		return ErrNameRequired
	}
	if u.Login == "" {
		return ErrLoginRequired
	}

	return nil
}
