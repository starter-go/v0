package isessions

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/security/random"
	"github.com/starter-go/v0/subjects/core/classes/sessions"
	"github.com/starter-go/v0/subjects/core/data/subjectdb"
	"gorm.io/gorm"
)

type SessionDaoImpl struct {

	//starter:component

	_as func(sessions.DAO) //starter:as("#")

	Agent   subjectdb.Agent    //starter:inject("#")
	UUIDSer random.UUIDService //starter:inject("#")

}

func (inst *SessionDaoImpl) innerMakeItem() *sessions.Entity {
	return new(sessions.Entity)
}

func (inst *SessionDaoImpl) innerMakeItemList() []*sessions.Entity {
	return make([]*sessions.Entity, 0)
}

func (inst *SessionDaoImpl) innerGenUUID() lang.UUID {
	ser := inst.UUIDSer
	b := ser.Build().Class("sessions.Entity")
	return b.Generate()
}

// PrepareDB implements sessions.DAO.
func (inst *SessionDaoImpl) PrepareDB(db *gorm.DB) *gorm.DB {
	return inst.innerPrepareDB(db)
}

func (inst *SessionDaoImpl) innerPrepareDB(db *gorm.DB) *gorm.DB {
	return inst.Agent.DB(db)
}

// Find implements sessions.DAO.
func (inst *SessionDaoImpl) Find(db *gorm.DB, id sessions.ID) (*sessions.Entity, error) {

	db = inst.innerPrepareDB(db)
	item := inst.innerMakeItem()
	res := db.Find(&item, id)
	err := res.Error
	return item, err
}

// Insert implements sessions.DAO.
func (inst *SessionDaoImpl) Insert(db *gorm.DB, item *sessions.Entity) (*sessions.Entity, error) {

	uuid := inst.innerGenUUID()
	db = inst.innerPrepareDB(db)

	item.ID = 0
	item.UUID = uuid

	res := db.Create(item)
	err := res.Error

	return item, err
}

// Query implements sessions.DAO.
func (inst *SessionDaoImpl) Query(db *gorm.DB, q *sessions.Query) ([]*sessions.Entity, error) {

	item := inst.innerMakeItem()
	list := inst.innerMakeItemList()
	want := q.Want
	page := q.Pagination

	var count int64

	db = inst.innerPrepareDB(db)

	// from
	db = db.Model(item)

	// where
	if want != nil {
		db = db.Where(want)
	}

	// count
	res := db.Count(&count)
	if res.Error == nil {
		q.Pagination.Total = count
	}

	// page
	db = db.Limit(int(page.Limit())).Offset(int(page.Offset()))

	// find
	res = db.Find(&list)
	if res.Error != nil {
		return nil, res.Error
	}

	return list, nil
}

// Remove implements sessions.DAO.
func (inst *SessionDaoImpl) Remove(db *gorm.DB, id sessions.ID) error {

	item := inst.innerMakeItem()
	db = inst.innerPrepareDB(db)
	res := db.Delete(item, id)
	return res.Error
}

// Update implements sessions.DAO.
func (inst *SessionDaoImpl) Update(db *gorm.DB, id sessions.ID, callback func(item *sessions.Entity) error) (*sessions.Entity, error) {

	db = inst.innerPrepareDB(db)

	item1, err := inst.Find(db, id)
	if err != nil {
		return nil, err
	}

	err = callback(item1)
	if err != nil {
		return nil, err
	}

	res := db.Save(item1)
	err = res.Error
	if err != nil {
		return nil, err
	}

	return item1, nil
}

func (inst *SessionDaoImpl) _impl() sessions.DAO {
	return inst
}
