package model

import (
	"time"
)

// File represents a file in the virtual filesystem.
type File struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`

	Owner  *User   `json:"-"`
	Folder *Folder `json:"-"`
}

// NewFile creates a new File.
func NewFile(owner *User, folder *Folder, name, description string) (*File, error) {
	err := ValidateInput(name)
	if err != nil {
		return nil, err
	}

	return &File{
		Name:        name,
		Description: description,
		CreatedAt:   time.Now(),
		Owner:       owner,
		Folder:      folder,
	}, nil
}
