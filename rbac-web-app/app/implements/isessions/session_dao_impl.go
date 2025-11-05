package isessions

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/security/random"
	"github.com/starter-go/v0/rbac-web-app/app/classes/sessions"
	"github.com/starter-go/v0/rbac-web-app/app/data/database"
	"github.com/starter-go/v0/rbac-web-app/app/data/dxo"
	"github.com/starter-go/v0/rbac-web-app/app/data/entity"
	"gorm.io/gorm"
)

type SessionDaoAPI interface {
	sessions.DAO
}

type SessionDaoImpl struct {

	//starter:component

	_as func(sessions.DAO) //starter:as("#")

	Agent database.Agent //starter:inject("#")

	UUIDSer random.UUIDService //starter:inject("#")

}

func (inst *SessionDaoImpl) _impl() SessionDaoAPI {
	return inst
}

func (inst *SessionDaoImpl) makeModel() *entity.Session {
	return new(entity.Session)
}

func (inst *SessionDaoImpl) makeModelList() []*entity.Session {
	return make([]*entity.Session, 0)
}

func (inst *SessionDaoImpl) makeUUID() lang.UUID {
	ub := inst.UUIDSer.Build()
	ub.Class("rbac.entity.Session")
	return ub.Generate()
}

func (inst *SessionDaoImpl) Find(db *gorm.DB, id dxo.SessionID) (*entity.Session, error) {

	db = inst.Agent.DB(db)
	item := inst.makeModel()

	res := db.First(&item, id)
	err := res.Error
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (inst *SessionDaoImpl) Insert(db *gorm.DB, item *entity.Session) (*entity.Session, error) {

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
