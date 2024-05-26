//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package vfs

import (
	"time"
)

// FolderManager defines the interface for folder management
type FolderManager interface {
	CreateFolder(username, foldername, description string) error
	DeleteFolder(username, foldername string) error
	RenameFolder(username, foldername, newFoldername string) error
	ListFolders(username string, sortBy string, order string) ([]*Folder, error)
}

// Folder represents a folder with name, description, creation time and a list of files
type Folder struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	CreatedAt   time.Time        `json:"created_at"`
	Files       map[string]*File `json:"files"`
}

// NewFolder creates a new Folder.
func NewFolder(name string, description string) *Folder {
	return &Folder{
		Name:        name,
		Description: description,
		CreatedAt:   time.Now(),
		Files:       make(map[string]*File),
	}
}
