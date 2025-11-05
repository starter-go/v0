package iauthx

import (
	"context"

	"github.com/starter-go/v0/rbac-web-app/app/classes/authx"
	"github.com/starter-go/v0/rbac-web-app/app/web/vo"
)

type AuthxServiceImpl struct {

	//starter:component

	_as func(authx.Service) //starter:as("#")

	Auth1 authx.AuthenticationService //starter:inject("#")
	Auth2 authx.AuthorizationService  //starter:inject("#")

}

func (inst *AuthxServiceImpl) _impl() authx.Service {
	return inst
}

func (inst *AuthxServiceImpl) Auth(ctx context.Context, view *vo.Authx) (*vo.Authx, error) {

	tc := new(innerAuthxServiceTaskContext)
	tc.view1 = view
	tc.context = ctx

	err := inst.doAuth1(tc)
	if err != nil {
		return nil, err
	}

	err = inst.doAuth2(tc)
	if err != nil {
		return nil, err
	}

	return tc.view2, nil
}

func (inst *AuthxServiceImpl) doAuth1(x *innerAuthxServiceTaskContext) error {

	ser := inst.Auth1
	list := x.view1.Authentications
	actx := new(authx.AuthenticationContext)

	actx.Context = x.context
	actx.Auth = nil
	actx.View = x.view1

	for _, item := range list {
		actx.Auth = item
		err := ser.Authenticate(actx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *AuthxServiceImpl) doAuth2(x *innerAuthxServiceTaskContext) error {

	ser := inst.Auth2
	list := x.view1.Authorizations
	actx := new(authx.AuthorizationContext)

	actx.Context = x.context
	actx.Auth = nil
	actx.View = x.view1

	for _, item := range list {
		actx.Auth = item
		err := ser.Authorize(actx)
		if err != nil {
			return err
		}
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type innerAuthxServiceTaskContext struct {
	context context.Context
	view1   *vo.Authx // request
	view2   *vo.Authx // response
}
