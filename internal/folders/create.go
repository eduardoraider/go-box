package folders

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func (h *handler) Create(rw http.ResponseWriter, r *http.Request) {
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

	id, err := Insert(h.db, f)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	f.ID = id

	rw.WriteHeader(http.StatusCreated)
	rw.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(rw).Encode(f)
	if err != nil {
		return
	}
}

func Insert(db *sql.DB, f *Folder) (id int64, err error) {
	stmt := `INSERT INTO folders (parent_id, name, modified_at) VALUES ($1, $2, $3) RETURNING id`
	err = db.QueryRow(stmt, f.ParentId, f.Name, f.ModifiedAt).Scan(&id)
	if err != nil {
		return -1, err
	}
	return
}
