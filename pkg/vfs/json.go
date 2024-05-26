package vfs

type jsonFile struct {
	filePath string // Path to the JSON file
}

// NewJSONFile creates a new JSON file system.
func NewJSONFile(filePath string) VirtualFileSystem {
	return &jsonFile{
		filePath: filePath,
	}
}

func (vfs *jsonFile) Save() error {
	// TODO implement me
	panic("implement me")
}

func (vfs *jsonFile) Load() error {
	// TODO implement me
	panic("implement me")
}

func (vfs *jsonFile) GetByUsername(username string) (*User, error) {
	// TODO implement me
	panic("implement me")
}

func (vfs *jsonFile) RegisterUser(username string) error {
	// TODO implement me
	panic("implement me")
}

func (vfs *jsonFile) DeleteUser(username string) error {
	// TODO implement me
	panic("implement me")
}

func (vfs *jsonFile) ListUsers() []string {
	// TODO implement me
	panic("implement me")
}

func (vfs *jsonFile) CreateFolder(username, foldername, description string) error {
	// TODO implement me
	panic("implement me")
}

func (vfs *jsonFile) DeleteFolder(username, foldername string) error {
	// TODO implement me
	panic("implement me")
}

func (vfs *jsonFile) RenameFolder(username, foldername, newFoldername string) error {
	// TODO implement me
	panic("implement me")
}

func (vfs *jsonFile) ListFolders(username string, sortBy string, order string) ([]*Folder, error) {
	// TODO implement me
	panic("implement me")
}

func (vfs *jsonFile) CreateFile(username, foldername, filename, description string) error {
	// TODO implement me
	panic("implement me")
}

func (vfs *jsonFile) DeleteFile(username, foldername, filename string) error {
	// TODO implement me
	panic("implement me")
}

func (vfs *jsonFile) ListFiles(username, foldername, sortBy string, order string) ([]*File, error) {
	// TODO implement me
	panic("implement me")
}
