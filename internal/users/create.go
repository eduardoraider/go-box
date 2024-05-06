package users

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func (h *handler) Create(rw http.ResponseWriter, r *http.Request) {
	u := new(User)

	err := json.NewDecoder(r.Body).Decode(u)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = u.SetPassword(u.Password)
	if err != nil {
		return
	}

	err = u.Validate()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := Insert(h.db, u)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	u.ID = id

	rw.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(rw).Encode(u)
	if err != nil {
		return
	}

}

func Insert(db *sql.DB, u *User) (int64, error) {
	stmt := `INSERT INTO users(name, login, password, modified_at) values($1, $2, $3, $4)`
	result, err := db.Exec(stmt, u.Name, u.Login, u.Password, u.ModifiedAt)
	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}
