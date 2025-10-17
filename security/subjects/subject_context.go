package subjects

import (
	"github.com/starter-go/v0/security/authx"
	"github.com/starter-go/v0/security/sessions"

	"github.com/starter-go/v0/security/tokens"
)

type Context struct {

	// services

	SubjectService Service

	AuthxService   authx.Service
	JWTService     tokens.JWTService
	SessionService sessions.Service
	TokenService   tokens.TokenService
}
