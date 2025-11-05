package sessions

import (
	"github.com/starter-go/v0/rbac-web-app/app/data/dxo"
	"github.com/starter-go/v0/rbac-web-app/app/data/entity"
	"gorm.io/gorm"
)

type DAO interface {

	// query

	Find(db *gorm.DB, id dxo.SessionID) (*entity.Session, error)

	// edit

	Insert(db *gorm.DB, item *entity.Session) (*entity.Session, error)
}
