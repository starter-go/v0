package sessions

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/v0/rbac-web-app/app/data/dxo"
	"github.com/starter-go/v0/rbac-web-app/app/data/entity"
	"gorm.io/gorm"
)

type DAO interface {

	// query

	Find(db *gorm.DB, id dxo.SessionID) (*entity.Session, error)

	FindByUUID(db *gorm.DB, uuid lang.UUID) (*entity.Session, error)

	// edit

	Insert(db *gorm.DB, item *entity.Session) (*entity.Session, error)

	Update(db *gorm.DB, id dxo.SessionID, callback func(item *entity.Session) error) (*entity.Session, error)

	Remove(db *gorm.DB, id dxo.SessionID) error
}
