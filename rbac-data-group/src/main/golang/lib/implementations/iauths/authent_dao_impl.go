package iauths

import (
	"fmt"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac/lib/classes/authentications"
	"github.com/starter-go/security/random"
	"github.com/starter-go/v0/rbac-data-group/src/main/golang/api/daos"
	"gorm.io/gorm"
)

type AuthentDaoImpl struct {

	//starter:component

	_as func(daos.IAuthenticationDao) //starter:as("#")

	DBAgent    daos.IDatabaseAgent //starter:inject("#")
	UUIDGenSer random.UUIDService  //starter:inject("#")
}

func (inst *AuthentDaoImpl) innerGenUUID() lang.UUID {
	b := inst.UUIDGenSer.Build()
	b.Class("authentications.Entity")
	return b.Generate()
}

func (inst *AuthentDaoImpl) innerMakeItem() *authentications.Entity {
	return new(authentications.Entity)
}

func (inst *AuthentDaoImpl) innerMakeItemList() []*authentications.Entity {
	return make([]*authentications.Entity, 0)
}

// Delete implements [authentications.UserDAO].
func (inst *AuthentDaoImpl) Delete(db *gorm.DB, id authentications.ID) error {
	db = inst.GetDB(db)
	item := inst.innerMakeItem()
	item.ID = id
	res := db.Delete(item, id)
	err := res.Error
	return err
}

// Find implements [authentications.UserDAO].
func (inst *AuthentDaoImpl) Find(db *gorm.DB, id authentications.ID) (*authentications.Entity, error) {
	db = inst.GetDB(db)
	item := inst.innerMakeItem()
	res := db.First(item, id)
	err := res.Error
	return item, err
}

// GetDB implements [authentications.UserDAO].
func (inst *AuthentDaoImpl) GetDB(old *gorm.DB) *gorm.DB {
	return inst.DBAgent.DB(old)
}

// Insert implements [authentications.UserDAO].
func (inst *AuthentDaoImpl) Insert(db *gorm.DB, item *authentications.Entity) (*authentications.Entity, error) {

	db = inst.GetDB(db)
	uuid := inst.innerGenUUID()

	item.ID = 0
	item.UUID = uuid

	res := db.Create(item)
	err := res.Error
	return item, err
}

// Query implements [authentications.UserDAO].
func (inst *AuthentDaoImpl) Query(db *gorm.DB, q *authentications.Query) ([]*authentications.Entity, error) {

	panic("unimplemented")
}

// Update implements [authentications.UserDAO].
func (inst *AuthentDaoImpl) Update(db *gorm.DB, id authentications.ID, callback func(old *authentications.Entity) error) (*authentications.Entity, error) {

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

func (inst *AuthentDaoImpl) _impl() daos.IAuthenticationDao {
	return inst
}
