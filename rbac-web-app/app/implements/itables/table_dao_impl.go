package itables

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/security/random"
	"github.com/starter-go/v0/rbac-web-app/app/classes/tables"
	"github.com/starter-go/v0/rbac-web-app/app/data/database"
	"github.com/starter-go/v0/rbac-web-app/app/data/entity"
	"gorm.io/gorm"
)

type TableDaoImpl struct {

	//starter:component

	_as func(tables.DAO) //starter:as("#")

	Agent   database.Agent     //starter:inject("#")
	UUIDSer random.UUIDService //starter:inject("#")

}

// Find implements [tables.DAO].
func (inst *TableDaoImpl) Find(db *gorm.DB, id tables.ID) (*tables.Entity, error) {
	item := inst.innerMakeItem()
	db = inst.innerPrepareDB(db)
	res := db.First(item, id)
	err := res.Error
	return item, err
}

// Insert implements [tables.DAO].
func (inst *TableDaoImpl) Insert(db *gorm.DB, item *tables.Entity) (*tables.Entity, error) {
	item.ID = 0
	item.UUID = inst.innerGenUUID()
	db = inst.innerPrepareDB(db)
	res := db.Create(item)
	err := res.Error
	return item, err
}

// Query implements [tables.DAO].
func (inst *TableDaoImpl) Query(db *gorm.DB, q *tables.Query) ([]*tables.Entity, error) {

	finder := new(entity.Finder)
	item := inst.innerMakeItem()
	list := inst.innerMakeItemList()
	page := &q.Pagination
	want := q.Want

	db = inst.innerPrepareDB(db)
	finder.SetAll(q.All).SetDB(db).SetItem(item).SetList(&list).SetPagination(page)

	if want != nil {
		finder.SetWant(want)
	}

	err := finder.Find()
	return list, err
}

func (inst *TableDaoImpl) innerMakeItem() *tables.Entity {
	return new(tables.Entity)
}

func (inst *TableDaoImpl) innerMakeItemList() []*tables.Entity {
	return make([]*tables.Entity, 0)
}

func (inst *TableDaoImpl) innerPrepareDB(db *gorm.DB) *gorm.DB {
	return inst.Agent.DB(db)
}

func (inst *TableDaoImpl) innerGenUUID() lang.UUID {
	b := inst.UUIDSer.Build()
	b.Class("tables.Entity")
	return b.Generate()
}

func (inst *TableDaoImpl) _impl() tables.DAO {
	return inst
}
