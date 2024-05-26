package fsx

// VirtualFileSystem represents the entire file system with user management
type VirtualFileSystem struct {
	Users map[string]*User `json:"users"`
}

// NewVFS creates a new VirtualFileSystem
func NewVFS() *VirtualFileSystem {
	return &VirtualFileSystem{
		Users: make(map[string]*User),
	}
}
