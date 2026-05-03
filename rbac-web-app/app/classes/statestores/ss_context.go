package statestores

import (
	"context"
	"time"

	"github.com/starter-go/application/properties"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/v0/rbac-web-app/app/data/dxo"
	"github.com/starter-go/v0/rbac-web-app/app/web/dto"
)

type Context struct {
	CC context.Context

	Method Method

	SessionID   dxo.SessionID
	SessionUUID lang.UUID

	Properties properties.Table

	// dto(s)

	User *dto.User

	Session *dto.Session

	Token *dto.Token

	SessionMaxAge time.Duration

	TokenMaxAge time.Duration
}
