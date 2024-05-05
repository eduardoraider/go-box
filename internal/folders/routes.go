package folders

import (
	"database/sql"
	"github.com/go-chi/chi"
)

type handler struct {
	db *sql.DB
}

func SetRoutes(r chi.Router, db *sql.DB) {
	h := handler{db}

	r.Post("/", h.Create)
	r.Put("/{id}", h.Modify)
}
