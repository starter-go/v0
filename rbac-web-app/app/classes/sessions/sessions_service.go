package sessions

import (
	"context"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/v0/rbac-web-app/app/data/dxo"
	"github.com/starter-go/v0/rbac-web-app/app/web/dto"
)

type Service interface {

	// query

	Find(cc context.Context, id dxo.SessionID) (*dto.Session, error)

	FindByUUID(cc context.Context, uuid lang.UUID) (*dto.Session, error)

	// edit

	Insert(cc context.Context, item *dto.Session) (*dto.Session, error)

	Update(cc context.Context, id dxo.SessionID, item *dto.Session) (*dto.Session, error)

	Remove(cc context.Context, id dxo.SessionID) error
}
