package tables

import "gorm.io/gorm"

type DAO interface {
	Find(db *gorm.DB, id ID) (*Entity, error)

	Query(db *gorm.DB, q *Query) ([]*Entity, error)

	Insert(db *gorm.DB, item *Entity) (*Entity, error)
}
