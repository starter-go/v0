package authx

import (
	"context"

	"github.com/starter-go/v0/rbac-web-app/app/data/dxo"
	"github.com/starter-go/v0/rbac-web-app/app/data/entity"
)

type UserAuthNameService interface {
	FindUserEntity(cc context.Context, name dxo.UserAuthName) (*entity.User, error)
}
