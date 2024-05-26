package fsx

var _ UserManager = &VirtualFileSystem{}

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

func (vfs *VirtualFileSystem) RegisterUser(username string) error {
	// todo: 2024/5/26|sean|implement me
	panic("implement me")
}

func (vfs *VirtualFileSystem) DeleteUser(username string) error {
	// todo: 2024/5/26|sean|implement me
	panic("implement me")
}

func (vfs *VirtualFileSystem) ListUsers() []string {
	// todo: 2024/5/26|sean|implement me
	panic("implement me")
}
