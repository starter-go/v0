package subjects

import (
	"context"
	"fmt"

	"github.com/starter-go/v0/rbac-web-app/app/classes/webcontexts"
	"github.com/starter-go/v0/rbac-web-app/app/web/dto"
)

////////////////////////////////////////////////////////////////////////////////
// api

const (
	theSubjectAttrKey = "github.com/starter-go/v0/rbac-web-app/app/classes/subjects#Subject"
)

type Subject interface {
	GetToken() (*dto.Token, error)

	GetSession() (*dto.Session, error)

	GetUser() (*dto.User, error)

	GetContext() *Context
}

func Current(cc context.Context) (Subject, error) {

	const key = theSubjectAttrKey

	acc, err := webcontexts.GetAccess(cc)
	if err != nil {
		return nil, err
	}

	value := acc.GetAttr(key)
	if value != nil {
		sub, ok := value.(Subject)
		if ok && (sub != nil) {
			return sub, nil
		}
	}

	ctx := NewContext()
	sub := ctx.Subject
	acc.SetAttr(key, sub)
	return sub, nil
}

////////////////////////////////////////////////////////////////////////////////
// facade

type innerSubjectFacade struct {
	context *Context
}

func (inst *innerSubjectFacade) _impl() Subject {
	return inst
}

func (inst *innerSubjectFacade) innerLoadToken() (*dto.Token, error) {

	ctx := inst.context.Parent
	ser := inst.context.Services.Tokens

	return ser.LoadCurrentToken(ctx)
}

func (inst *innerSubjectFacade) innerLoadSession() (*dto.Session, error) {

	token, err := inst.GetToken()
	if err != nil {
		return nil, err
	}

	ctx := inst.context.Parent
	ser := inst.context.Services.Sessions
	id := token.SessionID

	return ser.Find(ctx, id)
}

func (inst *innerSubjectFacade) innerLoadUserInfo() (*dto.User, error) {

	ses, err := inst.GetSession()
	if err != nil {
		return nil, err
	}

	ctx := inst.context.Parent
	ser := inst.context.Services.Users
	id := ses.Owner

	return ser.Find(ctx, id)
}

func (inst *innerSubjectFacade) GetToken() (*dto.Token, error) {

	token := inst.context.Cache.Token

	if token == nil {
		obj, err := inst.innerLoadToken()
		if err != nil {
			return nil, err
		} else if obj == nil {
			return nil, fmt.Errorf("innerSubjectFacade.GetToken(): token is nil")
		} else {
			token = obj // ok
		}
		inst.context.Cache.Token = token
	}

	return token, nil
}

func (inst *innerSubjectFacade) GetSession() (*dto.Session, error) {

	session := inst.context.Cache.Session

	if session == nil {
		obj, err := inst.innerLoadSession()
		if err != nil {
			return nil, err
		} else if obj == nil {
			return nil, fmt.Errorf("innerSubjectFacade.GetSession(): session is nil")
		} else {
			session = obj // ok
		}
		inst.context.Cache.Session = session
	}

	return session, nil
}

func (inst *innerSubjectFacade) GetUser() (*dto.User, error) {

	user := inst.context.Cache.User

	if user == nil {
		obj, err := inst.innerLoadUserInfo()
		if err != nil {
			return nil, err
		} else if obj == nil {
			return nil, fmt.Errorf("innerSubjectFacade.GetUser(): user is nil")
		} else {
			user = obj // ok
		}
		inst.context.Cache.User = user
	}

	return user, nil
}

func (inst *innerSubjectFacade) GetContext() *Context {
	return inst.context
}

////////////////////////////////////////////////////////////////////////////////
// EOF
