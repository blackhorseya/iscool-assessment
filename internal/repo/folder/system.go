package folder

import (
	"context"

	"github.com/blackhorseya/iscool-assessment/entity/model"
	"github.com/blackhorseya/iscool-assessment/entity/repo"
)

type system struct {
	path string
}

// NewSystem is used to create a new System.
func NewSystem(path string) (repo.FolderManager, error) {
	return &system{
		path: path,
	}, nil
}

func (i *system) GetByName(ctx context.Context, owner *model.User, foldername string) (item *model.Folder, err error) {
	// TODO implement me
	panic("implement me")
}

func (i *system) Create(ctx context.Context, owner *model.User, foldername, description string) (item *model.Folder, err error) {
	// TODO implement me
	panic("implement me")
}

func (i *system) Delete(ctx context.Context, owner *model.User, foldername string) (err error) {
	// TODO implement me
	panic("implement me")
}

func (i *system) Rename(ctx context.Context, owner *model.User, foldername, newFoldername string) (item *model.Folder, err error) {
	// TODO implement me
	panic("implement me")
}

func (i *system) List(ctx context.Context, owner *model.User, sortBy string, order string) (items []*model.Folder, err error) {
	// TODO implement me
	panic("implement me")
}

func (i *system) CreateFile(ctx context.Context, owner *model.User, folder *model.Folder, filename, description string) (item *model.File, err error) {
	// TODO implement me
	panic("implement me")
}

func (i *system) DeleteFile(ctx context.Context, owner *model.User, folder *model.Folder, filename string) (err error) {
	// TODO implement me
	panic("implement me")
}

func (i *system) ListFiles(ctx context.Context, owner *model.User, folder *model.Folder, sortBy string, order string) (items []*model.File, err error) {
	// TODO implement me
	panic("implement me")
}
