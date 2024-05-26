package vfs

import (
	"encoding/json"
	"fmt"
	"os"
)

type jsonFile struct {
	Users map[string]*User `json:"users"`

	filePath string // Path to the JSON file
}

// NewJSONFile creates a new JSON file system.
func NewJSONFile(filePath string) VirtualFileSystem {
	return &jsonFile{
		Users:    make(map[string]*User),
		filePath: filePath,
	}
}

func (vfs *jsonFile) Save() error {
	// Ensure the directory exists
	if err := ensureDir(vfs.filePath); err != nil {
		return fmt.Errorf("failed to ensure directory: %w", err)
	}

	data, err := json.MarshalIndent(vfs, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	return os.WriteFile(vfs.filePath, data, 0600)
}

func (vfs *jsonFile) Load() error {
	data, err := os.ReadFile(vfs.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}

		return fmt.Errorf("failed to read file: %w", err)
	}

	return json.Unmarshal(data, vfs)
}

func (vfs *jsonFile) GetByUsername(username string) (*User, error) {
	user, exists := vfs.Users[username]
	if !exists {
		return nil, fmt.Errorf("the %s doesn't exist", username)
	}

	return user, nil
}

func (vfs *jsonFile) RegisterUser(username string) error {
	if _, exists := vfs.Users[username]; exists {
		return fmt.Errorf("the %s has already existed", username)
	}

	vfs.Users[username] = NewUser(username)

	return nil
}

func (vfs *jsonFile) DeleteUser(username string) error {
	// TODO implement me
	panic("implement me")
}

func (vfs *jsonFile) ListUsers() []string {
	// TODO implement me
	panic("implement me")
}

func (vfs *jsonFile) CreateFolder(username, foldername, description string) error {
	// TODO implement me
	panic("implement me")
}

func (vfs *jsonFile) DeleteFolder(username, foldername string) error {
	// TODO implement me
	panic("implement me")
}

func (vfs *jsonFile) RenameFolder(username, foldername, newFoldername string) error {
	// TODO implement me
	panic("implement me")
}

func (vfs *jsonFile) ListFolders(username string, sortBy string, order string) ([]*Folder, error) {
	// TODO implement me
	panic("implement me")
}

func (vfs *jsonFile) CreateFile(username, foldername, filename, description string) error {
	// TODO implement me
	panic("implement me")
}

func (vfs *jsonFile) DeleteFile(username, foldername, filename string) error {
	// TODO implement me
	panic("implement me")
}

func (vfs *jsonFile) ListFiles(username, foldername, sortBy string, order string) ([]*File, error) {
	// TODO implement me
	panic("implement me")
}
