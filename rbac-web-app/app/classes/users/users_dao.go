package users

import (
	"github.com/starter-go/v0/rbac-web-app/app/data/dxo"
	"github.com/starter-go/v0/rbac-web-app/app/data/entity"
	"gorm.io/gorm"
)

type DAO interface {

	// query

	Find(db *gorm.DB, id dxo.UserID) (*entity.User, error)

	FindByName(db *gorm.DB, name dxo.UserName) (*entity.User, error)

	FindByMobile(db *gorm.DB, num dxo.PhoneNumber) (*entity.User, error)

	FindByEmail(db *gorm.DB, addr dxo.EmailAddress) (*entity.User, error)

	// edit

	Insert(db *gorm.DB, item *entity.User) (*entity.User, error)
}
