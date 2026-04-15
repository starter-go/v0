package iauthx

import (
	"context"
	"fmt"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/v0/rbac-web-app/app/classes/authx"
	"github.com/starter-go/v0/rbac-web-app/app/classes/sessions"
	"github.com/starter-go/v0/rbac-web-app/app/classes/tokens"
	"github.com/starter-go/v0/rbac-web-app/app/data/dxo"
	"github.com/starter-go/v0/rbac-web-app/app/web/dto"
)

type ActionLoginAuthorizer struct {

	//starter:component

	_as func(authx.Authorizer) //starter:as(".")

	SessionService sessions.Service //starter:inject("#")
	TokenService   tokens.Service   //starter:inject("#")
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

func (inst *innerActionLoginAuthorizerSessionMaker) makeSession() error {

	ctx := inst.context
	ser := inst.caller.SessionService
	now := lang.Now()
	user := inst.user
	session := new(dto.Session)

	session.Owner = user.ID
	session.AliveFrom = now - 1000
	session.AliveTo = now + (24 * 3600 * 1000)
	session.Alive = true
	session.DisplayName = user.DisplayName
	session.Avatar = user.Avatar
	session.Roles = user.Roles

	ses2, err := ser.Insert(ctx, session)
	if err != nil {
		return err
	}

	inst.session = ses2
	return nil
}

func (inst *innerActionLoginAuthorizerSessionMaker) makeToken() error {

	ctx := inst.context
	ser := inst.caller.TokenService
	now := lang.Now()

	session := inst.session
	token := new(dto.Token)

	token.AliveFrom = now - 1000
	token.AliveTo = now + (1000 * 3600)
	token.SessionID = session.ID
	token.SessionUUID = session.UUID

	_, err := ser.SetCurrentToken(ctx, token)
	return err
}

func (inst *innerActionLoginAuthorizerSessionMaker) makeAll() error {

	err := inst.makeSession()
	if err != nil {
		return err
	}

	err = inst.makeToken()
	if err != nil {
		return err
	}

	return nil
}
