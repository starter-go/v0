package iauthx

import (
	"fmt"
	"sort"
	"sync"

	"github.com/starter-go/v0/rbac-web-app/app/classes/authx"
	"github.com/starter-go/v0/rbac-web-app/app/data/dxo"
)

type AuthorizationServiceImpl struct {

	//starter:component

	_as func(authx.AuthorizationService) //starter:as("#")

	RawProviderList []authx.Authorizer //starter:inject(".")

	cache *innerAuthorizerCache
}

func (inst *AuthorizationServiceImpl) _impl() authx.AuthorizationService {
	return inst
}

func (inst *AuthorizationServiceImpl) innerLoadCache() *innerAuthorizerCache {
	cache := new(innerAuthorizerCache)
	cache.loadFrom(inst.RawProviderList)
	return cache
}

func (inst *AuthorizationServiceImpl) innerGetCache() *innerAuthorizerCache {
	cache := inst.cache
	if cache == nil {
		cache = inst.innerLoadCache()
		inst.cache = cache
	}
	return cache
}

func (inst *AuthorizationServiceImpl) findAuthorizersForContext(ctx *authx.AuthorizationContext) ([]*authx.AuthorizerRegistration, error) {

	action := ctx.Auth.Action
	cache := inst.innerGetCache()
	list := cache.queryWithAction(action)

	return list, nil
}

func (inst *AuthorizationServiceImpl) Authorize(ctx *authx.AuthorizationContext) error {

	all, err := inst.findAuthorizersForContext(ctx)
	if err != nil {
		return err
	}

	for _, item := range all {
		authorizer := item.Authorizer
		if authorizer.Accept(ctx) {
			return authorizer.Authorize(ctx)
		}
	}

	action := ctx.Auth.Action
	return fmt.Errorf("no Authorizer can handle the Authorization (action=%s)", action)
}

////////////////////////////////////////////////////////////////////////////////

type innerAuthorizerBank struct {
	action dxo.AuthAction

	list []*authx.AuthorizerRegistration
}

func (inst *innerAuthorizerBank) isItemReady(item *authx.AuthorizerRegistration) bool {

	if item == nil {
		return false
	}

	if item.Authorizer == nil {
		return false
	}

	if !item.Enabled {
		return false
	}

	if item.Action == "" {
		return false
	}

	return true
}

func (inst *innerAuthorizerBank) sort() {
	sort.Sort(inst)
}

func (inst *innerAuthorizerBank) Len() int {
	return len(inst.list)
}
func (inst *innerAuthorizerBank) Less(i1, i2 int) bool {
	all := inst.list
	o1 := all[i1]
	o2 := all[i2]
	return (o1.Priority < o2.Priority)
}
func (inst *innerAuthorizerBank) Swap(i1, i2 int) {
	all := inst.list
	all[i1], all[i2] = all[i2], all[i1]
}

func (inst *innerAuthorizerBank) add(item *authx.AuthorizerRegistration) {
	if inst.isItemReady(item) {
		inst.list = append(inst.list, item)
	}
}

////////////////////////////////////////////////////////////////////////////////

type innerAuthorizerCache struct {
	banks map[dxo.AuthAction]*innerAuthorizerBank
	mutex sync.Mutex
}

func (inst *innerAuthorizerCache) loadFrom(src []authx.Authorizer) {

	inst.banks = make(map[dxo.AuthAction]*innerAuthorizerBank)

	for _, item1 := range src {
		info := item1.GetRegistration()
		action := info.Action
		bank := inst.openBank(action)
		bank.add(info)
	}

	// sort all banks
	banks := inst.banks
	for _, bank := range banks {
		bank.sort()
	}

}

func (inst *innerAuthorizerCache) openBank(action dxo.AuthAction) *innerAuthorizerBank {
	bank := inst.banks[action]
	if bank == nil {
		bank = new(innerAuthorizerBank)
		bank.action = action
		inst.banks[action] = bank
	}
	return bank
}

func (inst *innerAuthorizerCache) queryWithAction(action dxo.AuthAction) []*authx.AuthorizerRegistration {

	mtx := &inst.mutex
	mtx.Lock()
	defer mtx.Unlock()

	bank := inst.banks[action]

	if bank != nil {
		return bank.list
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////
