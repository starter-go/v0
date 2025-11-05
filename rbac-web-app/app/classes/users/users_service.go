package users

import (
	"context"

	"github.com/starter-go/v0/rbac-web-app/app/data/dxo"
	"github.com/starter-go/v0/rbac-web-app/app/web/dto"
)

type Service interface {

	// query

	Find(ctx context.Context, id dxo.UserID) (*dto.User, error)

	// edit

	Insert(cc context.Context, item *dto.User) (*dto.User, error)
}
