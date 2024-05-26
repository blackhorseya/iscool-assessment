package vfs

import (
	"os"
	"testing"

	"github.com/blackhorseya/iscool-assessment/entity/repo"
	"github.com/blackhorseya/iscool-assessment/internal/repo/user"
	"github.com/blackhorseya/iscool-assessment/pkg/vfs"
	"github.com/stretchr/testify/suite"
)

type suiteIntegration struct {
	suite.Suite

	users   repo.UserManager
	folders repo.FolderManager
	vfs     vfs.VirtualFileSystem
}

func (s *suiteIntegration) SetupTest() {
	users, err := user.NewJSONFile("out/users.json")
	s.Require().NoError(err)
	s.users = users

	s.folders = nil

	s.vfs = New(s.users, s.folders)
}

func (s *suiteIntegration) TearDownTest() {
	_ = os.Remove("out/users.json")
}

func TestIntegration(t *testing.T) {
	suite.Run(t, new(suiteIntegration))
}
