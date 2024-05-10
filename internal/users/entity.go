package users

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var (
	ErrNameRequired     = errors.New("username is required")
	ErrLoginRequired    = errors.New("username is required")
	ErrPasswordRequired = errors.New("password is required and cannot be empty")
	ErrPasswordLength   = errors.New("password must be at least 8 characters")
)

func New(name, login, password string) (*User, error) {

	u := User{
		Name:       name,
		Login:      login,
		ModifiedAt: time.Now(),
	}

	err := u.SetPassword(password)
	if err != nil {
		return nil, err
	}

	err = u.Validate()
	if err != nil {
		return nil, err
	}

	return &u, nil
}

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

func (u *User) SetPassword(password string) error {
	if password == "" {
		return ErrPasswordRequired
	}
	if len(password) < 8 {
		return ErrPasswordLength
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("error hashing password", err)
	}
	u.Password = string(hashedPassword)

	return nil
}

func (u *User) Validate() error {
	if u.Name == "" {
		return ErrNameRequired
	}
	if u.Login == "" {
		return ErrLoginRequired
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(""), bcrypt.DefaultCost)
	if u.Password == string(hashedPassword) {
		return ErrPasswordRequired
	}

	return nil
}
