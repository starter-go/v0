package iauthx

import (
	"fmt"

	"github.com/starter-go/v0/rbac-web-app/app/classes/authx"
)

type AuthorizationServiceImpl struct {

	//starter:component

	_as func(authx.AuthorizationService) //starter:as("#")

	RawProviderList []authx.Authorizer //starter:inject(".")

}

func (inst *AuthorizationServiceImpl) _impl() authx.AuthorizationService {
	return inst
}

func (inst *AuthorizationServiceImpl) Authorize(ctx *authx.AuthorizationContext) error {

	return fmt.Errorf("no impl ")

}
