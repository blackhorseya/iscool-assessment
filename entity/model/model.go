package model

import (
	"time"
)

// File represents a file in the virtual filesystem.
type File struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

// NewFile creates a new File.
func NewFile(name, description string) (*File, error) {
	err := ValidateInput(name)
	if err != nil {
		return nil, err
	}

	return &File{
		Name:        name,
		Description: description,
		CreatedAt:   time.Now(),
	}, nil
}

// Folder represents a folder with name, description, creation time and a list of files.
type Folder struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	CreatedAt   time.Time        `json:"created_at"`
	Files       map[string]*File `json:"files"`
}

// NewFolder creates a new Folder.
func NewFolder(name, description string) (*Folder, error) {
	err := ValidateInput(name)
	if err != nil {
		return nil, err
	}

	return &Folder{
		Name:        name,
		Description: description,
		CreatedAt:   time.Now(),
		Files:       make(map[string]*File),
	}, nil
}

// User represents a user with username and a list of folders.
type User struct {
	Username string             `json:"username"`
	Folders  map[string]*Folder `json:"folders"`
}

// NewUser creates a new User.
func NewUser(username string) (*User, error) {
	err := ValidateInput(username)
	if err != nil {
		return nil, err
	}

	return &User{
		Username: username,
		Folders:  make(map[string]*Folder),
	}, nil
}
