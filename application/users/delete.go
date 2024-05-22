package users

import (
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func (h *handler) Delete(rw http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.repo.Delete(int64(id))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
	rw.Header().Add("Content-Type", "application/json")
}
