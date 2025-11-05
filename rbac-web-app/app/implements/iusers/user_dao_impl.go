package iusers

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/security/random"
	"github.com/starter-go/v0/rbac-web-app/app/classes/users"
	"github.com/starter-go/v0/rbac-web-app/app/data/database"
	"github.com/starter-go/v0/rbac-web-app/app/data/dxo"
	"github.com/starter-go/v0/rbac-web-app/app/data/entity"
	"gorm.io/gorm"
)

type UserDaoAPI interface {
	users.DAO
}

type UserDaoImpl struct {

	//starter:component

	_as func(users.DAO) //starter:as("#")

	Agent database.Agent //starter:inject("#")

	UUIDSer random.UUIDService //starter:inject("#")

}

func (inst *UserDaoImpl) _impl() UserDaoAPI {
	return inst
}

func (inst *UserDaoImpl) makeModel() *entity.User {
	return new(entity.User)
}

func (inst *UserDaoImpl) makeModelList() []*entity.User {
	return make([]*entity.User, 0)
}

func (inst *UserDaoImpl) makeUUID() lang.UUID {
	ub := inst.UUIDSer.Build()
	ub.Class("rbac.entity.User")
	return ub.Generate()
}

func (inst *UserDaoImpl) Find(db *gorm.DB, id dxo.UserID) (*entity.User, error) {

	db = inst.Agent.DB(db)
	item := inst.makeModel()

	res := db.First(&item, id)
	err := res.Error
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (inst *UserDaoImpl) Insert(db *gorm.DB, item *entity.User) (*entity.User, error) {

	db = inst.Agent.DB(db)
	uuid := inst.makeUUID()

	item.ID = 0
	item.UUID = uuid

	res := db.Create(item)
	err := res.Error
	if err != nil {
		return nil, err
	}

	return item, nil
}
