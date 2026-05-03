package isessions

import (
	"github.com/starter-go/v0/rbac-web-app/app/classes/statestores"
)

type SimpleStoreChainLoader struct {

	//starter:component

	_as func(statestores.FilterChainLoader) //starter:as("#")

	RegistryList []statestores.FilterRegistry //starter:inject(".")

}

// LoadChain implements sessions.StoreChainLoader.
func (inst *SimpleStoreChainLoader) LoadChain() (statestores.FilterChain, error) {

	builder := new(statestores.FilterChainBuilder)
	src := inst.RegistryList

	for _, it := range src {
		builder.AddRegistry(it)
	}

	ch := builder.Build()
	return ch, nil
}

func (inst *SimpleStoreChainLoader) _impl() statestores.FilterChainLoader {
	return inst
}
