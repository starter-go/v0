package isessions

import (
	"github.com/starter-go/v0/rbac-web-app/app/classes/statestores"
)

type SimpleStoreChainHolder struct {

	//starter:component

	_as func(statestores.FilterChainHolder) //starter:as("#")

	Loader statestores.FilterChainLoader //starter:inject("#")

	chain statestores.FilterChain
}

// GetChain implements sessions.StoreChainHolder.
func (inst *SimpleStoreChainHolder) GetChain() (statestores.FilterChain, error) {

	ch := inst.chain
	if ch != nil {
		return ch, nil
	}

	ch, err := inst.Loader.LoadChain()
	if err != nil {
		return nil, err
	}

	inst.chain = ch
	return ch, nil
}

func (inst *SimpleStoreChainHolder) _impl() statestores.FilterChainHolder {
	return inst
}
