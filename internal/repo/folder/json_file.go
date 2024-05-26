package folder

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/blackhorseya/iscool-assessment/entity/model"
	"github.com/blackhorseya/iscool-assessment/entity/repo"
	"github.com/blackhorseya/iscool-assessment/pkg/utils"
)

type jsonFile struct {
	sync.Mutex

	users map[string]*model.User
	path  string
}

// NewJSONFile is used to create a new JSONFile.
func NewJSONFile(path string) (repo.FolderManager, error) {
	instance := &jsonFile{
		Mutex: sync.Mutex{},
		users: make(map[string]*model.User),
		path:  path,
	}

	err := instance.Load()
	if err != nil {
		return nil, err
	}

	return instance, nil
}

func (i *jsonFile) Create(
	ctx context.Context,
	owner *model.User,
	foldername, description string,
) (item *model.Folder, err error) {
	i.Lock()
	defer i.Unlock()

	err = i.Load()
	if err != nil {
		return nil, err
	}

	user, exists := i.users[owner.Username]
	if !exists {
		return nil, fmt.Errorf("the %s doesn't exist", owner.Username)
	}

	if _, exists = user.Folders[foldername]; exists {
		return nil, fmt.Errorf("the %s has already existed", foldername)
	}

	folder, err := model.NewFolder(owner, foldername, description)
	if err != nil {
		return nil, err
	}

	user.Folders[foldername] = folder

	err = i.Save()
	if err != nil {
		return nil, err
	}

	return folder, nil
}

func (i *jsonFile) Delete(ctx context.Context, owner *model.User, foldername string) (err error) {
	// TODO implement me
	panic("implement me")
}

func (i *jsonFile) Rename(
	ctx context.Context,
	owner *model.User,
	foldername, newFoldername string,
) (item *model.Folder, err error) {
	// TODO implement me
	panic("implement me")
}

func (i *jsonFile) List(
	ctx context.Context,
	owner *model.User,
	sortBy string,
	order string,
) (items []*model.Folder, err error) {
	// TODO implement me
	panic("implement me")
}

func (i *jsonFile) CreateFile(
	ctx context.Context,
	owner *model.User,
	folder *model.Folder,
	filename, description string,
) (item *model.File, err error) {
	// TODO implement me
	panic("implement me")
}

func (i *jsonFile) DeleteFile(
	ctx context.Context,
	owner *model.User,
	folder *model.Folder,
	filename string,
) (err error) {
	// TODO implement me
	panic("implement me")
}

func (i *jsonFile) ListFiles(
	ctx context.Context,
	owner *model.User,
	folder *model.Folder,
	sortBy string,
	order string,
) (items []*model.File, err error) {
	// TODO implement me
	panic("implement me")
}

// Save is used to save the data to the file.
func (i *jsonFile) Save() (err error) {
	// Ensure the directory exists
	if err = utils.EnsureDir(i.path); err != nil {
		return fmt.Errorf("failed to ensure directory: %w", err)
	}

	data, err := json.MarshalIndent(i.users, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	return os.WriteFile(i.path, data, 0600)
}

// Load is used to load the data from the file.
func (i *jsonFile) Load() (err error) {
	data, err := os.ReadFile(i.path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}

		return fmt.Errorf("failed to read file: %w", err)
	}

	return json.Unmarshal(data, &i.users)
}
