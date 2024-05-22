package users

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"

	domain "github.com/eduardoraider/go-box/internal/users"
)

func (h *handler) Modify(rw http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	u, err := h.factory.RestoreOne(int64(id))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	u, err = domain.DecodeAndModify(r.Body, u)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.repo.Update(int64(id), u)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(u)
}
