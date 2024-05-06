package files

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/eduardoraider/go-box/internal/queue"
	"net/http"
	"strconv"
)

func (h *handler) Create(rw http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		return
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	path := fmt.Sprintf("/%s", fileHeader.Filename)

	err = h.bucket.Upload(file, path)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	entity, err := New(1, fileHeader.Filename, fileHeader.Header.Get("Content-Type"), path)
	if err != nil {
		er := h.bucket.Delete(path)
		if er != nil {
			http.Error(rw, er.Error(), http.StatusBadRequest)
		}
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	folderID := r.Form.Get("folder_id")
	if folderID != "" {
		fid, err := strconv.Atoi(folderID)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		entity.FolderId = int64(fid)
	}

	id, err := Insert(h.db, *entity)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	entity.ID = id

	dto := queue.AppQueueDto{
		Filename: fileHeader.Filename,
		Path:     path,
		ID:       int(id),
	}

	msg, err := dto.Marshal()
	if err != nil {
		// TODO: rollback
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.queue.Publish(msg)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Add("Content-Type", "text/html")
	err = json.NewEncoder(rw).Encode(entity)
	if err != nil {
		return
	}

}

func Insert(db *sql.DB, f File) (int64, error) {
	stmt := `INSERT INTO files (folder_id, owner_id, name, type, path, modified_at) VALUES ($1, $2, $3, $4, $5, $6)`
	result, err := db.Exec(stmt, f.FolderId, f.OwnerId, f.Name, f.Type, f.Path, f.ModifiedAt)
	if err != nil {
		return 1, err
	}

	return result.LastInsertId()
}
