package tokens

import (
	"context"

	"github.com/starter-go/v0/rbac-web-app/app/web/dto"
)

type Service interface {

	// query

	GetCurrentToken(cc context.Context) (*dto.Token, error)

	LoadCurrentToken(cc context.Context) (*dto.Token, error)

	// edit

	SetCurrentToken(cc context.Context, token *dto.Token) (*dto.Token, error)
}
