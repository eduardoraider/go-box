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

func VerifyPassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func encryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (u *User) SetPassword(password string) error {
	if password == "" {
		return ErrPasswordRequired
	}
	if len(password) < 8 {
		return ErrPasswordLength
	}
	hashedPassword, err := encryptPassword(password)
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

	hashedPassword, _ := encryptPassword("")
	if u.Password == hashedPassword {
		return ErrPasswordRequired
	}

	return nil
}
