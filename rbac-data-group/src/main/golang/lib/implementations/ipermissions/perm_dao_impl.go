package ipermissions

import (
	"fmt"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac/lib/classes/permissions"
	"github.com/starter-go/security/random"
	"github.com/starter-go/v0/rbac-data-group/src/main/golang/api/daos"
	"gorm.io/gorm"
)

type PermissionDaoImpl struct {

	//starter:component

	_as func(daos.IPermissionDao) //starter:as("#")

	DBAgent    daos.IDatabaseAgent //starter:inject("#")
	UUIDGenSer random.UUIDService  //starter:inject("#")
}

func (inst *PermissionDaoImpl) innerGenUUID() lang.UUID {
	b := inst.UUIDGenSer.Build()
	b.Class("permissions.Entity")
	return b.Generate()
}

func (inst *PermissionDaoImpl) innerMakeItem() *permissions.Entity {
	return new(permissions.Entity)
}

func (inst *PermissionDaoImpl) innerMakeItemList() []*permissions.Entity {
	return make([]*permissions.Entity, 0)
}

// Delete implements [permissions.UserDAO].
func (inst *PermissionDaoImpl) Delete(db *gorm.DB, id permissions.ID) error {
	db = inst.GetDB(db)
	item := inst.innerMakeItem()
	item.ID = id
	res := db.Delete(item, id)
	err := res.Error
	return err
}

// Find implements [permissions.UserDAO].
func (inst *PermissionDaoImpl) Find(db *gorm.DB, id permissions.ID) (*permissions.Entity, error) {
	db = inst.GetDB(db)
	item := inst.innerMakeItem()
	res := db.First(item, id)
	err := res.Error
	return item, err
}

// GetDB implements [permissions.UserDAO].
func (inst *PermissionDaoImpl) GetDB(old *gorm.DB) *gorm.DB {
	return inst.DBAgent.DB(old)
}

// Insert implements [permissions.UserDAO].
func (inst *PermissionDaoImpl) Insert(db *gorm.DB, item *permissions.Entity) (*permissions.Entity, error) {

	db = inst.GetDB(db)
	uuid := inst.innerGenUUID()

	item.ID = 0
	item.UUID = uuid

	res := db.Create(item)
	err := res.Error
	return item, err
}

// Query implements [permissions.UserDAO].
func (inst *PermissionDaoImpl) Query(db *gorm.DB, q *permissions.Query) ([]*permissions.Entity, error) {

	panic("unimplemented")
}

// Update implements [permissions.UserDAO].
func (inst *PermissionDaoImpl) Update(db *gorm.DB, id permissions.ID, callback func(old *permissions.Entity) error) (*permissions.Entity, error) {

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

func (inst *PermissionDaoImpl) _impl() daos.IPermissionDao {
	return inst
}
