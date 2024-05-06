package files

import (
	"database/sql"
	"github.com/eduardoraider/go-box/internal/bucket"
	"github.com/eduardoraider/go-box/internal/queue"
	"github.com/go-chi/chi"
)

type handler struct {
	db     *sql.DB
	bucket *bucket.Bucket
	queue  *queue.Queue
}

func setRoutes(r chi.Router, db *sql.DB, b *bucket.Bucket, q *queue.Queue) {
	h := handler{db, b, q}

	r.Post("/", h.Create)
	r.Put("/{id}", h.Modify)
	r.Delete("/{id}", h.Delete)
}
