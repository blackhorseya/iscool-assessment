package model

import (
	"strings"
	"testing"
	"time"
)

func TestNewFolder(t *testing.T) {
	user := &User{}

	tests := []struct {
		name        string
		owner       *User
		folderName  string
		description string
		wantErr     bool
	}{
		{
			name:        "Valid input",
			owner:       user,
			folderName:  "valid-input",
			description: "description",
			wantErr:     false,
		},
		{
			name:        "Empty name",
			owner:       user,
			folderName:  "",
			description: "description",
			wantErr:     true,
		},
		{
			name:        "Name exceeding max length",
			owner:       user,
			folderName:  strings.Repeat("a", 256),
			description: "description",
			wantErr:     true,
		},
		{
			name:        "Invalid characters in name",
			owner:       user,
			folderName:  "invalid!",
			description: "description",
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			folder, err := NewFolder(tt.owner, tt.folderName, tt.description)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFolder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil && folder.CreatedAt.After(time.Now()) {
				t.Errorf("NewFolder() folder.CreatedAt is in the future")
			}
		})
	}
}
