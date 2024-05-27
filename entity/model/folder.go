package model

import (
	"time"
)

// Folder represents a folder with name, description, creation time and a list of files.
type Folder struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`

	Owner   *User              `json:"-"`
	Files   map[string]*File   `json:"files"`
	Folders map[string]*Folder `json:"-"`
}

// NewFolder creates a new Folder.
func NewFolder(owner *User, name, description string) (*Folder, error) {
	err := ValidateInput(name)
	if err != nil {
		return nil, err
	}

	return &Folder{
		Name:        name,
		Description: description,
		CreatedAt:   time.Now(),
		Files:       make(map[string]*File),
		Owner:       owner,
		Folders:     make(map[string]*Folder),
	}, nil
}
