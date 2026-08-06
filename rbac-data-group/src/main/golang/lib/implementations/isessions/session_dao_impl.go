package isessions

import (
	"fmt"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac/lib/classes/sessions"
	"github.com/starter-go/security/random"
	"github.com/starter-go/v0/rbac-data-group/src/main/golang/api/daos"
	"gorm.io/gorm"
)

type SessionDaoImpl struct {

	//starter:component

	_as func(daos.ISessionDao) //starter:as("#")

	DBAgent    daos.IDatabaseAgent //starter:inject("#")
	UUIDGenSer random.UUIDService  //starter:inject("#")
}

func (inst *SessionDaoImpl) innerGenUUID() lang.UUID {
	b := inst.UUIDGenSer.Build()
	b.Class("sessions.Entity")
	return b.Generate()
}

func (inst *SessionDaoImpl) innerMakeItem() *sessions.Entity {
	return new(sessions.Entity)
}

func (inst *SessionDaoImpl) innerMakeItemList() []*sessions.Entity {
	return make([]*sessions.Entity, 0)
}

// Delete implements [sessions.UserDAO].
func (inst *SessionDaoImpl) Delete(db *gorm.DB, id sessions.ID) error {
	db = inst.GetDB(db)
	item := inst.innerMakeItem()
	item.ID = id
	res := db.Delete(item, id)
	err := res.Error
	return err
}

// Find implements [sessions.UserDAO].
func (inst *SessionDaoImpl) Find(db *gorm.DB, id sessions.ID) (*sessions.Entity, error) {
	db = inst.GetDB(db)
	item := inst.innerMakeItem()
	res := db.First(item, id)
	err := res.Error
	return item, err
}

// GetDB implements [sessions.UserDAO].
func (inst *SessionDaoImpl) GetDB(old *gorm.DB) *gorm.DB {
	return inst.DBAgent.DB(old)
}

// Insert implements [sessions.UserDAO].
func (inst *SessionDaoImpl) Insert(db *gorm.DB, item *sessions.Entity) (*sessions.Entity, error) {

	db = inst.GetDB(db)
	uuid := inst.innerGenUUID()

	item.ID = 0
	item.UUID = uuid

	res := db.Create(item)
	err := res.Error
	return item, err
}

// Query implements [sessions.UserDAO].
func (inst *SessionDaoImpl) Query(db *gorm.DB, q *sessions.Query) ([]*sessions.Entity, error) {

	panic("unimplemented")
}

// Update implements [sessions.UserDAO].
func (inst *SessionDaoImpl) Update(db *gorm.DB, id sessions.ID, callback func(old *sessions.Entity) error) (*sessions.Entity, error) {

	if callback == nil {
		return nil, fmt.Errorf("callback func is nil")
	}

	db = inst.GetDB(db)

	// find
	item := inst.innerMakeItem()
	res := db.First(item, id)
	err := res.Error
	if err != nil {
		return nil, err
	}

	// callback
	err = callback(item)
	if err != nil {
		return nil, err
	}

	// save
	res = db.Save(item)
	err = res.Error
	return item, err
}

func (inst *SessionDaoImpl) _impl() daos.ISessionDao {
	return inst
}
