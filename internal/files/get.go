package files

import (
	"database/sql"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func (h *handler) Get(rw http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Get(db *sql.DB, id int64) (*File, error) {
	stmt := `SELECT * FROM files WHERE id = $1;`
	row := db.QueryRow(stmt, id)

	var f File
	err := row.Scan(&f.ID, &f.FolderId, &f.OwnerId, &f.Name, &f.Type,
		&f.Path, &f.CreatedAt, &f.ModifiedAt, &f.Deleted)
	if err != nil {
		return nil, err
	}
	return &f, nil
}
