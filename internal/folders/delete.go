package folders

import (
	"database/sql"
	"github.com/eduardoraider/go-box/internal/files"
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

	err = deleteFolderContent(h.db, int64(id))
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

func deleteFolderContent(db *sql.DB, folderId int64) error {
	err := deleteFiles(db, folderId)
	if err != nil {
		return err
	}

	return deleteSubFolders(db, folderId)
}

func deleteSubFolders(db *sql.DB, folderId int64) error {
	subFolders, err := getSubFolders(db, folderId)
	if err != nil {
		return err
	}

	removedFolders := make([]Folder, 0, len(subFolders))
	for _, sf := range subFolders {
		err := Delete(db, sf.ID)
		if err != nil {
			break
		}

		err = deleteFolderContent(db, sf.ID)
		if err != nil {
			err := Update(db, sf.ID, &sf)
			if err != nil {
				break
			}
			break
		}

		removedFolders = append(removedFolders, sf)
	}

	if len(removedFolders) != len(subFolders) {
		for _, sf := range removedFolders {
			err := Update(db, sf.ID, &sf)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func deleteFiles(db *sql.DB, folderId int64) error {
	f, err := files.List(db, int64(folderId))
	if err != nil {
		return err
	}

	removedFiles := make([]files.File, 0, len(f))
	for _, file := range f {
		file.Deleted = true
		err := files.Update(db, file.ID, &file)
		if err != nil {
			break
		}
		removedFiles = append(removedFiles, file)
	}

	if len(f) != len(removedFiles) {
		for _, file := range removedFiles {
			file.Deleted = false
			err := files.Update(db, file.ID, &file)
			if err != nil {
				return err
			}
		}
		return err
	}

	return nil
}

func Delete(db *sql.DB, id int64) error {
	stmt := `UPDATE folders SET modified_at=$1, deleted=true WHERE id=$2`
	_, err := db.Exec(stmt, time.Now(), id)
	return err
}
