//go:build wireinject

//go:generate wire

package cmd

import (
	"github.com/blackhorseya/iscool-assessment/internal/repo/folder"
	"github.com/blackhorseya/iscool-assessment/internal/repo/user"
	vfsI "github.com/blackhorseya/iscool-assessment/internal/vfs"
	vfs2 "github.com/blackhorseya/iscool-assessment/pkg/vfs"
	"github.com/google/wire"
)

func NewVFSWithJSON(path string) (vfs2.VirtualFileSystem, error) {
	panic(wire.Build(
		vfsI.New,
		folder.NewJSONFile,
		user.NewJSONFile,
	))
}

func NewVFSWithSystem(path string) (vfs2.VirtualFileSystem, error) {
	panic(wire.Build(
		vfsI.New,
		folder.NewSystem,
		user.NewSystem,
	))
}
