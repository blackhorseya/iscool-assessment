package user

import (
	"context"
	"os"
	"strings"

	"github.com/blackhorseya/iscool-assessment/entity/model"
	"github.com/blackhorseya/iscool-assessment/entity/repo"
)

type system struct {
	path string
}

// NewSystem is used to create a new System.
func NewSystem(path string) (repo.UserManager, error) {
	return &system{
		path: strings.TrimRight(path, "/"),
	}, nil
}

func (i *system) Register(ctx context.Context, username string) (item *model.User, err error) {
	user, err := model.NewUser(username)
	if err != nil {
		return nil, err
	}

	// create a folder for the user
	err = os.MkdirAll(i.path+"/"+user.Username, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (i *system) GetByUsername(ctx context.Context, username string) (item *model.User, err error) {
	user, err := model.NewUser(username)
	if err != nil {
		return nil, err
	}

	// check if the user folder exists
	info, err := os.Stat(i.path + "/" + user.Username)
	if err != nil {
		return nil, err
	}

	if !info.IsDir() {
		return nil, os.ErrNotExist
	}

	return user, nil
}
