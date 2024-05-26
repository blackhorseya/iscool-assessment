package vfs

import (
	"testing"

	"github.com/blackhorseya/iscool-assessment/pkg/vfs"
	"github.com/stretchr/testify/suite"
)

type suiteTester struct {
	suite.Suite

	vfs vfs.VirtualFileSystem
}

func TestAll(t *testing.T) {
	suite.Run(t, new(suiteTester))
}
