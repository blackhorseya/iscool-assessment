package folder

import (
	"context"

	"github.com/blackhorseya/iscool-assessment/entity/model"
	"github.com/blackhorseya/iscool-assessment/entity/repo"
)

type jsonFile struct {
	path string
}

// NewJSONFile is used to create a new JSONFile.
func NewJSONFile(path string) (repo.FolderManager, error) {
	return &jsonFile{
		path: path,
	}, nil
}

func (i *jsonFile) Create(
	ctx context.Context,
	owner *model.User,
	foldername, description string,
) (item *model.Folder, err error) {
	// TODO implement me
	panic("implement me")
}

func (i *jsonFile) Delete(ctx context.Context, owner *model.User, foldername string) (err error) {
	// TODO implement me
	panic("implement me")
}

func (i *jsonFile) Rename(
	ctx context.Context,
	owner *model.User,
	foldername, newFoldername string,
) (item *model.Folder, err error) {
	// TODO implement me
	panic("implement me")
}

func (i *jsonFile) List(
	ctx context.Context,
	owner *model.User,
	sortBy string,
	order string,
) (items []*model.Folder, err error) {
	// TODO implement me
	panic("implement me")
}

func (i *jsonFile) CreateFile(
	ctx context.Context,
	owner *model.User,
	folder *model.Folder,
	filename, description string,
) (item *model.File, err error) {
	// TODO implement me
	panic("implement me")
}

func (i *jsonFile) DeleteFile(
	ctx context.Context,
	owner *model.User,
	folder *model.Folder,
	filename string,
) (err error) {
	// TODO implement me
	panic("implement me")
}

func (i *jsonFile) ListFiles(
	ctx context.Context,
	owner *model.User,
	folder *model.Folder,
	sortBy string,
	order string,
) (items []*model.File, err error) {
	// TODO implement me
	panic("implement me")
}
