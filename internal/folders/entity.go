package folders

import (
	"errors"
	"time"
)

var (
	ErrNameRequired = errors.New("folder name is required")
)

func New(name string, parentId int64) (*Folder, error) {
	f := Folder{
		Name: name,
	}

	if parentId > 0 {
		f.ParentId = parentId
	}

	err := f.Validate()
	if err != nil {
		return nil, err
	}

	return &f, nil
}

type Folder struct {
	ID         int64     `json:"id"`
	ParentId   int64     `json:"parent_id"`
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
