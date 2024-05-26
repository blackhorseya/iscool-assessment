package vfs

import (
	"fmt"
	"regexp"
	"time"
)

// Constants for input validation
const (
	MaxInputLength = 255
	ValidChars     = "^[a-zA-Z0-9_-]+$"
)

// File represents a file in the virtual filesystem.
type File struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

// NewFile creates a new File.
func NewFile(name, description string) (*File, error) {
	if len(name) == 0 || len(name) > MaxInputLength {
		return nil, fmt.Errorf("file name length must be between 1 and %d characters", MaxInputLength)
	}
	if match, _ := regexp.MatchString(ValidChars, name); !match {
		return nil, fmt.Errorf("file name contains invalid characters")
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
	if len(name) == 0 || len(name) > MaxInputLength {
		return nil, fmt.Errorf("folder name length must be between 1 and %d characters", MaxInputLength)
	}
	if match, _ := regexp.MatchString(ValidChars, name); !match {
		return nil, fmt.Errorf("folder name contains invalid characters")
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
	if len(username) == 0 || len(username) > MaxInputLength {
		return nil, fmt.Errorf("username length must be between 1 and %d characters", MaxInputLength)
	}
	if match, _ := regexp.MatchString(ValidChars, username); !match {
		return nil, fmt.Errorf("username contains invalid characters")
	}

	return &User{
		Username: username,
		Folders:  make(map[string]*Folder),
	}, nil
}
