package folders

import (
	"database/sql"
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
	"time"
)

func (h *handler) Modify(rw http.ResponseWriter, r *http.Request) {
	f := new(Folder)

	err := json.NewDecoder(r.Body).Decode(f)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = f.Validate()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = Update(h.db, int64(id), f)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO: Get folder

	rw.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(rw).Encode(f)
	if err != nil {
		return
	}
}

func Update(db *sql.DB, id int64, f *Folder) error {
	f.ModifiedAt = time.Now()

	stmt := `UPDATE folder SET name=$1, modified_at=$2 WHERE id=$3`
	_, err := db.Exec(stmt, f.Name, f.ModifiedAt, id)
	return err
}
