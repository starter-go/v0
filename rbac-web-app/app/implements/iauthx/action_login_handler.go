package iauthx

import (
	"context"
	"fmt"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/v0/rbac-web-app/app/classes/authx"
	"github.com/starter-go/v0/rbac-web-app/app/data/dxo"
	"github.com/starter-go/v0/rbac-web-app/app/web/dto"
	"github.com/starter-go/v0/subjects"
)

type ActionLoginAuthorizer struct {

	//starter:component

	_as func(authx.Authorizer) //starter:as(".")

	SessionSpanSecondsDef int64 //starter:inject("${security.session.default-age-sec}")
	SessionSpanSecondsMax int64 //starter:inject("${security.session.max-age-sec}")
	SessionSpanSecondsMin int64 //starter:inject("${security.session.min-age-sec}")

}

// Accept implements authx.Authorizer.
func (inst *ActionLoginAuthorizer) Accept(ctx *authx.AuthorizationContext) bool {
	action := ctx.Auth.Action
	return (action == dxo.AuthActionLogin)
}

// Authorize implements authx.Authorizer.
func (inst *ActionLoginAuthorizer) Authorize(ctx *authx.AuthorizationContext) error {

	if !inst.Accept(ctx) {
		return fmt.Errorf("the authx.AuthorizationContext is not a login action")
	}

	// 创建新的会话,并返回它的token
	maker := new(innerActionLoginAuthorizerSessionMaker)
	maker.caller = inst
	maker.user = ctx.AuthUser
	maker.context = ctx.Context

	return maker.makeAll()
}

// GetRegistration implements authx.Authorizer.
func (inst *ActionLoginAuthorizer) GetRegistration() *authx.AuthorizerRegistration {

	r1 := new(authx.AuthorizerRegistration)
	r1.Name = "Action-Login-Authorizer"
	r1.Priority = 0
	r1.Enabled = true
	r1.Authorizer = inst
	r1.Action = dxo.AuthActionLogin

	return r1
}

func (inst *ActionLoginAuthorizer) _impl() authx.Authorizer {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

type innerActionLoginAuthorizerSessionMaker struct {
	caller  *ActionLoginAuthorizer
	context context.Context
	session *dto.Session
	user    *dto.User
}

func (inst *innerActionLoginAuthorizerSessionMaker) computeNotAfter(notBefore lang.Time) lang.Time {

	cl := inst.caller
	sec := cl.SessionSpanSecondsDef
	max := cl.SessionSpanSecondsMax
	min := cl.SessionSpanSecondsMin

	if sec < min {
		sec = min
	}
	if sec > max {
		sec = max
	}

	ms := sec * 1000
	return notBefore + lang.Time(ms)
}

func (inst *innerActionLoginAuthorizerSessionMaker) makeAll() error {

	ctx := inst.context
	sub, err := subjects.GetCurrent(ctx)
	if err != nil {
		return err
	}

	sett, err := sub.DoSet()
	if err != nil {
		return err
	}

	src := inst.user
	roles := src.Roles.Format()
	userID := src.ID
	now := lang.Now()
	notAfter := inst.computeNotAfter(now)

	if userID < 1 {
		return fmt.Errorf("bad user-id")
	}

	sett.SetUserID(userID)
	sett.SetAvatar(string(src.Avatar))
	sett.SetDisplayName(src.DisplayName)
	sett.SetUserEmail(src.Email.String())
	sett.SetUserName(src.Name)
	sett.SetLocale(src.Language)
	sett.SetAuthenticated(true)
	sett.SetNotBefore(now - 3)
	sett.SetNotAfter(notAfter)

	sett.SetProperty(subjects.PNameRoles, string(roles))

	err = sub.Create()
	if err != nil {
		return err
	}
	return sub.Flush()
}
