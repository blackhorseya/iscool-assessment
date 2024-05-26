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
		return fmt.Errorf("failed to read file: %w", err)
	}

	return json.Unmarshal(data, vfs)
}
