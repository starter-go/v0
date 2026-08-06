package iroles

import (
	"fmt"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac/lib/classes/roles"
	"github.com/starter-go/security/random"
	"github.com/starter-go/v0/rbac-data-group/src/main/golang/api/daos"
	"gorm.io/gorm"
)

type RoleDaoImpl struct {

	//starter:component

	_as func(daos.IRoleDao) //starter:as("#")

	DBAgent    daos.IDatabaseAgent //starter:inject("#")
	UUIDGenSer random.UUIDService  //starter:inject("#")
}

func (inst *RoleDaoImpl) innerGenUUID() lang.UUID {
	b := inst.UUIDGenSer.Build()
	b.Class("roles.Entity")
	return b.Generate()
}

func (inst *RoleDaoImpl) innerMakeItem() *roles.Entity {
	return new(roles.Entity)
}

func (inst *RoleDaoImpl) innerMakeItemList() []*roles.Entity {
	return make([]*roles.Entity, 0)
}

// Delete implements [roles.UserDAO].
func (inst *RoleDaoImpl) Delete(db *gorm.DB, id roles.ID) error {
	db = inst.GetDB(db)
	item := inst.innerMakeItem()
	item.ID = id
	res := db.Delete(item, id)
	err := res.Error
	return err
}

// Find implements [roles.UserDAO].
func (inst *RoleDaoImpl) Find(db *gorm.DB, id roles.ID) (*roles.Entity, error) {
	db = inst.GetDB(db)
	item := inst.innerMakeItem()
	res := db.First(item, id)
	err := res.Error
	return item, err
}

// GetDB implements [roles.UserDAO].
func (inst *RoleDaoImpl) GetDB(old *gorm.DB) *gorm.DB {
	return inst.DBAgent.DB(old)
}

// Insert implements [roles.UserDAO].
func (inst *RoleDaoImpl) Insert(db *gorm.DB, item *roles.Entity) (*roles.Entity, error) {

	db = inst.GetDB(db)
	uuid := inst.innerGenUUID()

	item.ID = 0
	item.UUID = uuid

	res := db.Create(item)
	err := res.Error
	return item, err
}

// Query implements [roles.UserDAO].
func (inst *RoleDaoImpl) Query(db *gorm.DB, q *roles.Query) ([]*roles.Entity, error) {

	panic("unimplemented")
}

// Update implements [roles.UserDAO].
func (inst *RoleDaoImpl) Update(db *gorm.DB, id roles.ID, callback func(old *roles.Entity) error) (*roles.Entity, error) {

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

func (inst *RoleDaoImpl) _impl() daos.IRoleDao {
	return inst
}
