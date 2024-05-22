package services

import (
	"github.com/eduardoraider/go-box/factories"
	domain "github.com/eduardoraider/go-box/internal/users"
	"github.com/eduardoraider/go-box/repositories"
	"time"
)

func NewAuthService(repo repositories.UserWriteRepository, fact *factories.UserFactory) *AuthService {
	return &AuthService{
		repo:    repo,
		factory: fact,
	}
}

type AuthService struct {
	repo    repositories.UserWriteRepository
	factory *factories.UserFactory
}

func (svc *AuthService) authenticate(login, password string) (*domain.User, error) {
	return svc.factory.Authenticate(login, password)
}

func (svc *AuthService) updateLastLogin(u *domain.User) error {
	u.LastLogin = time.Now()
	return svc.repo.Update(u.ID, u)
}

func (svc *AuthService) Authenticate(login, password string) (u *domain.User, err error) {
	u, err = svc.authenticate(login, password)
	if err != nil {
		return
	}

	err = svc.updateLastLogin(u)
	if err != nil {
		return nil, err
	}

	return
}
