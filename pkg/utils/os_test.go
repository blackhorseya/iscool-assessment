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
			name: "check type of non-existing json file",
			path: "non_existing_file.json",
			want: "json",
		},
		{
			name: "check type of non-existing folder",
			path: "non_existing_folder",
			want: "folder",
		},
		{
			name: "check type of existing json file",
			path: "testdata/existing_file.json",
			want: "json",
		},
		{
			name: "check type of existing folder",
			path: "testdata/existing_folder",
			want: "folder",
		},
		{
			name: "check type of path with error",
			path: "",
			want: "error",
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
