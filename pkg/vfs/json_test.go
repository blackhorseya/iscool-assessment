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
