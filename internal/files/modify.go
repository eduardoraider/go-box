package files

import (
	"database/sql"
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
	"time"
)

func (h *handler) Modify(rw http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	file, err := Get(h.db, int64(id))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&file)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = file.Validate()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	err = Update(h.db, int64(file.ID), file)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(rw).Encode(file)
	if err != nil {
		return
	}

}

func Update(db *sql.DB, id int64, f *File) error {
	f.ModifiedAt = time.Now()

	stmt := `UPDATE file SET name=$1, modified_at=$2, delete=$3 WHERE id=$4`
	_, err := db.Exec(stmt, f.Name, f.ModifiedAt, f.Deleted, id)
	return err
}
