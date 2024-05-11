package files

import (
	"database/sql"
)

func Get(db *sql.DB, id int64) (*File, error) {
	stmt := `SELECT * FROM files WHERE id=$1`
	row := db.QueryRow(stmt, id)

	var f File
	err := row.Scan(&f.ID, &f.FolderId, &f.OwnerId, &f.Name, &f.Type,
		&f.Path, &f.CreatedAt, &f.ModifiedAt, &f.Deleted)
	if err != nil {
		return nil, err
	}
	return &f, nil
}
