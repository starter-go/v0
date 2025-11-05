package subjects

import (
	"context"

	"github.com/starter-go/v0/rbac-web-app/app/classes/sessions"
	"github.com/starter-go/v0/rbac-web-app/app/classes/tokens"
	"github.com/starter-go/v0/rbac-web-app/app/classes/users"
	"github.com/starter-go/v0/rbac-web-app/app/web/dto"
)

////////////////////////////////////////////////////////////////////////////////
// context

type Services struct {
	Tokens tokens.Service

	Sessions sessions.Service

	Users users.Service
}

type Cache struct {
	Token *dto.Token

	Session *dto.Session

	User *dto.User
}

type Context struct {

	// parent-context
	Parent context.Context

	// services
	Services Services

	// cache
	Cache Cache

	// facade
	Subject Subject
}

func NewContext() *Context {

	ctx := new(Context)
	sub := new(innerSubjectFacade)

	ctx.Subject = sub
	sub.context = ctx

	return ctx
}

////////////////////////////////////////////////////////////////////////////////
// EOF
