package iauths

import (
	"fmt"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/api/classes/authentications"
	"github.com/starter-go/v0/libdao"
	"github.com/starter-go/v0/libdao/api/libdaoapi"
	"github.com/starter-go/v0/rbac-data-group/src/main/golang/api/daos"
	"gorm.io/gorm"
)

type AuthenticationDaoImpl struct {

	//starter:component

	_as func(rbac.AuthenticationDAO) //starter:as(".")

	ConfigClass    string //starter:inject("${rbac-data-group.sql.class}")
	ConfigEnabled  bool   //starter:inject("${rbac-data-group.sql.enabled}")
	ConfigPriority int    //starter:inject("${rbac-data-group.sql.priority}")

	DBAgent daos.IDatabaseAgent //starter:inject("#")

}

// GetRegistration implements [authentications.DAO].
func (inst *AuthenticationDaoImpl) GetRegistration() *libdaoapi.DaoRegistration {

	r1 := &libdao.DaoRegistration{
		Name:     "AuthenticationDaoImpl",
		ID:       "sql-rbac-authentication-dao",
		Class:    inst.ConfigClass,
		Enabled:  inst.ConfigEnabled,
		Priority: inst.ConfigPriority,
		DAO:      inst,
	}

	return r1

}

func (inst *AuthenticationDaoImpl) innerGenUUID(item any) lang.UUID {
	ser := lang.DefaultUUIDService()
	b := ser.NewBuilder()
	b.ForObject(item)
	return b.Build()
}

func (inst *AuthenticationDaoImpl) innerMakeItem() *authentications.Entity {
	return new(authentications.Entity)
}

func (inst *AuthenticationDaoImpl) innerMakeItemList() []*authentications.Entity {
	return make([]*authentications.Entity, 0)
}

// Delete implements [authentications.UserDAO].
func (inst *AuthenticationDaoImpl) Delete(db *gorm.DB, id authentications.ID) error {
	db = inst.GetDB(db)
	item := inst.innerMakeItem()
	item.ID = id
	res := db.Delete(item, id)
	err := res.Error
	return err
}

// Find implements [authentications.UserDAO].
func (inst *AuthenticationDaoImpl) Find(db *gorm.DB, id authentications.ID) (*authentications.Entity, error) {
	db = inst.GetDB(db)
	item := inst.innerMakeItem()
	res := db.First(item, id)
	err := res.Error
	return item, err
}

// GetDB implements [authentications.UserDAO].
func (inst *AuthenticationDaoImpl) GetDB(old *gorm.DB) *gorm.DB {
	return inst.DBAgent.DB(old)
}

// Insert implements [authentications.UserDAO].
func (inst *AuthenticationDaoImpl) Insert(db *gorm.DB, item *authentications.Entity) (*authentications.Entity, error) {

	db = inst.GetDB(db)
	uuid := inst.innerGenUUID(item)

	item.ID = 0
	item.UUID = uuid

	res := db.Create(item)
	err := res.Error
	return item, err
}

// Query implements [authentications.UserDAO].
func (inst *AuthenticationDaoImpl) Query(db *gorm.DB, q *authentications.Query) ([]*authentications.Entity, error) {

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

// Update implements [authentications.UserDAO].
func (inst *AuthenticationDaoImpl) Update(db *gorm.DB, id authentications.ID, callback func(old *authentications.Entity) error) (*authentications.Entity, error) {

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

func (inst *AuthenticationDaoImpl) _impl() rbac.AuthenticationDAO {
	return inst
}
