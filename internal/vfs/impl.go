package vfs

import (
	"context"

	"github.com/blackhorseya/iscool-assessment/entity/model"
	"github.com/blackhorseya/iscool-assessment/entity/repo"
	"github.com/blackhorseya/iscool-assessment/pkg/vfs"
)

type impl struct {
	users   repo.UserManager
	folders repo.FolderManager
}

// New is used to create a new VirtualFileSystem.
func New(users repo.UserManager, folders repo.FolderManager) vfs.VirtualFileSystem {
	return &impl{
		users:   users,
		folders: folders,
	}
}

func (i *impl) RegisterUser(username string) (item *model.User, err error) {
	return i.users.Register(context.TODO(), username)
}

func (i *impl) CreateFolder(username, foldername, description string) (item *model.Folder, err error) {
	// TODO implement me
	panic("implement me")
}

func (i *impl) DeleteFolder(username, foldername string) (err error) {
	// TODO implement me
	panic("implement me")
}

func (i *impl) ListFolders(username string, sortBy string, order string) (items []*model.Folder, err error) {
	// TODO implement me
	panic("implement me")
}

func (i *impl) RenameFolder(username, foldername, newFoldername string) (item *model.Folder, err error) {
	// TODO implement me
	panic("implement me")
}

func (i *impl) CreateFile(username, foldername, filename, description string) (item *model.File, err error) {
	// TODO implement me
	panic("implement me")
}

func (i *impl) DeleteFile(username, foldername, filename string) (err error) {
	// TODO implement me
	panic("implement me")
}

func (i *impl) ListFiles(username, foldername string, sortBy string, order string) (items []*model.File, err error) {
	// TODO implement me
	panic("implement me")
}
