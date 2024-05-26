package vfs

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/blackhorseya/iscool-assessment/entity/model"
	"github.com/blackhorseya/iscool-assessment/entity/repo"
	"github.com/blackhorseya/iscool-assessment/pkg/vfs"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type suiteTester struct {
	suite.Suite

	ctrl    *gomock.Controller
	users   *repo.MockUserManager
	folders *repo.MockFolderManager
	vfs     vfs.VirtualFileSystem
}

func (s *suiteTester) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.users = repo.NewMockUserManager(s.ctrl)
	s.folders = repo.NewMockFolderManager(s.ctrl)
	s.vfs = New(s.users, s.folders)
}

func (s *suiteTester) TearDownTest() {
	s.ctrl.Finish()
}

func TestAll(t *testing.T) {
	suite.Run(t, new(suiteTester))
}

func (s *suiteTester) Test_impl_RegisterUser() {
	type args struct {
		username string
		mock     func()
	}
	tests := []struct {
		name     string
		args     args
		wantItem *model.User
		wantErr  bool
	}{
		{
			name: "register user with valid username",
			args: args{
				username: "validUsername",
				mock: func() {
					user := &model.User{Username: "validUsername"}
					s.users.EXPECT().Register(gomock.Any(), user.Username).Return(user, nil).Times(1)
				},
			},
			wantItem: &model.User{Username: "validUsername"},
			wantErr:  false,
		},
		{
			name: "register user with empty username",
			args: args{
				username: "",
				mock: func() {
					s.users.EXPECT().Register(gomock.Any(), gomock.Any()).Return(
						nil,
						fmt.Errorf("username cannot be empty"),
					).Times(1)
				},
			},
			wantItem: nil,
			wantErr:  true,
		},
		{
			name: "register user with username already exists",
			args: args{
				username: "existingUsername",
				mock: func() {
					s.users.EXPECT().Register(gomock.Any(), gomock.Any()).Return(
						nil,
						fmt.Errorf("username already exists"),
					).Times(1)
				},
			},
			wantItem: nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

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
