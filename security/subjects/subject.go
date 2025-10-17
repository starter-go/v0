package subjects

import (
	"context"
	"fmt"

	"github.com/starter-go/rbac"
	"github.com/starter-go/v0/security/sessions"
	"github.com/starter-go/v0/security/tokens"
)

type Subject struct {
	Context *Context

	JWT tokens.JWT

	Token *tokens.TokenDTO

	Session *sessions.SessionDTO

	// cached fileds of session
	UID rbac.UserID
}

func GetSubject(cc context.Context) (*Subject, error) {

	return nil, fmt.Errorf("todo .... no impl")
}
