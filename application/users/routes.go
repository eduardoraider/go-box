package users

import (
	"github.com/eduardoraider/go-box/factories"
	"github.com/eduardoraider/go-box/internal/auth"
	"github.com/eduardoraider/go-box/repositories"
	"github.com/go-chi/chi"
)

var gh handler

type handler struct {
	repo    repositories.UserWriteRepository
	factory *factories.UserFactory
}

func SetRoutes(r chi.Router, repo repositories.UserWriteRepository, uf *factories.UserFactory) {
	gh = handler{repo, uf}

	r.Route("/users", func(r chi.Router) {
		r.Post("/", gh.Create)

		r.Group(func(r chi.Router) {
			r.Use(auth.Middleware)

			r.Put("/{id}", gh.Modify)
			r.Delete("/{id}", gh.Delete)
			r.Get("/{id}", gh.GetByID)
			r.Get("/", gh.List)
		})
	})
}
