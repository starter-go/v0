package sessions

import "gorm.io/gorm"

type DAO interface {

	// prepare

	PrepareDB(db *gorm.DB) *gorm.DB

	// query

	Query(db *gorm.DB, q *Query) ([]*Entity, error)

	Find(db *gorm.DB, id ID) (*Entity, error)

	// edit

	Insert(db *gorm.DB, item *Entity) (*Entity, error)

	Update(db *gorm.DB, id ID, callback func(item *Entity) error) (*Entity, error)

	Remove(db *gorm.DB, id ID) error
}
