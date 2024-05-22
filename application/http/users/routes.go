package users

import (
	"github.com/eduardoraider/go-box/factories"
	"github.com/eduardoraider/go-box/internal/auth"
	"github.com/eduardoraider/go-box/repositories"
	"github.com/go-chi/chi"
)

type handler struct {
	repo    repositories.UserWriteRepository
	factory *factories.UserFactory
}

func SetRoutes(r chi.Router, repo repositories.UserWriteRepository, uf *factories.UserFactory) {
	h := handler{repo, uf}

	r.Route("/users", func(r chi.Router) {
		r.Post("/", h.Create)

		r.Group(func(r chi.Router) {
			r.Use(auth.Middleware)

			r.Put("/{id}", h.Modify)
			r.Delete("/{id}", h.Delete)
			r.Get("/{id}", h.GetByID)
			r.Get("/", h.List)
		})
	})
}
