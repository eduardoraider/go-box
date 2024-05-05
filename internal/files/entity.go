package files

import (
	"errors"
	"time"
)

var (
	ErrOwnerRequired = errors.New("owner required")
	ErrNameRequired  = errors.New("folder name is required")
	ErrTypeRequired  = errors.New("folder type is required")
	ErrPathRequired  = errors.New("folder path is required")
)

type File struct {
	ID         int64     `json:"id"`
	FolderId   int64     `json:"-"`
	OwnerId    int64     `json:"owner_id"`
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	Path       string    `json:"-"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	Deleted    bool      `json:"-"`
}

func (f *File) Validate() error {
	if f.OwnerId == 0 {
		return ErrOwnerRequired
	}
	if f.Name == "" {
		return ErrNameRequired
	}
	if f.Type == "" {
		return ErrTypeRequired
	}
	if f.Path == "" {
		return ErrPathRequired
	}
	return nil
}
