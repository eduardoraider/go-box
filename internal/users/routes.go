package users

import (
	"database/sql"
	"github.com/go-chi/chi"
)

var gh handler

type handler struct {
	db *sql.DB
}

func SetRoutes(r chi.Router, db *sql.DB) {
	gh = handler{db}

	r.Post("/", gh.Create)
	r.Put("/{id}", gh.Modify)
	r.Delete("/{id}", gh.Delete)
	r.Get("/{id}", gh.GetByID)
	r.Get("/", gh.List)
}
