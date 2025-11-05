package sessions

import (
	"context"

	"github.com/starter-go/v0/rbac-web-app/app/data/dxo"
	"github.com/starter-go/v0/rbac-web-app/app/web/dto"
)

type Service interface {

	// query

	Find(cc context.Context, id dxo.SessionID) (*dto.Session, error)

	// edit

	Insert(cc context.Context, item *dto.Session) (*dto.Session, error)
}
