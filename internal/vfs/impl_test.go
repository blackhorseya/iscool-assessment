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

func (s *suiteTester) Test_impl_DeleteFolder() {
	user1, _ := model.NewUser("validUsername")

	type args struct {
		username   string
		foldername string
		mock       func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "delete folder with valid username and foldername",
			args: args{
				username:   "validUsername",
				foldername: "validFoldername",
				mock: func() {
					s.users.EXPECT().GetByUsername(gomock.Any(), user1.Username).Return(user1, nil).Times(1)
					s.folders.EXPECT().Delete(gomock.Any(), user1, "validFoldername").Return(nil).Times(1)
				},
			},
			wantErr: false,
		},
		{
			name: "delete folder with empty username",
			args: args{
				username:   "",
				foldername: "validFoldername",
				mock: func() {
					s.users.EXPECT().GetByUsername(gomock.Any(), "").Return(nil, errors.New("empty user")).Times(1)
				},
			},
			wantErr: true,
		},
		{
			name: "delete folder with empty foldername",
			args: args{
				username:   "validUsername",
				foldername: "",
				mock: func() {
					s.users.EXPECT().GetByUsername(gomock.Any(), "validUsername").Return(user1, nil).Times(1)
					s.folders.EXPECT().Delete(gomock.Any(), user1, "").Return(fmt.Errorf("foldername cannot be empty")).Times(1)
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			err := s.vfs.DeleteFolder(tt.args.username, tt.args.foldername)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteFolder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func (s *suiteTester) Test_impl_ListFolders() {
	user1, _ := model.NewUser("validUsername")
	folders := []*model.Folder{
		{Name: "Folder1", Description: "Description1"},
		{Name: "Folder2", Description: "Description2"},
	}

	type args struct {
		username string
		sortBy   string
		order    string
		mock     func()
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
				username: "validUsername",
				sortBy:   "name",
				order:    "asc",
				mock: func() {
					s.users.EXPECT().GetByUsername(gomock.Any(), user1.Username).Return(user1, nil).Times(1)
					s.folders.EXPECT().List(gomock.Any(), user1, "name", "asc").Return(folders, nil).Times(1)
				},
			},
			want:    folders,
			wantErr: false,
		},
		{
			name: "list folders with empty username",
			args: args{
				username: "",
				sortBy:   "name",
				order:    "asc",
				mock: func() {
					s.users.EXPECT().GetByUsername(gomock.Any(), "").Return(nil, errors.New("empty user")).Times(1)
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "list folders with invalid sort field",
			args: args{
				username: "validUsername",
				sortBy:   "invalid",
				order:    "asc",
				mock: func() {
					s.users.EXPECT().GetByUsername(gomock.Any(), user1.Username).Return(user1, nil).Times(1)
					s.folders.EXPECT().List(
						gomock.Any(),
						user1,
						"invalid",
						"asc",
					).Return(nil, fmt.Errorf("invalid sort field")).Times(1)
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			got, err := s.vfs.ListFolders(tt.args.username, tt.args.sortBy, tt.args.order)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListFolders() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListFolders() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func (s *suiteTester) Test_impl_RenameFolder() {
	user1, _ := model.NewUser("validUsername")
	folder1, _ := model.NewFolder(user1, "validFoldername", "validDescription")

	type args struct {
		username      string
		foldername    string
		newFoldername string
		mock          func()
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
				username:      "validUsername",
				foldername:    "validFoldername",
				newFoldername: "newFoldername",
				mock: func() {
					s.users.EXPECT().GetByUsername(gomock.Any(), user1.Username).Return(user1, nil).Times(1)
					s.folders.EXPECT().Rename(gomock.Any(), user1, folder1.Name, "newFoldername").Return(folder1, nil).Times(1)
				},
			},
			want:    folder1,
			wantErr: false,
		},
		{
			name: "rename folder with empty username",
			args: args{
				username:      "",
				foldername:    "validFoldername",
				newFoldername: "newFoldername",
				mock: func() {
					s.users.EXPECT().GetByUsername(gomock.Any(), "").Return(nil, errors.New("empty user")).Times(1)
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "rename folder with empty foldername",
			args: args{
				username:      "validUsername",
				foldername:    "",
				newFoldername: "newFoldername",
				mock: func() {
					s.users.EXPECT().GetByUsername(gomock.Any(), "validUsername").Return(user1, nil).Times(1)
					s.folders.EXPECT().Rename(
						gomock.Any(),
						user1,
						"",
						"newFoldername",
					).Return(nil, fmt.Errorf("foldername cannot be empty")).Times(1)
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "rename folder with empty new foldername",
			args: args{
				username:      "validUsername",
				foldername:    "validFoldername",
				newFoldername: "",
				mock: func() {
					s.users.EXPECT().GetByUsername(gomock.Any(), "validUsername").Return(user1, nil).Times(1)
					s.folders.EXPECT().Rename(
						gomock.Any(),
						user1,
						"validFoldername",
						"",
					).Return(nil, fmt.Errorf("new foldername cannot be empty")).Times(1)
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			got, err := s.vfs.RenameFolder(tt.args.username, tt.args.foldername, tt.args.newFoldername)
			if (err != nil) != tt.wantErr {
				t.Errorf("RenameFolder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RenameFolder() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func (s *suiteTester) Test_impl_CreateFile() {
	user1, _ := model.NewUser("validUsername")
	folder1, _ := model.NewFolder(user1, "validFoldername", "validDescription")
	file1, _ := model.NewFile(user1, folder1, "validFilename", "validDescription")

	type args struct {
		username    string
		foldername  string
		filename    string
		description string
		mock        func()
	}
	tests := []struct {
		name    string
		args    args
		want    *model.File
		wantErr bool
	}{
		{
			name: "create file with valid username, foldername and filename",
			args: args{
				username:    "validUsername",
				foldername:  "validFoldername",
				filename:    "validFilename",
				description: "validDescription",
				mock: func() {
					s.users.EXPECT().GetByUsername(gomock.Any(), user1.Username).Return(user1, nil).Times(1)
					s.folders.EXPECT().GetByName(gomock.Any(), user1, folder1.Name).Return(folder1, nil).Times(1)
					s.folders.EXPECT().CreateFile(
						gomock.Any(),
						user1,
						folder1,
						file1.Name,
						file1.Description,
					).Return(file1, nil).Times(1)
				},
			},
			want:    file1,
			wantErr: false,
		},
		{
			name: "create file with empty username",
			args: args{
				username:    "",
				foldername:  "validFoldername",
				filename:    "validFilename",
				description: "validDescription",
				mock: func() {
					s.users.EXPECT().GetByUsername(gomock.Any(), "").Return(nil, errors.New("empty user")).Times(1)
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "create file with empty foldername",
			args: args{
				username:    "validUsername",
				foldername:  "",
				filename:    "validFilename",
				description: "validDescription",
				mock: func() {
					s.users.EXPECT().GetByUsername(gomock.Any(), "validUsername").Return(user1, nil).Times(1)
					s.folders.EXPECT().GetByName(
						gomock.Any(),
						user1,
						"",
					).Return(nil, fmt.Errorf("foldername cannot be empty")).Times(1)
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "create file with empty filename",
			args: args{
				username:    "validUsername",
				foldername:  "validFoldername",
				filename:    "",
				description: "validDescription",
				mock: func() {
					s.users.EXPECT().GetByUsername(gomock.Any(), "validUsername").Return(user1, nil).Times(1)
					s.folders.EXPECT().GetByName(gomock.Any(), user1, "validFoldername").Return(folder1, nil).Times(1)
					s.folders.EXPECT().CreateFile(
						gomock.Any(),
						user1,
						folder1,
						"",
						"validDescription",
					).Return(nil, fmt.Errorf("filename cannot be empty")).Times(1)
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			got, err := s.vfs.CreateFile(tt.args.username, tt.args.foldername, tt.args.filename, tt.args.description)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func (s *suiteTester) Test_impl_DeleteFile() {
	user1, _ := model.NewUser("validUsername")
	folder1, _ := model.NewFolder(user1, "validFoldername", "validDescription")
	file1, _ := model.NewFile(user1, folder1, "validFilename", "validDescription")

	type args struct {
		username   string
		foldername string
		filename   string
		mock       func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "delete file with valid username, foldername and filename",
			args: args{
				username:   "validUsername",
				foldername: "validFoldername",
				filename:   "validFilename",
				mock: func() {
					s.users.EXPECT().GetByUsername(gomock.Any(), user1.Username).Return(user1, nil).Times(1)
					s.folders.EXPECT().GetByName(gomock.Any(), user1, folder1.Name).Return(folder1, nil).Times(1)
					s.folders.EXPECT().DeleteFile(gomock.Any(), user1, folder1, file1.Name).Return(nil).Times(1)
				},
			},
			wantErr: false,
		},
		{
			name: "delete file with empty username",
			args: args{
				username:   "",
				foldername: "validFoldername",
				filename:   "validFilename",
				mock: func() {
					s.users.EXPECT().GetByUsername(gomock.Any(), "").Return(nil, errors.New("empty user")).Times(1)
				},
			},
			wantErr: true,
		},
		{
			name: "delete file with empty foldername",
			args: args{
				username:   "validUsername",
				foldername: "",
				filename:   "validFilename",
				mock: func() {
					s.users.EXPECT().GetByUsername(gomock.Any(), "validUsername").Return(user1, nil).Times(1)
					s.folders.EXPECT().GetByName(
						gomock.Any(),
						user1,
						"",
					).Return(nil, fmt.Errorf("foldername cannot be empty")).Times(1)
				},
			},
			wantErr: true,
		},
		{
			name: "delete file with empty filename",
			args: args{
				username:   "validUsername",
				foldername: "validFoldername",
				filename:   "",
				mock: func() {
					s.users.EXPECT().GetByUsername(gomock.Any(), "validUsername").Return(user1, nil).Times(1)
					s.folders.EXPECT().GetByName(gomock.Any(), user1, "validFoldername").Return(folder1, nil).Times(1)
					s.folders.EXPECT().DeleteFile(
						gomock.Any(),
						user1,
						folder1,
						"",
					).Return(fmt.Errorf("filename cannot be empty")).Times(1)
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			err := s.vfs.DeleteFile(tt.args.username, tt.args.foldername, tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func (s *suiteTester) Test_impl_ListFiles() {
	user1, _ := model.NewUser("validUsername")
	folder1, _ := model.NewFolder(user1, "validFoldername", "validDescription")
	files := []*model.File{
		{Name: "File1", Description: "Description1"},
		{Name: "File2", Description: "Description2"},
	}

	type args struct {
		username   string
		foldername string
		sortBy     string
		order      string
		mock       func()
	}
	tests := []struct {
		name    string
		args    args
		want    []*model.File
		wantErr bool
	}{
		{
			name: "list files with valid username and foldername",
			args: args{
				username:   "validUsername",
				foldername: "validFoldername",
				sortBy:     "name",
				order:      "asc",
				mock: func() {
					s.users.EXPECT().GetByUsername(gomock.Any(), user1.Username).Return(user1, nil).Times(1)
					s.folders.EXPECT().GetByName(gomock.Any(), user1, folder1.Name).Return(folder1, nil).Times(1)
					s.folders.EXPECT().ListFiles(gomock.Any(), user1, folder1, "name", "asc").Return(files, nil).Times(1)
				},
			},
			want:    files,
			wantErr: false,
		},
		{
			name: "list files with empty username",
			args: args{
				username:   "",
				foldername: "validFoldername",
				sortBy:     "name",
				order:      "asc",
				mock: func() {
					s.users.EXPECT().GetByUsername(gomock.Any(), "").Return(nil, errors.New("empty user")).Times(1)
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "list files with empty foldername",
			args: args{
				username:   "validUsername",
				foldername: "",
				sortBy:     "name",
				order:      "asc",
				mock: func() {
					s.users.EXPECT().GetByUsername(gomock.Any(), "validUsername").Return(user1, nil).Times(1)
					s.folders.EXPECT().GetByName(
						gomock.Any(),
						user1,
						"",
					).Return(nil, fmt.Errorf("foldername cannot be empty")).Times(1)
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "list files with invalid sort field",
			args: args{
				username:   "validUsername",
				foldername: "validFoldername",
				sortBy:     "invalid",
				order:      "asc",
				mock: func() {
					s.users.EXPECT().GetByUsername(gomock.Any(), user1.Username).Return(user1, nil).Times(1)
					s.folders.EXPECT().GetByName(gomock.Any(), user1, folder1.Name).Return(folder1, nil).Times(1)
					s.folders.EXPECT().ListFiles(
						gomock.Any(),
						user1,
						folder1,
						"invalid",
						"asc",
					).Return(nil, fmt.Errorf("invalid sort field")).Times(1)
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			got, err := s.vfs.ListFiles(tt.args.username, tt.args.foldername, tt.args.sortBy, tt.args.order)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListFiles() got = %v, want %v", got, tt.want)
			}
		})
	}
}
