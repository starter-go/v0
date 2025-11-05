package iauthx

import (
	"fmt"

	"github.com/starter-go/v0/rbac-web-app/app/classes/authx"
)

type AuthenticationServiceImpl struct {

	//starter:component

	_as func(authx.AuthenticationService) //starter:as("#")

	RawProviderList []authx.Authenticator //starter:inject(".")

	cached []*authx.AuthenticatorRegistration
}

func (inst *AuthenticationServiceImpl) _impl_1() authx.AuthenticationService {
	return inst
}

func (inst *AuthenticationServiceImpl) getProviderList() []*authx.AuthenticatorRegistration {
	list := inst.cached
	if list == nil {
		list = inst.loadProviderList()
		inst.cached = list
	}
	return list
}

func (inst *AuthenticationServiceImpl) loadProviderList() []*authx.AuthenticatorRegistration {

	src := inst.RawProviderList
	dst := make([]*authx.AuthenticatorRegistration, 0)
	for _, it1 := range src {
		if it1 == nil {
			continue
		}
		it2 := it1.GetRegistration()
		if inst.isProviderReady(it2) {
			dst = append(dst, it2)
		}
	}
	return dst
}

func (inst *AuthenticationServiceImpl) isProviderReady(p *authx.AuthenticatorRegistration) bool {

	if p == nil {
		return false
	}

	if !p.Enabled {
		return false
	}

	if p.Authenticator == nil {
		return false
	}
	if p.Name == "" {
		return false
	}
	if p.Mechanism == "" {
		return false
	}

	return true
}

func (inst *AuthenticationServiceImpl) findProviderByMechanism(ctx *authx.AuthenticationContext) (authx.Authenticator, error) {

	want := ctx.Auth.Mechanism
	all := inst.getProviderList()

	for _, item := range all {
		have := item.Mechanism
		if have == want {
			return item.Authenticator, nil
		}
	}

	return nil, fmt.Errorf("no Authenticator found")
}

func (inst *AuthenticationServiceImpl) findProviderByAccept(ctx *authx.AuthenticationContext) (authx.Authenticator, error) {

	all := inst.getProviderList()

	for _, item := range all {
		p := item.Authenticator
		if p.Accept(ctx) {
			return p, nil
		}
	}

	return nil, fmt.Errorf("no Authenticator found")
}

func (inst *AuthenticationServiceImpl) findProvider(ctx *authx.AuthenticationContext) (authx.Authenticator, error) {

	provider, err := inst.findProviderByMechanism(ctx)
	if (provider != nil) && (err == nil) {
		return provider, nil
	}

	provider, err = inst.findProviderByAccept(ctx)
	if (provider != nil) && (err == nil) {
		return provider, nil
	}

	mech := ctx.Auth.Mechanism
	return nil, fmt.Errorf("no Authenticator handled the auth request, Mechanism=[%s]", mech)
}

func (inst *AuthenticationServiceImpl) Authenticate(ctx *authx.AuthenticationContext) error {

	provider, err := inst.findProvider(ctx)
	if err != nil {
		return err
	}
	return provider.Authenticate(ctx)
}
