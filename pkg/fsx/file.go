//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package fsx

import (
	"time"
)

// FileManager defines the interface for file management
type FileManager interface {
	CreateFile(username, foldername, filename, description string) error
	DeleteFile(username, foldername, filename string) error
	ListFiles(username, foldername, sortBy string, order string) ([]*File, error)
}

// File represents a file in the virtual filesystem.
type File struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
