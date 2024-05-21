package folders

import (
	"errors"
	"github.com/guregu/null/v5"
	"time"
)

var (
	ErrNameRequired = errors.New("folder name is required")
)

type Folder struct {
	ID         int64     `json:"id"`
	ParentId   null.Int  `json:"parent_id"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	Deleted    bool      `json:"-"`
}

func (f *Folder) Validate() error {
	if f.Name == "" {
		return ErrNameRequired
	}
	return nil
}

type FolderContent struct {
	Folder  Folder           `json:"folder"`
	Content []FolderResource `json:"content"`
}

type FolderResource struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}
