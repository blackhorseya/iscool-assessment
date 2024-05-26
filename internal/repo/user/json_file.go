package user

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
func NewJSONFile(path string) (repo.UserManager, error) {
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

func (i *jsonFile) Register(ctx context.Context, username string) (item *model.User, err error) {
	i.Lock()
	defer i.Unlock()

	if _, exists := i.users[username]; exists {
		return nil, fmt.Errorf("the %s has already existed", username)
	}

	user, err := model.NewUser(username)
	if err != nil {
		return nil, err
	}

	i.users[username] = user

	err = i.Save()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (i *jsonFile) GetByUsername(ctx context.Context, username string) (item *model.User, err error) {
	i.Lock()
	defer i.Unlock()

	user, exists := i.users[username]
	if !exists {
		return nil, fmt.Errorf("the %s doesn't exist", username)
	}

	return user, nil
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
