package istore

import (
	"context"

	"github.com/starter-go/v0/rbac-web-app/app/classes/statestores"
)

type StateStoreServiceImpl struct {

	//starter:component

	_as func(statestores.Service) //starter:as("#")

	Holder statestores.FilterChainHolder //starter:inject("#")

}

// GetStore implements statestores.Service.
func (inst *StateStoreServiceImpl) GetStore(cc context.Context) (statestores.Store, error) {

	ch, err := inst.Holder.GetChain()

	return ch, err
}

func (inst *StateStoreServiceImpl) _impl() statestores.Service {
	return inst
}
