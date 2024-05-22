package factories

import (
	"errors"
	domain "github.com/eduardoraider/go-box/internal/users"
	"github.com/eduardoraider/go-box/repositories"
)

type scan interface {
	Scan(...any) error
}

func NewUserFactory(repo repositories.UserReadRepository) *UserFactory {
	return &UserFactory{repo}
}

type UserFactory struct {
	repo repositories.UserReadRepository
}

func restore(row scan) (*domain.User, error) {
	var u domain.User
	err := row.Scan(&u.ID, &u.Name, &u.Login, &u.Password,
		&u.CreatedAt, &u.ModifiedAt, &u.Deleted, &u.LastLogin)

	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (uf *UserFactory) RestoreOne(id int64) (*domain.User, error) {
	row := uf.repo.Get(id)
	return restore(row)
}

func (uf *UserFactory) Authenticate(login, password string) (*domain.User, error) {

	row := uf.repo.Login(login)
	u, err := restore(row)
	if err != nil {
		return nil, err
	}

	if !domain.VerifyPassword(u.Password, password) {
		return nil, errors.New("invalid password")
	}

	return u, nil
}

func (uf *UserFactory) RestoreAll() ([]domain.User, error) {
	rows, err := uf.repo.List()
	if err != nil {
		return nil, err
	}

	us := make([]domain.User, 0)
	for rows.Next() {
		u, err := restore(rows)
		if err != nil {
			continue
		}

		us = append(us, *u)
	}

	return us, nil
}
