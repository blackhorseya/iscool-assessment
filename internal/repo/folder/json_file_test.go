package folder

import (
	"context"
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
