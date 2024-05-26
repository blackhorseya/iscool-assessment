package vfs

import (
	"os"
	"reflect"
	"testing"

	"github.com/blackhorseya/iscool-assessment/entity/model"
	"github.com/blackhorseya/iscool-assessment/entity/repo"
	"github.com/blackhorseya/iscool-assessment/internal/repo/folder"
	"github.com/blackhorseya/iscool-assessment/internal/repo/user"
	"github.com/blackhorseya/iscool-assessment/pkg/vfs"
	"github.com/stretchr/testify/suite"
)

const defaultPath = "out/data.json"

type suiteIntegration struct {
	suite.Suite

	users   repo.UserManager
	folders repo.FolderManager
	vfs     vfs.VirtualFileSystem
}

func (s *suiteIntegration) SetupTest() {
	users, err := user.NewJSONFile(defaultPath)
	s.Require().NoError(err)
	s.users = users

	folders, err := folder.NewJSONFile(defaultPath)
	s.Require().NoError(err)
	s.folders = folders

	s.vfs = New(s.users, s.folders)
}

func (s *suiteIntegration) TearDownTest() {
	_ = os.Remove(defaultPath)
}

func TestIntegration(t *testing.T) {
	suite.Run(t, new(suiteIntegration))
}

func (s *suiteIntegration) Test_impl_RegisterUser() {
	user1, _ := model.NewUser("validUsername")

	type args struct {
		username string
	}
	tests := []struct {
		name     string
		args     args
		wantItem *model.User
		wantErr  bool
	}{
		{
			name:     "register user with valid username",
			args:     args{username: user1.Username},
			wantItem: user1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			gotItem, err := s.vfs.RegisterUser(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotItem, tt.wantItem) {
				t.Errorf("RegisterUser() gotItem = %v, want %v", gotItem, tt.wantItem)
			}
		})
	}
}
