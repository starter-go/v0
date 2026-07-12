package users

import (
	"github.com/starter-go/v0/rbac-web-app/app/data/dxo"
	"gorm.io/gorm"
)

type DAO interface {

	// query

	Find(db *gorm.DB, id ID) (*Entity, error)

	FindByName(db *gorm.DB, name dxo.UserName) (*Entity, error)

	FindByMobile(db *gorm.DB, num dxo.PhoneNumber) (*Entity, error)

	FindByEmail(db *gorm.DB, addr dxo.EmailAddress) (*Entity, error)

	Query(db *gorm.DB, q *Query) ([]*Entity, error)

	// edit

	Insert(db *gorm.DB, item *Entity) (*Entity, error)
}
