package utils

import (
	"os"
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
