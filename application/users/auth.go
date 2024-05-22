package users

import (
	"time"

	domain "github.com/eduardoraider/go-box/internal/users"
)

func (h *handler) authenticate(login, password string) (*domain.User, error) {
	return h.factory.Authenticate(login, password)
}

func (h *handler) updateLastLogin(u *domain.User) error {
	u.LastLogin = time.Now()
	return h.repo.Update(u.ID, u)
}

func Authenticate(login, password string) (u *domain.User, err error) {
	u, err = gh.authenticate(login, password)
	if err != nil {
		return
	}

	err = gh.updateLastLogin(u)
	if err != nil {
		return nil, err
	}

	return
}
