package itables

import (
	"fmt"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac/lib/classes/tables"
	"github.com/starter-go/security/random"
	"github.com/starter-go/v0/rbac-data-group/src/main/golang/api/daos"
	"gorm.io/gorm"
)

type TableDaoImpl struct {

	//starter:component

	_as func(daos.ITableDao) //starter:as("#")

	DBAgent    daos.IDatabaseAgent //starter:inject("#")
	UUIDGenSer random.UUIDService  //starter:inject("#")
}

func (inst *TableDaoImpl) innerGenUUID() lang.UUID {
	b := inst.UUIDGenSer.Build()
	b.Class("tables.Entity")
	return b.Generate()
}

func (inst *TableDaoImpl) innerMakeItem() *tables.Entity {
	return new(tables.Entity)
}

func (inst *TableDaoImpl) innerMakeItemList() []*tables.Entity {
	return make([]*tables.Entity, 0)
}

// Delete implements [tables.UserDAO].
func (inst *TableDaoImpl) Delete(db *gorm.DB, id tables.ID) error {
	db = inst.GetDB(db)
	item := inst.innerMakeItem()
	item.ID = id
	res := db.Delete(item, id)
	err := res.Error
	return err
}

// Find implements [tables.UserDAO].
func (inst *TableDaoImpl) Find(db *gorm.DB, id tables.ID) (*tables.Entity, error) {
	db = inst.GetDB(db)
	item := inst.innerMakeItem()
	res := db.First(item, id)
	err := res.Error
	return item, err
}

// GetDB implements [tables.UserDAO].
func (inst *TableDaoImpl) GetDB(old *gorm.DB) *gorm.DB {
	return inst.DBAgent.DB(old)
}

// Insert implements [tables.UserDAO].
func (inst *TableDaoImpl) Insert(db *gorm.DB, item *tables.Entity) (*tables.Entity, error) {

	db = inst.GetDB(db)
	uuid := inst.innerGenUUID()

	item.ID = 0
	item.UUID = uuid

	res := db.Create(item)
	err := res.Error
	return item, err
}

// Query implements [tables.UserDAO].
func (inst *TableDaoImpl) Query(db *gorm.DB, q *tables.Query) ([]*tables.Entity, error) {

	panic("unimplemented")
}

// Update implements [tables.UserDAO].
func (inst *TableDaoImpl) Update(db *gorm.DB, id tables.ID, callback func(old *tables.Entity) error) (*tables.Entity, error) {

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

func (inst *TableDaoImpl) _impl() daos.ITableDao {
	return inst
}
