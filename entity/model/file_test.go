package model

import (
	"strings"
	"testing"
	"time"
)

func TestNewFile(t *testing.T) {
	user := &User{}
	folder := &Folder{}

	tests := []struct {
		name        string
		owner       *User
		folder      *Folder
		fileName    string
		description string
		wantErr     bool
	}{
		{
			name:        "Valid input",
			owner:       user,
			folder:      folder,
			fileName:    "valid-input",
			description: "description",
			wantErr:     false,
		},
		{
			name:        "Empty name",
			owner:       user,
			folder:      folder,
			fileName:    "",
			description: "description",
			wantErr:     true,
		},
		{
			name:        "Name exceeding max length",
			owner:       user,
			folder:      folder,
			fileName:    strings.Repeat("a", 256),
			description: "description",
			wantErr:     true,
		},
		{
			name:        "Invalid characters in name",
			owner:       user,
			folder:      folder,
			fileName:    "invalid!",
			description: "description",
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, err := NewFile(tt.owner, tt.folder, tt.fileName, tt.description)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil && file.CreatedAt.After(time.Now()) {
				t.Errorf("NewFile() file.CreatedAt is in the future")
			}
		})
	}
}
