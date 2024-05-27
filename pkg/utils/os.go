package utils

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

// CheckPathType checks the type of the path
func CheckPathType(path string) string {
	if path == "" {
		return "error"
	}

	fileInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			if strings.HasSuffix(path, ".json") {
				return "json"
			}
			return "folder"
		}
		return "error"
	}

	if fileInfo.IsDir() {
		return "folder"
	}

	if strings.HasSuffix(path, ".json") {
		return "json"
	}

	return "file"
}

// EnsureDir checks if the directory for the file exists and creates it if not
func EnsureDir(fileName string) error {
	if fileName == "" {
		return errors.New("invalid path")
	}

	dir := filepath.Dir(fileName)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}
