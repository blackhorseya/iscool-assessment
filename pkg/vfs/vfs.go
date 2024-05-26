package vfs

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

const orderAsc = "asc"

var _ UserManager = &VFS{}
var _ FolderManager = &VFS{}
var _ FileManager = &VFS{}

// VFS represents the entire file system with user management
type VFS struct {
	Users map[string]*User `json:"users"`
}

// NewVFS creates a new VFS
func NewVFS() *VFS {
	return &VFS{
		Users: make(map[string]*User),
	}
}

func (vfs *VFS) CreateFile(username, foldername, filename, description string) error {
	user, exists := vfs.Users[username]
	if !exists {
		return fmt.Errorf("the %s doesn't exist", username)
	}
	folder, exists := user.Folders[foldername]
	if !exists {
		return fmt.Errorf("the %s doesn't exist", foldername)
	}
	if _, exists = folder.Files[filename]; exists {
		return fmt.Errorf("the %s has already existed", filename)
	}
	folder.Files[filename] = NewFile(filename, description)
	return nil
}

func (vfs *VFS) DeleteFile(username, foldername, filename string) error {
	user, exists := vfs.Users[username]
	if !exists {
		return fmt.Errorf("the %s doesn't exist", username)
	}
	folder, exists := user.Folders[foldername]
	if !exists {
		return fmt.Errorf("the %s doesn't exist", foldername)
	}
	if _, exists = folder.Files[filename]; !exists {
		return fmt.Errorf("the %s doesn't exist", filename)
	}
	delete(folder.Files, filename)
	return nil
}

func (vfs *VFS) ListFiles(username, foldername, sortBy string, order string) ([]*File, error) {
	user, exists := vfs.Users[username]
	if !exists {
		return nil, fmt.Errorf("the %s doesn't exist", username)
	}
	folder, exists := user.Folders[foldername]
	if !exists {
		return nil, fmt.Errorf("the %s doesn't exist", foldername)
	}
	var files []*File
	for _, file := range folder.Files {
		files = append(files, file)
	}
	sort.Slice(files, func(i, j int) bool {
		switch sortBy {
		case "name":
			if order == orderAsc {
				return files[i].Name < files[j].Name
			}
			return files[i].Name > files[j].Name
		case "created":
			if order == orderAsc {
				return files[i].CreatedAt.Before(files[j].CreatedAt)
			}
			return files[i].CreatedAt.After(files[j].CreatedAt)
		default:
			return files[i].Name < files[j].Name
		}
	})
	return files, nil
}

func (vfs *VFS) CreateFolder(username, foldername, description string) error {
	user, exists := vfs.Users[username]
	if !exists {
		return fmt.Errorf("the %s doesn't exist", username)
	}
	if _, exists = user.Folders[foldername]; exists {
		return fmt.Errorf("the %s has already existed", foldername)
	}
	user.Folders[foldername] = NewFolder(foldername, description)
	return nil
}

func (vfs *VFS) DeleteFolder(username, foldername string) error {
	user, exists := vfs.Users[username]
	if !exists {
		return fmt.Errorf("the %s doesn't exist", username)
	}
	if _, exists = user.Folders[foldername]; !exists {
		return fmt.Errorf("the %s doesn't exist", foldername)
	}
	delete(user.Folders, foldername)
	return nil
}

func (vfs *VFS) RenameFolder(username, foldername, newFoldername string) error {
	user, exists := vfs.Users[username]
	if !exists {
		return fmt.Errorf("the %s doesn't exist", username)
	}
	folder, exists := user.Folders[foldername]
	if !exists {
		return fmt.Errorf("the %s doesn't exist", foldername)
	}
	if _, exists = user.Folders[newFoldername]; exists {
		return fmt.Errorf("the %s has already existed", newFoldername)
	}
	delete(user.Folders, foldername)
	folder.Name = newFoldername
	user.Folders[newFoldername] = folder
	return nil
}

func (vfs *VFS) ListFolders(username string, sortBy string, order string) ([]*Folder, error) {
	user, exists := vfs.Users[username]
	if !exists {
		return nil, fmt.Errorf("the %s doesn't exist", username)
	}
	var folders []*Folder
	for _, folder := range user.Folders {
		folders = append(folders, folder)
	}
	sort.Slice(folders, func(i, j int) bool {
		switch sortBy {
		case "name":
			if order == orderAsc {
				return folders[i].Name < folders[j].Name
			}
			return folders[i].Name > folders[j].Name
		case "created":
			if order == orderAsc {
				return folders[i].CreatedAt.Before(folders[j].CreatedAt)
			}
			return folders[i].CreatedAt.After(folders[j].CreatedAt)
		default:
			return folders[i].Name < folders[j].Name
		}
	})
	return folders, nil
}

func (vfs *VFS) RegisterUser(username string) error {
	if _, exists := vfs.Users[username]; exists {
		return fmt.Errorf("the %s has already existed", username)
	}

	vfs.Users[username] = NewUser(username)

	return nil
}

func (vfs *VFS) DeleteUser(username string) error {
	if _, exists := vfs.Users[username]; !exists {
		return fmt.Errorf("the %s doesn't exist", username)
	}
	delete(vfs.Users, username)
	return nil
}

func (vfs *VFS) ListUsers() []string {
	var users []string
	for username := range vfs.Users {
		users = append(users, username)
	}
	return users
}

// SaveToFile saves the virtual filesystem to a file.
func (vfs *VFS) SaveToFile(filename string) error {
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

func (vfs *VFS) LoadFromFile(filename string) error {
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
