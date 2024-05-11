package files

import "database/sql"

func List(db *sql.DB, folderId int64) ([]File, error) {
	stmt := `SELECT * FROM files WHERE folder_id=$1 AND deleted=false`
	return selectAllFiles(db, stmt, folderId)
}

func ListRoot(db *sql.DB) ([]File, error) {
	stmt := `SELECT * FROM files WHERE folder_id IS NULL AND deleted=false`
	return selectAllFiles(db, stmt, 0)
}

func selectAllFiles(db *sql.DB, stmt string, folderId int64) ([]File, error) {
	rows, err := db.Query(stmt, folderId)
	if err != nil {
		return nil, err
	}

	files := make([]File, 0)
	for rows.Next() {
		var f File

		err := rows.Scan(&f.ID, &f.FolderId, &f.OwnerId, &f.Name, &f.Type,
			&f.Path, &f.CreatedAt, &f.ModifiedAt, &f.Deleted)
		if err != nil {
			continue
		}

		files = append(files, f)
	}

	return files, nil
}
