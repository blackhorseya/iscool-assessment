package vfs

import (
	"testing"

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
