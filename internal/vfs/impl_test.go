package vfs

import (
	"errors"
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

func (s *suiteTester) Test_impl_CreateFolder() {
	user1, _ := model.NewUser("validUsername")
	folder1, _ := model.NewFolder(user1, "validFoldername", "validDescription")

	type args struct {
		username    string
		foldername  string
		description string
		mock        func()
	}
	tests := []struct {
		name     string
		args     args
		wantItem *model.Folder
		wantErr  bool
	}{
		{
			name: "create folder with valid username and foldername",
			args: args{
				username:    "validUsername",
				foldername:  "validFoldername",
				description: "validDescription",
				mock: func() {
					s.users.EXPECT().GetByUsername(gomock.Any(), user1.Username).Return(user1, nil).Times(1)

					s.folders.EXPECT().Create(
						gomock.Any(),
						user1,
						folder1.Name,
						folder1.Description,
					).Return(folder1, nil).Times(1)
				},
			},
			wantItem: folder1,
			wantErr:  false,
		},
		{
			name: "create folder with empty username",
			args: args{
				username:    "",
				foldername:  "validFoldername",
				description: "validDescription",
				mock: func() {
					s.users.EXPECT().GetByUsername(gomock.Any(), "").Return(nil, errors.New("empty user")).Times(1)
				},
			},
			wantItem: nil,
			wantErr:  true,
		},
		{
			name: "create folder with empty foldername",
			args: args{
				username:    "validUsername",
				foldername:  "",
				description: "validDescription",
				mock: func() {
					s.users.EXPECT().GetByUsername(gomock.Any(), "validUsername").Return(user1, nil).Times(1)

					s.folders.EXPECT().Create(
						gomock.Any(),
						gomock.Any(),
						gomock.Any(),
						gomock.Any(),
					).Return(nil, fmt.Errorf("foldername cannot be empty"))
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

			gotItem, err := s.vfs.CreateFolder(tt.args.username, tt.args.foldername, tt.args.description)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateFolder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotItem, tt.wantItem) {
				t.Errorf("CreateFolder() gotItem = %v, want %v", gotItem, tt.wantItem)
			}
		})
	}
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
