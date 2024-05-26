//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package repo

import (
	"context"

	"github.com/blackhorseya/iscool-assessment/entity/model"
)

// UserManager defines the interface for user management.
type UserManager interface {
	Register(ctx context.Context, username string) (item *model.User, err error)
	GetByUsername(ctx context.Context, username string) (item *model.User, err error)
}
