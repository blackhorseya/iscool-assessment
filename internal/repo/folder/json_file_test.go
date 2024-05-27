package folder

import (
	"context"
	"os"
	"reflect"
	"sync"
	"testing"

	"github.com/blackhorseya/iscool-assessment/entity/model"
)

func Test_jsonFile_GetByName(t *testing.T) {
	type fields struct {
		users map[string]*model.User
		path  string
	}
	type args struct {
		ctx        context.Context
		owner      *model.User
		foldername string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantItem *model.Folder
		wantErr  bool
	}{
		{
			name: "Valid folder retrieval",
			fields: fields{
				users: map[string]*model.User{
					"user1": {
						Username: "user1",
						Folders: map[string]*model.Folder{
							"folder1": {
								Name: "folder1",
							},
						},
					},
				},
				path: "out/vfs.json",
			},
			args: args{
				ctx:        context.Background(),
				owner:      &model.User{Username: "user1"},
				foldername: "folder1",
			},
			wantItem: &model.Folder{Name: "folder1"},
			wantErr:  false,
		},
		{
			name: "User not found",
			fields: fields{
				users: map[string]*model.User{},
				path:  "out/vfs.json",
			},
			args: args{
				ctx:        context.Background(),
				owner:      &model.User{Username: "user1"},
				foldername: "folder1",
			},
			wantItem: nil,
			wantErr:  true,
		},
		{
			name: "Folder not found",
			fields: fields{
				users: map[string]*model.User{
					"user1": {
						Username: "user1",
						Folders:  map[string]*model.Folder{},
					},
				},
				path: "out/vfs.json",
			},
			args: args{
				ctx:        context.Background(),
				owner:      &model.User{Username: "user1"},
				foldername: "nonexistentfolder",
			},
			wantItem: nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &jsonFile{
				Mutex: sync.Mutex{},
				users: tt.fields.users,
				path:  tt.fields.path,
			}
			gotItem, err := i.GetByName(tt.args.ctx, tt.args.owner, tt.args.foldername)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotItem, tt.wantItem) {
				t.Errorf("GetByName() gotItem = %v, want %v", gotItem, tt.wantItem)
			}
		})
	}
}

func Test_jsonFile_Create(t *testing.T) {
	user1, _ := model.NewUser("validUsername")
	folder1, _ := model.NewFolder(user1, "validFoldername", "validDescription")
	folder2, _ := model.NewFolder(user1, "folder2", "validDescription")

	type args struct {
		owner       *model.User
		foldername  string
		description string
	}
	tests := []struct {
		name    string
		args    args
		want    *model.Folder
		wantErr bool
	}{
		{
			name: "create folder with valid username and foldername",
			args: args{
				owner:       user1,
				foldername:  "folder2",
				description: "validDescription",
			},
			want:    folder2,
			wantErr: false,
		},
		{
			name: "create folder with non-existing username",
			args: args{
				owner:       &model.User{Username: "nonExistingUsername"},
				foldername:  "validFoldername",
				description: "validDescription",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "create folder with existing foldername",
			args: args{
				owner:       user1,
				foldername:  "validFoldername",
				description: "validDescription",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &jsonFile{
				Mutex: sync.Mutex{},
				users: map[string]*model.User{
					user1.Username: user1,
				},
				path: "out/vfs.json",
			}
			user1.Folders[folder1.Name] = folder1

			got, err := i.Create(context.Background(), tt.args.owner, tt.args.foldername, tt.args.description)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil && got.Name != tt.want.Name {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}

			// Clean up
			_ = os.Remove("out/vfs.json")
		})
	}
}

func Test_jsonFile_Delete(t *testing.T) {
	user1, _ := model.NewUser("validUsername")
	folder1, _ := model.NewFolder(user1, "validFoldername", "validDescription")

	type args struct {
		owner      *model.User
		foldername string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "delete folder with valid username and foldername",
			args: args{
				owner:      user1,
				foldername: "validFoldername",
			},
			wantErr: false,
		},
		{
			name: "delete folder with non-existing username",
			args: args{
				owner:      &model.User{Username: "nonExistingUsername"},
				foldername: "validFoldername",
			},
			wantErr: true,
		},
		{
			name: "delete folder with non-existing foldername",
			args: args{
				owner:      user1,
				foldername: "nonExistingFoldername",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &jsonFile{
				Mutex: sync.Mutex{},
				users: map[string]*model.User{
					user1.Username: user1,
				},
				path: "out/vfs.json",
			}
			user1.Folders[folder1.Name] = folder1

			err := i.Delete(context.Background(), tt.args.owner, tt.args.foldername)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Clean up
			_ = os.Remove("out/vfs.json")
		})
	}
}

func Test_jsonFile_Rename(t *testing.T) {
	user1, _ := model.NewUser("validUsername")
	folder1, _ := model.NewFolder(user1, "validFoldername", "validDescription")
	folder2, _ := model.NewFolder(user1, "newValidFoldername", "validDescription")

	type args struct {
		owner         *model.User
		foldername    string
		newFoldername string
	}
	tests := []struct {
		name    string
		args    args
		want    *model.Folder
		wantErr bool
	}{
		{
			name: "rename folder with valid username and foldername",
			args: args{
				owner:         user1,
				foldername:    "validFoldername",
				newFoldername: "newValidFoldername",
			},
			want:    folder2,
			wantErr: false,
		},
		{
			name: "rename folder with non-existing username",
			args: args{
				owner:         &model.User{Username: "nonExistingUsername"},
				foldername:    "validFoldername",
				newFoldername: "newValidFoldername",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "rename folder with non-existing foldername",
			args: args{
				owner:         user1,
				foldername:    "nonExistingFoldername",
				newFoldername: "newValidFoldername",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "rename folder with existing new foldername",
			args: args{
				owner:         user1,
				foldername:    "validFoldername",
				newFoldername: "validFoldername",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &jsonFile{
				Mutex: sync.Mutex{},
				users: map[string]*model.User{
					user1.Username: user1,
				},
				path: "out/vfs.json",
			}
			user1.Folders[folder1.Name] = folder1

			got, err := i.Rename(context.Background(), tt.args.owner, tt.args.foldername, tt.args.newFoldername)
			if (err != nil) != tt.wantErr {
				t.Errorf("Rename() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil && got.Name != tt.want.Name {
				t.Errorf("Rename() got = %v, want %v", got, tt.want)
			}

			// Clean up
			_ = os.Remove("out/vfs.json")
		})
	}
}

func Test_jsonFile_List(t *testing.T) {
	user1, _ := model.NewUser("validUsername")
	folder1, _ := model.NewFolder(user1, "folder1", "validDescription")
	folder2, _ := model.NewFolder(user1, "folder2", "validDescription")

	type args struct {
		owner  *model.User
		sortBy string
		order  string
	}
	tests := []struct {
		name    string
		args    args
		want    []*model.Folder
		wantErr bool
	}{
		{
			name: "list folders with valid username",
			args: args{
				owner:  user1,
				sortBy: "name",
				order:  "asc",
			},
			want:    []*model.Folder{folder1, folder2},
			wantErr: false,
		},
		{
			name: "list folders with non-existing username",
			args: args{
				owner:  &model.User{Username: "nonExistingUsername"},
				sortBy: "name",
				order:  "asc",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "list folders with valid username and invalid sort field",
			args: args{
				owner:  user1,
				sortBy: "invalidField",
				order:  "asc",
			},
			want:    []*model.Folder{folder1, folder2},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &jsonFile{
				Mutex: sync.Mutex{},
				users: map[string]*model.User{
					user1.Username: user1,
				},
				path: "out/vfs.json",
			}
			user1.Folders[folder1.Name] = folder1
			user1.Folders[folder2.Name] = folder2

			got, err := i.List(context.Background(), tt.args.owner, tt.args.sortBy, tt.args.order)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List() got = %v, want %v", got, tt.want)
			}

			// Clean up
			_ = os.Remove("out/vfs.json")
		})
	}
}
