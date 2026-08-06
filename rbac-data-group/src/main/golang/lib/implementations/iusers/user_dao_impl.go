package iusers

import (
	"fmt"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac/lib/classes/users"
	"github.com/starter-go/security/random"
	"github.com/starter-go/v0/rbac-data-group/src/main/golang/api/daos"
	"gorm.io/gorm"
)

type UserDaoImpl struct {

	//starter:component

	_as func(daos.IUserDao) //starter:as("#")

	DBAgent    daos.IDatabaseAgent //starter:inject("#")
	UUIDGenSer random.UUIDService  //starter:inject("#")
}

func (inst *UserDaoImpl) innerGenUUID() lang.UUID {
	b := inst.UUIDGenSer.Build()
	b.Class("users.Entity")
	return b.Generate()
}

func (inst *UserDaoImpl) innerMakeItem() *users.Entity {
	return new(users.Entity)
}

func (inst *UserDaoImpl) innerMakeItemList() []*users.Entity {
	return make([]*users.Entity, 0)
}

// Delete implements [users.UserDAO].
func (inst *UserDaoImpl) Delete(db *gorm.DB, id users.ID) error {
	db = inst.GetDB(db)
	item := inst.innerMakeItem()
	item.ID = id
	res := db.Delete(item, id)
	err := res.Error
	return err
}

// Find implements [users.UserDAO].
func (inst *UserDaoImpl) Find(db *gorm.DB, id users.ID) (*users.Entity, error) {
	db = inst.GetDB(db)
	item := inst.innerMakeItem()
	res := db.First(item, id)
	err := res.Error
	return item, err
}

// GetDB implements [users.UserDAO].
func (inst *UserDaoImpl) GetDB(old *gorm.DB) *gorm.DB {
	return inst.DBAgent.DB(old)
}

// Insert implements [users.UserDAO].
func (inst *UserDaoImpl) Insert(db *gorm.DB, item *users.Entity) (*users.Entity, error) {

	db = inst.GetDB(db)
	uuid := inst.innerGenUUID()

	item.ID = 0
	item.UUID = uuid

	res := db.Create(item)
	err := res.Error
	return item, err
}

// Query implements [users.UserDAO].
func (inst *UserDaoImpl) Query(db *gorm.DB, q *users.Query) ([]*users.Entity, error) {

	panic("unimplemented")
}

// Update implements [users.UserDAO].
func (inst *UserDaoImpl) Update(db *gorm.DB, id users.ID, callback func(old *users.Entity) error) (*users.Entity, error) {

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

func (inst *UserDaoImpl) _impl() daos.IUserDao {
	return inst
}
