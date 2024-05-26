//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package vfs

import (
	"github.com/blackhorseya/iscool-assessment/entity/model"
)

// VirtualFileSystem represents the entire file system with user management.
type VirtualFileSystem interface {
	// RegisterUser registers a new user.
	RegisterUser(username string) (item *model.User, err error)

	CreateFolder(username, foldername, description string) (item *model.Folder, err error)
	DeleteFolder(username, foldername string) (err error)
	ListFolders(username string, sortBy string, order string) (items []*model.Folder, err error)
	RenameFolder(username, foldername, newFoldername string) (item *model.Folder, err error)

	CreateFile(username, foldername, filename, description string) (item *model.File, err error)
	DeleteFile(username, foldername, filename string) (err error)
	ListFiles(username, foldername string, sortBy string, order string) (items []*model.File, err error)
}
