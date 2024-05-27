package utils

import (
	"testing"
)

func TestCheckPathType(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{
			name: "Empty path",
			path: "",
			want: "error",
		},
		{
			name: "Non-existing json file",
			path: "non_existing_file.json",
			want: "json",
		},
		{
			name: "Non-existing folder",
			path: "non_existing_folder",
			want: "folder",
		},
		{
			name: "Existing json file",
			path: "testdata/existing_file.json",
			want: "json",
		},
		{
			name: "Existing folder",
			path: "testdata/existing_folder",
			want: "folder",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPathType(tt.path); got != tt.want {
				t.Errorf("CheckPathType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnsureDir(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{
			name:    "Creates non-existent directory",
			path:    "test_dir/file.txt",
			wantErr: false,
		},
		{
			name:    "Does not create existing directory",
			path:    "test_dir/file.txt",
			wantErr: false,
		},
		{
			name:    "Returns error for invalid path",
			path:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := EnsureDir(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("EnsureDir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
