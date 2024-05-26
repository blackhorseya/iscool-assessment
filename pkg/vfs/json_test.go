package vfs

import (
	"os"
	"testing"
)

func TestJsonFile_Save(t *testing.T) {
	type fields struct {
		Users    map[string]*User
		filePath string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "save successfully",
			fields: fields{
				Users: map[string]*User{
					"user1": NewUser("user1"),
				},
				filePath: "test.json",
			},
			wantErr: false,
		},
		{
			name: "save with invalid file path",
			fields: fields{
				Users: map[string]*User{
					"user1": NewUser("user1"),
				},
				filePath: "/invalid/path/test.json",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vfs := &jsonFile{
				Users:    tt.fields.Users,
				filePath: tt.fields.filePath,
			}
			if err := vfs.Save(); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
			// Clean up
			if !tt.wantErr {
				_ = os.Remove(tt.fields.filePath)
			}
		})
	}
}

func TestJsonFile_Load(t *testing.T) {
	type fields struct {
		Users    map[string]*User
		filePath string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "load from existing file",
			fields: fields{
				Users: map[string]*User{
					"user1": NewUser("user1"),
				},
				filePath: "test.json",
			},
			wantErr: false,
		},
		{
			name: "load from non-existing file",
			fields: fields{
				Users: map[string]*User{
					"user1": NewUser("user1"),
				},
				filePath: "non_existing.json",
			},
			wantErr: false, // it should not error out, but return nil
		},
		{
			name: "load from file with invalid json",
			fields: fields{
				Users: map[string]*User{
					"user1": NewUser("user1"),
				},
				filePath: "invalid.json",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vfs := &jsonFile{
				Users:    tt.fields.Users,
				filePath: tt.fields.filePath,
			}
			// Save to file first to ensure it exists
			if !tt.wantErr && tt.name != "load from file with invalid json" {
				_ = vfs.Save()
			}
			if err := vfs.Load(); (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
			}
			// Clean up
			if !tt.wantErr {
				_ = os.Remove(tt.fields.filePath)
			}
		})
	}
}
