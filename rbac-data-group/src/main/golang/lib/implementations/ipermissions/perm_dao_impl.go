package ipermissions

import (
	"fmt"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/api/classes/permissions"

	"github.com/starter-go/v0/libdao"
	"github.com/starter-go/v0/libdao/api/libdaoapi"
	"github.com/starter-go/v0/rbac-data-group/src/main/golang/api/daos"
	"gorm.io/gorm"
)

type PermissionDaoImpl struct {

	//starter:component

	_as func(rbac.PermissionDAO) //starter:as(".")

	ConfigClass    string //starter:inject("${rbac-data-group.sql.class}")
	ConfigEnabled  bool   //starter:inject("${rbac-data-group.sql.enabled}")
	ConfigPriority int    //starter:inject("${rbac-data-group.sql.priority}")

	DBAgent daos.IDatabaseAgent //starter:inject("#")

}

// GetRegistration implements [permissions.DAO].
func (inst *PermissionDaoImpl) GetRegistration() *libdaoapi.DaoRegistration {

	r1 := &libdao.DaoRegistration{
		Name:     "PermissionDaoImpl",
		ID:       "sql-rbac-permission-dao",
		Class:    inst.ConfigClass,
		Enabled:  inst.ConfigEnabled,
		Priority: inst.ConfigPriority,
		DAO:      inst,
	}

	return r1
}

func (inst *PermissionDaoImpl) innerGenUUID(item any) lang.UUID {
	ser := lang.DefaultUUIDService()
	b := ser.NewBuilder()
	b.ForObject(item)
	return b.Build()
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
	uuid := inst.innerGenUUID(item)

	item.ID = 0
	item.UUID = uuid

	res := db.Create(item)
	err := res.Error
	return item, err
}

// Query implements [permissions.UserDAO].
func (inst *PermissionDaoImpl) Query(db *gorm.DB, q *permissions.Query) ([]*permissions.Entity, error) {

	db = inst.GetDB(db)
	finder := new(rbac.Finder)
	list := inst.innerMakeItemList()
	m := inst.innerMakeItem()
	p := &q.Pagination

	finder.SetDB(db).SetPagination(p).SetAll(q.All)
	finder.SetList(&list).SetWant(q.Want).SetModel(m)

	err := finder.Find()
	return list, err
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

func (inst *PermissionDaoImpl) _impl() rbac.PermissionDAO {
	return inst
}
