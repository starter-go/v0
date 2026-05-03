package isessions

import (
	"github.com/starter-go/v0/rbac-web-app/app/classes/statestores"
)

type CacheSessionStoreFilter struct {

	//starter:component

	_as func(statestores.FilterRegistry) //starter:as(".")

}

// Handle implements sessions.StoreFilter.
func (inst *CacheSessionStoreFilter) Handle(sc *statestores.Context, next statestores.FilterChain) error {

	return next.Handle(sc)
}

// ListRegistrations implements sessions.StoreFilterRegistry.
func (inst *CacheSessionStoreFilter) ListRegistrations() []*statestores.FilterRegistration {

	reg := &statestores.FilterRegistration{
		Name:     "CacheSessionStoreFilter",
		Enabled:  true,
		Priority: statestores.FilterPriorityCache,
		Filter:   inst,
	}

	return []*statestores.FilterRegistration{reg}
}

func (inst *CacheSessionStoreFilter) _impl() (statestores.FilterRegistry, statestores.Filter) {
	return inst, inst
}
