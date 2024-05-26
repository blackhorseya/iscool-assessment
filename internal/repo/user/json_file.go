package user

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/blackhorseya/iscool-assessment/entity/model"
	"github.com/blackhorseya/iscool-assessment/entity/repo"
)

type jsonFile struct {
	users map[string]*model.User
	path  string
}

// NewJSONFile is used to create a new JSONFile.
func NewJSONFile(path string) (repo.UserManager, error) {
	instance := &jsonFile{
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
	// TODO implement me
	panic("implement me")
}

func (i *jsonFile) GetByUsername(ctx context.Context, username string) (item *model.User, err error) {
	// TODO implement me
	panic("implement me")
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
