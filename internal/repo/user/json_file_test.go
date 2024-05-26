package user

import (
	"testing"

	"github.com/blackhorseya/iscool-assessment/entity/model"
)

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
				path:  "testdata/valid.json",
			},
			wantErr: false,
		},
		{
			name: "load from non-existing file",
			fields: fields{
				users: make(map[string]*model.User),
				path:  "testdata/non-existing.json",
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
		})
	}
}
