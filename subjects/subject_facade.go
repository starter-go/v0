package subjects

import (
	"context"
)

func NewSubject(ctx *Context) Subject {

	facade := &innerSubjectFacade{
		context: ctx,
	}
	ctx.Facade = facade
	return facade
}

////////////////////////////////////////////////////////////////////////////////

type innerSubjectFacade struct {
	context *Context
}

// Update implements Subject.
func (inst *innerSubjectFacade) Update() error {
	inst.context.FlagDirty = true
	return nil
}

// DoGet implements Subject.
func (inst *innerSubjectFacade) DoGet() (Getter, error) {

	cache, err := inst.innerGetCache()
	if err != nil {
		return nil, err
	}

	gett := new(innerGetterImpl)
	gett.init(inst.context, cache)
	return gett, nil
}

// DoSet implements Subject.
func (inst *innerSubjectFacade) DoSet() (Setter, error) {

	sett := new(innerSetterImpl)
	sett.init(inst.context)
	return sett, nil
}

// Exit implements Subject.
func (inst *innerSubjectFacade) Exit() error {

	ctx1 := inst.context
	ctx2 := ctx1.NewIOC(MethodDelete)
	wtr := ctx1.GetWriter(true)

	defer func() {
		ctx1.FlagCreate = false
		ctx1.FlagDirty = false
		ctx1.Buffer = nil
		ctx1.Cache = nil
	}()

	return wtr.Write(ctx2)
}

// Load implements Subject.
func (inst *innerSubjectFacade) Load() error {
	_, err := inst.innerGetCache()
	return err
}

// GetContext implements Subject.
func (inst *innerSubjectFacade) GetContext() context.Context {
	return inst.context.CC
}

// Create implements Subject.
func (inst *innerSubjectFacade) Create() error {
	ctx := inst.context
	ctx.FlagCreate = true
	ctx.FlagDirty = true
	return nil
}

// Reload implements Subject.
func (inst *innerSubjectFacade) Reload() error {
	ctx := inst.context
	ctx.Cache = nil
	_, err := inst.innerGetCache()
	return err
}

// Save implements Subject.
func (inst *innerSubjectFacade) Flush() error {

	ctx1 := inst.context

	defer func() {
		ctx1.FlagCreate = false
		ctx1.FlagDirty = false
		ctx1.Buffer = nil
		ctx1.Cache = nil
	}()

	// check flag

	dirty := ctx1.FlagDirty
	create := ctx1.FlagCreate
	if !dirty {
		return nil
	}
	method := MethodPut
	if create {
		method = MethodPost
	}

	// do write

	ctx2 := ctx1.NewIOC(method)
	wtr := ctx1.GetWriter(true)
	buffer := ctx1.Buffer
	want := &ctx2.Want

	if buffer != nil {
		want.Properties = buffer.Properties
		want.Buffer = buffer
	}

	return wtr.Write(ctx2)
}

func (inst *innerSubjectFacade) innerGetCache() (*Cache, error) {
	ctx := inst.context
	ch := ctx.GetCacheHolder(true)
	return ch.GetCache(ctx)
}

func (inst *innerSubjectFacade) _impl() Subject {
	return inst
}

////////////////////////////////////////////////////////////////////////////////
