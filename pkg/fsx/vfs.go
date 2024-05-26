package fsx

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

var _ UserManager = &VirtualFileSystem{}

// VirtualFileSystem represents the entire file system with user management
type VirtualFileSystem struct {
	Users map[string]*User `json:"users"`
}

// NewVFS creates a new VirtualFileSystem
func NewVFS() *VirtualFileSystem {
	return &VirtualFileSystem{
		Users: make(map[string]*User),
	}
}

// SaveToFile saves the virtual filesystem to a file.
func (vfs *VirtualFileSystem) SaveToFile(filename string) error {
	data, err := json.MarshalIndent(vfs, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	return os.WriteFile(filename, data, 0600)
}

func (vfs *VirtualFileSystem) LoadFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}

		return fmt.Errorf("failed to read file: %w", err)
	}

	return json.Unmarshal(data, vfs)
}

func (vfs *VirtualFileSystem) RegisterUser(username string) error {
	if _, exists := vfs.Users[username]; exists {
		return errors.New("the username has already existed")
	}

	vfs.Users[username] = NewUser(username)

	return nil
}

func (vfs *VirtualFileSystem) DeleteUser(username string) error {
	// todo: 2024/5/26|sean|implement me
	panic("implement me")
}

func (vfs *VirtualFileSystem) ListUsers() []string {
	// todo: 2024/5/26|sean|implement me
	panic("implement me")
}
