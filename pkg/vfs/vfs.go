package vfs

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

var _ UserManager = &VirtualFileSystem{}
var _ FolderManager = &VirtualFileSystem{}

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

func (vfs *VirtualFileSystem) CreateFolder(username, foldername, description string) error {
	user, exists := vfs.Users[username]
	if !exists {
		return errors.New("the username doesn't exist")
	}
	if _, exists = user.Folders[foldername]; exists {
		return errors.New("the foldername has already existed")
	}
	user.Folders[foldername] = NewFolder(foldername, description)
	return nil
}

func (vfs *VirtualFileSystem) DeleteFolder(username, foldername string) error {
	user, exists := vfs.Users[username]
	if !exists {
		return errors.New("the username doesn't exist")
	}
	if _, exists = user.Folders[foldername]; !exists {
		return errors.New("the foldername doesn't exist")
	}
	delete(user.Folders, foldername)
	return nil
}

func (vfs *VirtualFileSystem) RenameFolder(username, foldername, newFoldername string) error {
	user, exists := vfs.Users[username]
	if !exists {
		return errors.New("the username doesn't exist")
	}
	folder, exists := user.Folders[foldername]
	if !exists {
		return errors.New("the foldername doesn't exist")
	}
	if _, exists = user.Folders[newFoldername]; exists {
		return errors.New("the new foldername already exists")
	}
	delete(user.Folders, foldername)
	folder.Name = newFoldername
	user.Folders[newFoldername] = folder
	return nil
}

func (vfs *VirtualFileSystem) ListFolders(username string, sortBy string, order string) ([]*Folder, error) {
	user, exists := vfs.Users[username]
	if !exists {
		return nil, errors.New("the username doesn't exist")
	}
	var folders []*Folder
	for _, folder := range user.Folders {
		folders = append(folders, folder)
	}
	sort.Slice(folders, func(i, j int) bool {
		switch sortBy {
		case "name":
			if order == "asc" {
				return folders[i].Name < folders[j].Name
			}
			return folders[i].Name > folders[j].Name
		case "created":
			if order == "asc" {
				return folders[i].CreatedAt.Before(folders[j].CreatedAt)
			}
			return folders[i].CreatedAt.After(folders[j].CreatedAt)
		default:
			return folders[i].Name < folders[j].Name
		}
	})
	return folders, nil
}

func (vfs *VirtualFileSystem) RegisterUser(username string) error {
	if _, exists := vfs.Users[username]; exists {
		return errors.New("the username has already existed")
	}

	vfs.Users[username] = NewUser(username)

	return nil
}

func (vfs *VirtualFileSystem) DeleteUser(username string) error {
	if _, exists := vfs.Users[username]; !exists {
		return errors.New("the username doesn't exist")
	}
	delete(vfs.Users, username)
	return nil
}

func (vfs *VirtualFileSystem) ListUsers() []string {
	var users []string
	for username := range vfs.Users {
		users = append(users, username)
	}
	return users
}

// SaveToFile saves the virtual filesystem to a file.
func (vfs *VirtualFileSystem) SaveToFile(filename string) error {
	// Ensure the directory exists
	if err := ensureDir(filename); err != nil {
		return fmt.Errorf("failed to ensure directory: %w", err)
	}

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

// ensureDir checks if the directory for the file exists and creates it if not
func ensureDir(fileName string) error {
	dir := filepath.Dir(fileName)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}
