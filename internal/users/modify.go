package users

import (
	"database/sql"
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
	"time"
)

func (h *handler) Modify(rw http.ResponseWriter, r *http.Request) {
	u := new(User)

	err := json.NewDecoder(r.Body).Decode(u)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if u.Name == "" {
		http.Error(rw, ErrNameRequired.Error(), http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = Update(h.db, int64(id), u)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	u, err = Get(h.db, int64(id))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(rw).Encode(u)
	if err != nil {
		return
	}
}

func Update(db *sql.DB, id int64, u *User) error {
	u.ModifiedAt = time.Now()

	stmt := `UPDATE users SET name=$1, modified_at=$2, last_login=$3 WHERE id=$4`
	_, err := db.Exec(stmt, u.Name, u.ModifiedAt, u.LastLogin, id)
	return err
}
