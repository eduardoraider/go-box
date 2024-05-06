package users

import (
	"database/sql"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
	"time"
)

func (h *handler) Delete(rw http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = Delete(h.db, int64(id))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Add("Content-Type", "application/json")
}

func Delete(db *sql.DB, id int64) error {
	stmt := `UPDATE users SET modified_at=$1, deleted=true WHERE id=$2`
	_, err := db.Exec(stmt, time.Now(), id)
	return err
}
