package folders

import (
	"database/sql"
	"encoding/json"
	"github.com/eduardoraider/go-box/internal/files"
	"net/http"
)

func (h *handler) List(rw http.ResponseWriter, r *http.Request) {
	c, err := GetRootFolderContent(h.db)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	fc := FolderContent{
		Folder: Folder{
			Name: "Root",
		},
		Content: c}

	rw.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(rw).Encode(fc)
	if err != nil {
		return
	}
}

func getRootSubFolders(db *sql.DB) ([]Folder, error) {
	stmt := `SELECT * FROM folders WHERE parent_id IS NULL AND deleted=false`
	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}

	f := make([]Folder, 0)
	for rows.Next() {
		var folder Folder
		err := rows.Scan(&folder.ID, &folder.ParentId, &folder.Name, &folder.CreatedAt, &folder.ModifiedAt, &folder.Deleted)
		if err != nil {
			continue
		}

		f = append(f, folder)
	}

	return f, nil
}

func GetRootFolderContent(db *sql.DB) ([]FolderResource, error) {
	subFolders, err := getRootSubFolders(db)
	if err != nil {
		return nil, err
	}

	fr := make([]FolderResource, 0, len(subFolders))
	for _, sf := range subFolders {
		r := FolderResource{
			ID:         sf.ID,
			Name:       sf.Name,
			Type:       "directory",
			CreatedAt:  sf.CreatedAt,
			ModifiedAt: sf.ModifiedAt,
		}
		fr = append(fr, r)
	}

	folderFiles, err := files.ListRoot(db)
	if err != nil {
		return nil, err
	}

	for _, f := range folderFiles {
		r := FolderResource{
			ID:         f.ID,
			Name:       f.Name,
			Type:       f.Type,
			CreatedAt:  f.CreatedAt,
			ModifiedAt: f.ModifiedAt,
		}
		fr = append(fr, r)
	}

	return fr, nil
}
