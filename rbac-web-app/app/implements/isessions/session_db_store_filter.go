package isessions

import (
	"github.com/starter-go/application/properties"
	"github.com/starter-go/v0/rbac-web-app/app/classes/sessions"
	"github.com/starter-go/v0/rbac-web-app/app/classes/statestores"
)

type DatabaseSessionStoreFilter struct {

	//starter:component

	_as func(statestores.FilterRegistry) //starter:as(".")

	Sessions sessions.Service //starter:inject("#")
}

// Handle implements sessions.StoreFilter.
func (inst *DatabaseSessionStoreFilter) Handle(sc *statestores.Context, next statestores.FilterChain) error {

	method := sc.Method

	switch method {
	case statestores.MethodGet:
		return inst.handleGet(sc, next)

	case statestores.MethodPut:
		return inst.handlePut(sc, next)

	case statestores.MethodPost:
		return inst.handlePost(sc, next)

	case statestores.MethodDelete:
		return inst.handleDel(sc, next)
	}

	return next.Handle(sc)
}

func (inst *DatabaseSessionStoreFilter) handleGet(sc *statestores.Context, next statestores.FilterChain) error {

	ctx := sc.CC
	uuid := sc.SessionUUID
	props := sc.Properties

	if props == nil {
		props = properties.NewTable(nil)
		sc.Properties = props
	}

	item, err := inst.Sessions.FindByUUID(ctx, uuid)
	if err != nil {
		return err
	}

	props.Import(item.Properties)

	sc.Session = item
	sc.Properties = props

	return next.Handle(sc)
}

func (inst *DatabaseSessionStoreFilter) handlePost(sc *statestores.Context, next statestores.FilterChain) error {

	ctx := sc.CC
	ser := inst.Sessions
	o1 := sc.Session

	o2, err := ser.Insert(ctx, o1)
	if err != nil {
		return err
	}

	sc.Session = o2
	sc.SessionID = o2.ID
	sc.SessionUUID = o2.UUID

	return next.Handle(sc)
}

func (inst *DatabaseSessionStoreFilter) handlePut(sc *statestores.Context, next statestores.FilterChain) error {

	ctx := sc.CC
	ser := inst.Sessions
	o1 := sc.Session
	id := sc.SessionID

	o2, err := ser.Update(ctx, id, o1)
	if err != nil {
		return err
	}

	sc.Session = o2

	return next.Handle(sc)
}

func (inst *DatabaseSessionStoreFilter) handleDel(sc *statestores.Context, next statestores.FilterChain) error {

	ctx := sc.CC
	ser := inst.Sessions
	id := sc.SessionID

	err := ser.Remove(ctx, id)
	if err != nil {
		return err
	}

	sc.Session = new(sessions.DTO)

	return next.Handle(sc)
}

// ListRegistrations implements sessions.StoreFilterRegistry.
func (inst *DatabaseSessionStoreFilter) ListRegistrations() []*statestores.FilterRegistration {

	reg := &statestores.FilterRegistration{
		Name:     "DatabaseSessionStoreFilter",
		Enabled:  true,
		Priority: statestores.FilterPriorityDB,
		Filter:   inst,
	}

	return []*statestores.FilterRegistration{reg}
}

func (inst *DatabaseSessionStoreFilter) _impl() (statestores.FilterRegistry, statestores.Filter) {
	return inst, inst
}
