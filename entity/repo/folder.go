//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package repo

import (
	"context"

	"github.com/blackhorseya/iscool-assessment/entity/model"
)

// FolderManager defines the interface for folder management.
type FolderManager interface {
	GetByName(ctx context.Context, owner *model.User, foldername string) (item *model.Folder, err error)
	Create(ctx context.Context, owner *model.User, foldername, description string) (item *model.Folder, err error)
	Delete(ctx context.Context, owner *model.User, foldername string) (err error)
	Rename(ctx context.Context, owner *model.User, foldername, newFoldername string) (item *model.Folder, err error)
	List(ctx context.Context, owner *model.User, sortBy string, order string) (items []*model.Folder, err error)

	CreateFile(
		ctx context.Context,
		owner *model.User,
		folder *model.Folder,
		filename, description string,
	) (item *model.File, err error)
	DeleteFile(
		ctx context.Context,
		owner *model.User,
		folder *model.Folder,
		filename string,
	) (err error)
	ListFiles(
		ctx context.Context,
		owner *model.User,
		folder *model.Folder,
		sortBy string,
		order string,
	) (items []*model.File, err error)
}
