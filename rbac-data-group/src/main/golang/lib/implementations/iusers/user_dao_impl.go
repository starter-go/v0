package iusers

import (
	"fmt"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/api/classes/users"

	"github.com/starter-go/v0/libdao"
	"github.com/starter-go/v0/libdao/api/libdaoapi"
	"github.com/starter-go/v0/rbac-data-group/src/main/golang/api/daos"
	"gorm.io/gorm"
)

type UserDaoImpl struct {

	//starter:component

	_as func(rbac.UserDAO) //starter:as(".")

	ConfigClass    string //starter:inject("${rbac-data-group.sql.class}")
	ConfigEnabled  bool   //starter:inject("${rbac-data-group.sql.enabled}")
	ConfigPriority int    //starter:inject("${rbac-data-group.sql.priority}")

	DBAgent daos.IDatabaseAgent //starter:inject("#")

}

// FindByEmail implements [users.UserDAO].
func (inst *UserDaoImpl) FindByEmail(db *gorm.DB, addr users.EmailAddress) (*users.Entity, error) {
	panic("unimplemented")
}

// FindByName implements [users.UserDAO].
func (inst *UserDaoImpl) FindByName(db *gorm.DB, name users.UserName) (*users.Entity, error) {
	panic("unimplemented")
}

// FindByPhone implements [users.UserDAO].
func (inst *UserDaoImpl) FindByPhone(db *gorm.DB, num users.PhoneNumber) (*users.Entity, error) {
	panic("unimplemented")
}

// GetRegistration implements [users.UserDAO].
func (inst *UserDaoImpl) GetRegistration() *libdaoapi.DaoRegistration {
	r1 := &libdao.DaoRegistration{
		Name:     "UserDaoImpl",
		ID:       "sql-rbac-user-dao",
		Class:    inst.ConfigClass,
		Enabled:  inst.ConfigEnabled,
		Priority: inst.ConfigPriority,
		DAO:      inst,
	}
	return r1
}

func (inst *UserDaoImpl) innerGenUUID(item any) lang.UUID {
	ser := lang.DefaultUUIDService()
	b := ser.NewBuilder()
	b.ForObject(item)
	return b.Build()
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
	uuid := inst.innerGenUUID(item)

	item.ID = 0
	item.UUID = uuid

	res := db.Create(item)
	err := res.Error
	return item, err
}

// Query implements [users.UserDAO].
func (inst *UserDaoImpl) Query(db *gorm.DB, q *users.Query) ([]*users.Entity, error) {

	db = inst.GetDB(db)
	finder := new(rbac.Finder)
	list := inst.innerMakeItemList()
	p := &q.Pagination
	m := inst.innerMakeItem()

	finder.SetDB(db).SetPagination(p).SetAll(q.All)
	finder.SetList(&list).SetWant(q.Want).SetModel(m)

	err := finder.Find()
	return list, err
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

func (inst *UserDaoImpl) _impl() rbac.UserDAO {
	return inst
}
