package subjects

import (
	"context"

	"github.com/starter-go/application/properties"
	"github.com/starter-go/v0/rbac-web-app/app/classes/sessions"
	"github.com/starter-go/v0/rbac-web-app/app/classes/statestores"
	"github.com/starter-go/v0/rbac-web-app/app/classes/tokens"
	"github.com/starter-go/v0/rbac-web-app/app/classes/users"
	"github.com/starter-go/v0/rbac-web-app/app/web/dto"
)

////////////////////////////////////////////////////////////////////////////////
// context

type Services struct {
	Tokens tokens.Service

	Sessions sessions.Service

	Subjects Service

	Users users.Service
}

type Cache struct {
	Token *dto.Token

	Session *dto.Session

	User *dto.User

	Properties properties.Table

	Dirty bool
}

type Context struct {

	// parent-context
	Parent context.Context

	// services
	Services Services

	// facade
	Subject Subject

	// store
	Store statestores.Store

	// cache
	Cache *Cache
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
