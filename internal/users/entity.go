package users

import "time"

func New(name, login, password string) (*User, error) {
	now := time.Now()
	u := User{
		Name:       name,
		Login:      login,
		Password:   password,
		CreatedAt:  now,
		ModifiedAt: now,
	}

	return &u, nil
}

type User struct {
	ID         int64
	Name       string
	Login      string
	Password   string
	CreatedAt  time.Time
	ModifiedAt time.Time
	Deleted    bool
	LastLogin  time.Time
}
