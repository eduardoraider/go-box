package folders

import (
	"database/sql"
	"github.com/eduardoraider/go-box/internal/auth"
	"github.com/go-chi/chi"
)

type handler struct {
	db *sql.DB
}

func SetRoutes(r chi.Router, db *sql.DB) {
	h := handler{db}

	r.Route("/folders", func(r chi.Router) {
		r.Use(auth.Middleware)

		r.Post("/", h.Create)
		r.Put("/{id}", h.Modify)
		r.Delete("/{id}", h.Delete)
		r.Get("/", h.List)
		r.Get("/{id}", h.Get)
	})
}
