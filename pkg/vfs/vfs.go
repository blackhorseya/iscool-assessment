//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package vfs

// VirtualFileSystem represents the entire file system with user management.
type VirtualFileSystem interface {
	Save() error
	Load() error

	UserManager
	FolderManager
	FileManager
}
