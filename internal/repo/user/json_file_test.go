package user

import (
	"context"
	"os"
	"testing"

	"github.com/blackhorseya/iscool-assessment/entity/model"
)

func Test_NewJSONFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "create new JSON file with valid path",
			args: args{
				path: "out/valid.json",
			},
			wantErr: false,
		},
		{
			name: "create new JSON file with invalid path",
			args: args{
				path: "/invalid/path.json",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewJSONFile(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewJSONFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_jsonFile_Save(t *testing.T) {
	type fields struct {
		users map[string]*model.User
		path  string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "save to valid path",
			fields: fields{
				users: map[string]*model.User{"user1": {Username: "user1"}},
				path:  "out/valid.json",
			},
			wantErr: false,
		},
		{
			name: "save to invalid path",
			fields: fields{
				users: map[string]*model.User{"user1": {Username: "user1"}},
				path:  "/invalid/path.json",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &jsonFile{
				users: tt.fields.users,
				path:  tt.fields.path,
			}
			if err := i.Save(); (err != nil) != tt.wantErr {
				t.Errorf("jsonFile.Save() error = %v, wantErr %v", err, tt.wantErr)
			}

			// clean up
			_ = os.Remove(tt.fields.path)
		})
	}
}

func Test_jsonFile_Load(t *testing.T) {
	type fields struct {
		users map[string]*model.User
		path  string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "load from existing and valid json file",
			fields: fields{
				users: make(map[string]*model.User),
				path:  "out/valid.json",
			},
			wantErr: false,
		},
		{
			name: "load from non-existing file",
			fields: fields{
				users: make(map[string]*model.User),
				path:  "out/non-existing.json",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &jsonFile{
				users: tt.fields.users,
				path:  tt.fields.path,
			}
			if err := i.Load(); (err != nil) != tt.wantErr {
				t.Errorf("jsonFile.Load() error = %v, wantErr %v", err, tt.wantErr)
			}

			// clean up
			_ = os.Remove(tt.fields.path)
		})
	}
}

func Test_jsonFile_Register(t *testing.T) {
	type fields struct {
		users map[string]*model.User
		path  string
	}
	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "register new user",
			fields: fields{
				users: make(map[string]*model.User),
				path:  "out/valid.json",
			},
			args: args{
				ctx:      context.Background(),
				username: "newUser",
			},
			wantErr: false,
		},
		{
			name: "register existing user",
			fields: fields{
				users: map[string]*model.User{"existingUser": {Username: "existingUser"}},
				path:  "out/valid.json",
			},
			args: args{
				ctx:      context.Background(),
				username: "existingUser",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &jsonFile{
				users: tt.fields.users,
				path:  tt.fields.path,
			}
			_, err := i.Register(tt.args.ctx, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("jsonFile.Register() error = %v, wantErr %v", err, tt.wantErr)
			}

			// clean up
			_ = os.Remove(tt.fields.path)
		})
	}
}

func Test_jsonFile_GetByUsername(t *testing.T) {
	type fields struct {
		users map[string]*model.User
		path  string
	}
	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "get by existing username",
			fields: fields{
				users: map[string]*model.User{"existingUser": {Username: "existingUser"}},
				path:  "out/valid.json",
			},
			args: args{
				ctx:      context.Background(),
				username: "existingUser",
			},
			wantErr: false,
		},
		{
			name: "get by non-existing username",
			fields: fields{
				users: map[string]*model.User{"existingUser": {Username: "existingUser"}},
				path:  "out/valid.json",
			},
			args: args{
				ctx:      context.Background(),
				username: "nonExistingUser",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &jsonFile{
				users: tt.fields.users,
				path:  tt.fields.path,
			}
			_, err := i.GetByUsername(tt.args.ctx, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("jsonFile.GetByUsername() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
