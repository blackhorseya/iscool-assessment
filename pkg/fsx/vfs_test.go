package fsx

import (
	"os"
	"testing"
)

func TestVirtualFileSystem_SaveToFile(t *testing.T) {
	type fields struct {
		Users map[string]*User
	}
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "save to file successfully",
			fields:  fields{Users: make(map[string]*User)},
			args:    args{filename: "vfs.json"},
			wantErr: false,
		},
		{
			name:    "save to file successfully with path",
			fields:  fields{Users: make(map[string]*User)},
			args:    args{filename: "out/vfs.json"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vfs := &VirtualFileSystem{
				Users: tt.fields.Users,
			}
			if err := vfs.SaveToFile(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("SaveToFile() error = %v, wantErr %v", err, tt.wantErr)
			}

			// clean up
			_ = os.Remove(tt.args.filename)
		})
	}
}
