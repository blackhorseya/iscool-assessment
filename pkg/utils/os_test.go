package utils

import (
	"os"
	"testing"
)

func TestCheckPathType(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
		mock func()
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
			path: "out/existing_file.json",
			want: "json",
			mock: func() {
				_ = EnsureDir("out")
				_, _ = os.Create("out/existing_file.json")
			},
		},
		{
			name: "Existing folder",
			path: "out/existing_folder",
			want: "folder",
			mock: func() {
				_ = os.MkdirAll("out/existing_folder", 0755)
			},
		},
		{
			name: "Non-existing non-json file",
			path: "non_existing_file.txt",
			want: "folder",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock()
			}

			if got := CheckPathType(tt.path); got != tt.want {
				t.Errorf("CheckPathType() = %v, want %v", got, tt.want)
			}
		})
	}

	_ = os.RemoveAll("out")
}

func TestEnsureDir(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{
			name:    "Creates non-existent directory",
			path:    "out/file.txt",
			wantErr: false,
		},
		{
			name:    "Does not create existing directory",
			path:    "out/file.txt",
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

	_ = os.RemoveAll("out")
}
