package iroles

import (
	"fmt"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/security/random"
	"github.com/starter-go/v0/rbac-web-app/app/classes/roles"
	"github.com/starter-go/v0/rbac-web-app/app/data/database"
	"github.com/starter-go/v0/rbac-web-app/app/data/entity"
	"gorm.io/gorm"
)

type RoleDaoImpl struct {

	//starter:component

	_as func(roles.DAO) //starter:as("#")

	Agent   database.Agent     //starter:inject("#")
	UUIDSer random.UUIDService //starter:inject("#")

}

// Insert implements roles.DAO.
func (inst *RoleDaoImpl) Insert(db *gorm.DB, item *roles.Entity) (*roles.Entity, error) {

	item.ID = 0
	item.UUID = inst.innerGenUUID()
	db = inst.innerPrepareDB(db)

	res := db.Create(item)
	err := res.Error

	return item, err
}

func (inst *RoleDaoImpl) innerPrepareDB(db *gorm.DB) *gorm.DB {
	return inst.Agent.DB(db)
}

func (inst *RoleDaoImpl) innerMakeItemList() []*roles.Entity {
	return make([]*roles.Entity, 0)
}

func (inst *RoleDaoImpl) innerMakeItem() *roles.Entity {
	return new(roles.Entity)
}

func (inst *RoleDaoImpl) innerGenUUID() lang.UUID {
	b := inst.UUIDSer.Build()
	b.Class("roles.Entity")
	return b.Generate()
}

// Query implements roles.DAO.
func (inst *RoleDaoImpl) Query(db *gorm.DB, q *roles.Query) ([]*roles.Entity, error) {

	if q == nil {
		return nil, fmt.Errorf("RoleDaoImpl.Query() : query is nil")
	}

	db = inst.innerPrepareDB(db)
	finder := new(entity.Finder)
	list := inst.innerMakeItemList()
	item := inst.innerMakeItem()
	page := &q.Pagination
	want := q.Want

	finder.SetAll(q.All).SetDB(db).SetItem(item).SetList(&list).SetPagination(page).SetWant(want)

	err := finder.Find()
	return list, err
}

// Find implements roles.DAO.
func (inst *RoleDaoImpl) Find(db *gorm.DB, id roles.ID) (*roles.Entity, error) {

	db = inst.innerPrepareDB(db)
	item := inst.innerMakeItem()

	res := db.First(item, id)
	err := res.Error

	return item, err
}

func (inst *RoleDaoImpl) _impl() roles.DAO {
	return inst
}
