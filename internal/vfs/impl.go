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
	user, err := i.getUserByUsername(username)
	if err != nil {
		return nil, err
	}

	return i.folders.Create(context.TODO(), user, foldername, description)
}

func (i *impl) DeleteFolder(username, foldername string) (err error) {
	user, err := i.getUserByUsername(username)
	if err != nil {
		return err
	}

	return i.folders.Delete(context.TODO(), user, foldername)
}

func (i *impl) ListFolders(username string, sortBy string, order string) (items []*model.Folder, err error) {
	user, err := i.getUserByUsername(username)
	if err != nil {
		return nil, err
	}

	return i.folders.List(context.TODO(), user, sortBy, order)
}

func (i *impl) RenameFolder(username, foldername, newFoldername string) (item *model.Folder, err error) {
	user, err := i.getUserByUsername(username)
	if err != nil {
		return nil, err
	}

	return i.folders.Rename(context.TODO(), user, foldername, newFoldername)
}

func (i *impl) CreateFile(username, foldername, filename, description string) (item *model.File, err error) {
	user, err := i.getUserByUsername(username)
	if err != nil {
		return nil, err
	}

	folder, err := i.folders.GetByName(context.TODO(), user, foldername)
	if err != nil {
		return nil, err
	}

	return i.folders.CreateFile(context.TODO(), user, folder, filename, description)
}

func (i *impl) DeleteFile(username, foldername, filename string) (err error) {
	user, err := i.getUserByUsername(username)
	if err != nil {
		return err
	}

	folder, err := i.folders.GetByName(context.TODO(), user, foldername)
	if err != nil {
		return err
	}

	return i.folders.DeleteFile(context.TODO(), user, folder, filename)
}

func (i *impl) ListFiles(username, foldername string, sortBy string, order string) (items []*model.File, err error) {
	user, err := i.getUserByUsername(username)
	if err != nil {
		return nil, err
	}

	folder, err := i.folders.GetByName(context.TODO(), user, foldername)
	if err != nil {
		return nil, err
	}

	return i.folders.ListFiles(context.TODO(), user, folder, sortBy, order)
}

func (i *impl) getUserByUsername(username string) (item *model.User, err error) {
	return i.users.GetByUsername(context.TODO(), username)
}
