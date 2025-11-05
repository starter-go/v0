package iusers

import (
	"context"

	"github.com/starter-go/v0/rbac-web-app/app/classes/users"
	"github.com/starter-go/v0/rbac-web-app/app/data/dxo"
	"github.com/starter-go/v0/rbac-web-app/app/data/entity"
	"github.com/starter-go/v0/rbac-web-app/app/web/dto"
)

type UserServiceAPI interface {
	users.Service
}

type UserServiceImpl struct {

	//starter:component

	_as func(users.Service) //starter:as("#")

	Dao users.DAO //starter:inject("#")

}

func (inst *UserServiceImpl) _impl() UserServiceAPI {
	return inst
}

func (inst *UserServiceImpl) Find(ctx context.Context, id dxo.UserID) (*dto.User, error) {

	// return nil, fmt.Errorf("no impl ")

	item := new(dto.User)

	return item, nil
}

func (inst *UserServiceImpl) Insert(cc context.Context, item *dto.User) (*dto.User, error) {

	it2 := new(entity.User)
	users.ConvertD2E(item, it2)

	it3, err := inst.Dao.Insert(nil, it2)
	if err != nil {
		return nil, err
	}

	it4 := new(dto.User)
	users.ConvertE2D(it3, it4)
	return it4, nil
}
