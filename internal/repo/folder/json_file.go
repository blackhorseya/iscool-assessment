package folder

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"sync"

	"github.com/blackhorseya/iscool-assessment/entity/model"
	"github.com/blackhorseya/iscool-assessment/entity/repo"
	"github.com/blackhorseya/iscool-assessment/pkg/utils"
)

const orderAsc = "asc"

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

	return instance, nil
}

func (i *jsonFile) GetByName(
	ctx context.Context,
	owner *model.User,
	foldername string,
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

	folder, exists := user.Folders[foldername]
	if !exists {
		return nil, fmt.Errorf("the %s doesn't exist", foldername)
	}

	return folder, nil
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
	i.Lock()
	defer i.Unlock()

	err = i.Load()
	if err != nil {
		return err
	}

	user, exists := i.users[owner.Username]
	if !exists {
		return fmt.Errorf("the %s doesn't exist", owner.Username)
	}

	if _, exists = user.Folders[foldername]; !exists {
		return fmt.Errorf("the %s doesn't exist", foldername)
	}

	delete(user.Folders, foldername)

	err = i.Save()
	if err != nil {
		return err
	}

	return nil
}

func (i *jsonFile) Rename(
	ctx context.Context,
	owner *model.User,
	foldername, newFoldername string,
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

	folder, exists := user.Folders[foldername]
	if !exists {
		return nil, fmt.Errorf("the %s doesn't exist", foldername)
	}

	if _, exists = user.Folders[newFoldername]; exists {
		return nil, fmt.Errorf("the %s has already existed", newFoldername)
	}

	delete(user.Folders, foldername)
	folder.Name = newFoldername
	user.Folders[newFoldername] = folder

	return folder, nil
}

func (i *jsonFile) List(
	ctx context.Context,
	owner *model.User,
	sortBy string,
	order string,
) (items []*model.Folder, err error) {
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

	var folders []*model.Folder
	for _, folder := range user.Folders {
		folders = append(folders, folder)
	}

	sort.Slice(folders, func(i, j int) bool {
		switch sortBy {
		case "name":
			if order == orderAsc {
				return folders[i].Name < folders[j].Name
			}
			return folders[i].Name > folders[j].Name
		case "created":
			if order == orderAsc {
				return folders[i].CreatedAt.Before(folders[j].CreatedAt)
			}
			return folders[i].CreatedAt.After(folders[j].CreatedAt)
		default:
			return folders[i].Name < folders[j].Name
		}
	})

	return folders, nil
}

func (i *jsonFile) CreateFile(
	ctx context.Context,
	owner *model.User,
	folder *model.Folder,
	filename, description string,
) (item *model.File, err error) {
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

	folder, exists = user.Folders[folder.Name]
	if !exists {
		return nil, fmt.Errorf("the %s doesn't exist", folder.Name)
	}

	if _, exists = folder.Files[filename]; exists {
		return nil, fmt.Errorf("the %s has already existed", filename)
	}

	file, err := model.NewFile(owner, folder, filename, description)
	if err != nil {
		return nil, err
	}

	folder.Files[filename] = file

	err = i.Save()
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (i *jsonFile) DeleteFile(
	ctx context.Context,
	owner *model.User,
	folder *model.Folder,
	filename string,
) (err error) {
	i.Lock()
	defer i.Unlock()

	err = i.Load()
	if err != nil {
		return err
	}

	user, exists := i.users[owner.Username]
	if !exists {
		return fmt.Errorf("the %s doesn't exist", owner.Username)
	}

	folder, exists = user.Folders[folder.Name]
	if !exists {
		return fmt.Errorf("the %s doesn't exist", folder.Name)
	}

	if _, exists = folder.Files[filename]; !exists {
		return fmt.Errorf("the %s doesn't exist", filename)
	}

	delete(folder.Files, filename)

	err = i.Save()
	if err != nil {
		return err
	}

	return nil
}

func (i *jsonFile) ListFiles(
	ctx context.Context,
	owner *model.User,
	folder *model.Folder,
	sortBy string,
	order string,
) (items []*model.File, err error) {
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

	folder, exists = user.Folders[folder.Name]
	if !exists {
		return nil, fmt.Errorf("the %s doesn't exist", folder.Name)
	}

	var files []*model.File
	for _, file := range folder.Files {
		files = append(files, file)
	}

	sort.Slice(files, func(i, j int) bool {
		switch sortBy {
		case "name":
			if order == orderAsc {
				return files[i].Name < files[j].Name
			}
			return files[i].Name > files[j].Name
		case "created":
			if order == orderAsc {
				return files[i].CreatedAt.Before(files[j].CreatedAt)
			}
			return files[i].CreatedAt.After(files[j].CreatedAt)
		default:
			return files[i].Name < files[j].Name
		}
	})

	return files, nil
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
