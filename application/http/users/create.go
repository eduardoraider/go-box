package users

import (
	"encoding/json"
	"net/http"

	domain "github.com/eduardoraider/go-box/internal/users"
)

func (h *handler) Create(rw http.ResponseWriter, r *http.Request) {
	u, err := domain.DecodeAndCreate(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.repo.Create(u)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	u.ID = id

	rw.WriteHeader(http.StatusCreated)
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(u)
}
