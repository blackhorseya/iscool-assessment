package user

import (
	"context"

	"github.com/blackhorseya/iscool-assessment/entity/model"
	"github.com/blackhorseya/iscool-assessment/entity/repo"
)

type system struct {
	path string
}

// NewSystem is used to create a new System.
func NewSystem(path string) (repo.UserManager, error) {
	return &system{
		path: path,
	}, nil
}

func (i *system) Register(ctx context.Context, username string) (item *model.User, err error) {
	// TODO implement me
	panic("implement me")
}

func (i *system) GetByUsername(ctx context.Context, username string) (item *model.User, err error) {
	// TODO implement me
	panic("implement me")
}
