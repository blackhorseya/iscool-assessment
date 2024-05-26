package vfs

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
			vfs := &VFS{
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

func TestVirtualFileSystem_CreateFile(t *testing.T) {
	type fields struct {
		Users map[string]*User
	}
	type args struct {
		username    string
		foldername  string
		filename    string
		description string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "create file successfully",
			fields: fields{Users: map[string]*User{
				"user1": {
					Username: "user1",
					Folders:  map[string]*Folder{"folder1": NewFolder("folder1", "description1")},
				},
			}},
			args:    args{username: "user1", foldername: "folder1", filename: "file1", description: "description1"},
			wantErr: false,
		},
		{
			name: "create file with non-existing user",
			fields: fields{Users: map[string]*User{
				"user1": {
					Username: "user1",
					Folders:  map[string]*Folder{"folder1": NewFolder("folder1", "description1")},
				},
			}},
			args:    args{username: "user2", foldername: "folder1", filename: "file1", description: "description1"},
			wantErr: true,
		},
		{
			name: "create file with non-existing folder",
			fields: fields{Users: map[string]*User{
				"user1": {
					Username: "user1",
					Folders:  map[string]*Folder{"folder1": NewFolder("folder1", "description1")},
				},
			}},
			args:    args{username: "user1", foldername: "folder2", filename: "file1", description: "description1"},
			wantErr: true,
		},
		{
			name: "create file with existing file name",
			fields: fields{Users: map[string]*User{
				"user1": {
					Username: "user1",
					Folders: map[string]*Folder{"folder1": {
						Name:  "folder1",
						Files: map[string]*File{"file1": NewFile("file1", "description1")},
					}},
				},
			}},
			args:    args{username: "user1", foldername: "folder1", filename: "file1", description: "description1"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vfs := &VFS{
				Users: tt.fields.Users,
			}
			if err := vfs.CreateFile(
				tt.args.username,
				tt.args.foldername,
				tt.args.filename,
				tt.args.description,
			); (err != nil) != tt.wantErr {
				t.Errorf("CreateFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestVirtualFileSystem_DeleteFile(t *testing.T) {
	type fields struct {
		Users map[string]*User
	}
	type args struct {
		username   string
		foldername string
		filename   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "delete file successfully",
			fields: fields{Users: map[string]*User{
				"user1": {
					Username: "user1",
					Folders: map[string]*Folder{"folder1": {
						Name:  "folder1",
						Files: map[string]*File{"file1": NewFile("file1", "description1")},
					}},
				},
			}},
			args:    args{username: "user1", foldername: "folder1", filename: "file1"},
			wantErr: false,
		},
		{
			name: "delete file with non-existing user",
			fields: fields{Users: map[string]*User{
				"user1": {
					Username: "user1",
					Folders: map[string]*Folder{"folder1": {
						Name:  "folder1",
						Files: map[string]*File{"file1": NewFile("file1", "description1")},
					}},
				},
			}},
			args:    args{username: "user2", foldername: "folder1", filename: "file1"},
			wantErr: true,
		},
		{
			name: "delete file with non-existing folder",
			fields: fields{Users: map[string]*User{
				"user1": {
					Username: "user1",
					Folders: map[string]*Folder{"folder1": {
						Name:  "folder1",
						Files: map[string]*File{"file1": NewFile("file1", "description1")},
					}},
				},
			}},
			args:    args{username: "user1", foldername: "folder2", filename: "file1"},
			wantErr: true,
		},
		{
			name: "delete file with non-existing file name",
			fields: fields{Users: map[string]*User{
				"user1": {
					Username: "user1",
					Folders: map[string]*Folder{"folder1": {
						Name:  "folder1",
						Files: map[string]*File{"file1": NewFile("file1", "description1")},
					}},
				},
			}},
			args:    args{username: "user1", foldername: "folder1", filename: "file2"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vfs := &VFS{
				Users: tt.fields.Users,
			}
			if err := vfs.DeleteFile(tt.args.username, tt.args.foldername, tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("DeleteFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestVirtualFileSystem_ListFiles(t *testing.T) {
	type fields struct {
		Users map[string]*User
	}
	type args struct {
		username   string
		foldername string
		sortBy     string
		order      string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "list files successfully",
			fields: fields{Users: map[string]*User{
				"user1": {
					Username: "user1",
					Folders: map[string]*Folder{"folder1": {
						Name:  "folder1",
						Files: map[string]*File{"file1": NewFile("file1", "description1"), "file2": NewFile("file2", "description2")},
					}},
				},
			}},
			args:    args{username: "user1", foldername: "folder1", sortBy: "name", order: "asc"},
			wantErr: false,
		},
		{
			name: "list files with non-existing user",
			fields: fields{Users: map[string]*User{
				"user1": {
					Username: "user1",
					Folders: map[string]*Folder{"folder1": {
						Name:  "folder1",
						Files: map[string]*File{"file1": NewFile("file1", "description1"), "file2": NewFile("file2", "description2")},
					}},
				},
			}},
			args:    args{username: "user2", foldername: "folder1", sortBy: "name", order: "asc"},
			wantErr: true,
		},
		{
			name: "list files with non-existing folder",
			fields: fields{Users: map[string]*User{
				"user1": {
					Username: "user1",
					Folders: map[string]*Folder{"folder1": {
						Name:  "folder1",
						Files: map[string]*File{"file1": NewFile("file1", "description1"), "file2": NewFile("file2", "description2")},
					}},
				},
			}},
			args:    args{username: "user1", foldername: "folder2", sortBy: "name", order: "asc"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vfs := &VFS{
				Users: tt.fields.Users,
			}
			_, err := vfs.ListFiles(tt.args.username, tt.args.foldername, tt.args.sortBy, tt.args.order)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListFiles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
