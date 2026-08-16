package itables

import (
	"fmt"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/api/classes/tables"

	"github.com/starter-go/v0/libdao"
	"github.com/starter-go/v0/libdao/api/libdaoapi"
	"github.com/starter-go/v0/rbac-data-group/src/main/golang/api/daos"

	"gorm.io/gorm"
)

type TableDaoImpl struct {

	//starter:component

	_as func(rbac.TableDAO) //starter:as(".")

	ConfigClass    string //starter:inject("${rbac-data-group.sql.class}")
	ConfigEnabled  bool   //starter:inject("${rbac-data-group.sql.enabled}")
	ConfigPriority int    //starter:inject("${rbac-data-group.sql.priority}")

	DBAgent daos.IDatabaseAgent //starter:inject("#")

}

// GetRegistration implements [tables.DAO].
func (inst *TableDaoImpl) GetRegistration() *libdaoapi.DaoRegistration {

	r1 := &libdao.DaoRegistration{
		Name:     "TableDaoImpl",
		ID:       "sql-rbac-table-dao",
		Class:    inst.ConfigClass,
		Enabled:  inst.ConfigEnabled,
		Priority: inst.ConfigPriority,
		DAO:      inst,
	}

	return r1
}

func (inst *TableDaoImpl) innerGenUUID(item any) lang.UUID {
	ser := lang.DefaultUUIDService()
	b := ser.NewBuilder()
	b.ForObject(item)
	return b.Build()
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
	uuid := inst.innerGenUUID(item)

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

func (inst *TableDaoImpl) _impl() rbac.TableDAO {
	return inst
}
