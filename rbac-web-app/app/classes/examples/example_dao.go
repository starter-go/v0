package examples

import "gorm.io/gorm"

type DAO interface {
	Find(db *gorm.DB, id ID) (*Entity, error)
}
