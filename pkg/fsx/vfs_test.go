package fsx

import (
	"testing"
)

func TestNewVFS(t *testing.T) {
	vfs := NewVFS()
	if vfs == nil {
		t.Error("NewVFS() returned nil")
	}
}
