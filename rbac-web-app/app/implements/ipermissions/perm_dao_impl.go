package ipermissions

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/security/random"
	"github.com/starter-go/v0/rbac-web-app/app/classes/permissions"
	"github.com/starter-go/v0/rbac-web-app/app/data/database"
	"github.com/starter-go/v0/rbac-web-app/app/data/entity"
	"gorm.io/gorm"
)

type PermissionDaoImpl struct {

	//starter:component

	_as func(permissions.DAO) //starter:as("#")

	Agent   database.Agent     //starter:inject("#")
	UUIDSer random.UUIDService //starter:inject("#")

}

func (inst *PermissionDaoImpl) innerPrepareDB(db *gorm.DB) *gorm.DB {
	return inst.Agent.DB(db)
}

func (inst *PermissionDaoImpl) innerMakeItem() *permissions.Entity {
	return new(permissions.Entity)
}

func (inst *PermissionDaoImpl) innerMakeItemList() []*permissions.Entity {
	return make([]*permissions.Entity, 0)
}

func (inst *PermissionDaoImpl) innerGenUUID() lang.UUID {
	b := inst.UUIDSer.Build()
	b.Class("permissions.Entity")
	return b.Generate()
}

// Query implements permissions.DAO.
func (inst *PermissionDaoImpl) Query(db *gorm.DB, q *permissions.Query) ([]*permissions.Entity, error) {

	db = inst.innerPrepareDB(db)

	finder := new(entity.Finder)
	item := inst.innerMakeItem()
	list := inst.innerMakeItemList()

	finder.SetDB(db).SetItem(item).SetList(&list).SetWant(q.Want).SetPagination(&q.Pagination)

	err := finder.Find()
	return list, err
}

// Find implements permissions.DAO.
func (inst *PermissionDaoImpl) Find(db *gorm.DB, id permissions.ID) (*permissions.Entity, error) {
	db = inst.innerPrepareDB(db)
	item := inst.innerMakeItem()
	res := db.First(item, id)
	err := res.Error
	return item, err
}

// Insert implements permissions.DAO.
func (inst *PermissionDaoImpl) Insert(db *gorm.DB, item *permissions.Entity) (*permissions.Entity, error) {

	item.ID = 0
	item.UUID = inst.innerGenUUID()

	db = inst.innerPrepareDB(db)
	res := db.Create(item)
	err := res.Error
	return item, err
}

func (inst *PermissionDaoImpl) _impl() permissions.DAO {
	return inst
}
