package subjects

import (
	"context"
	"fmt"
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

	return wtr.Write(ctx2)
}

// Load implements Subject.
func (inst *innerSubjectFacade) Load() error {
	_, err := inst.innerGetCache()
	return err
}

// // SetProperties implements Subject.
// func (inst *innerSubjectFacade) SetProperties(pt properties.Table) error {

// 	if pt == nil {
// 		return nil
// 	}

// 	ctx := inst.context
// 	buffer := ctx.GetBuffer(true)
// 	older := buffer.Properties

// 	if older == nil {
// 		older = properties.NewTable(nil)
// 		buffer.Properties = older
// 	}

// 	older.Import(pt.Export(nil))
// 	return nil
// }

// GetContext implements Subject.
func (inst *innerSubjectFacade) GetContext() context.Context {
	return inst.context.CC
}

// // GetSession implements Subject.
// func (inst *innerSubjectFacade) GetSession() (*rbac.SessionDTO, error) {

// 	cache, err := inst.innerGetCache()
// 	if err != nil {
// 		return nil, err
// 	}

// 	ses := cache.SessionDTO
// 	if ses == nil {
// 		// return nil, fmt.Errorf("subject: Session is nil")
// 		ses = new(rbac.SessionDTO)
// 	}

// 	return ses, nil
// }

// Create implements Subject.
func (inst *innerSubjectFacade) Create() error {

	cache, err := inst.innerGetCache()
	if err == nil && cache != nil {
		sid := cache.SessionID
		suuid := cache.SessionUUID
		if sid > 0 && (suuid != "") {
			return fmt.Errorf("the session exists")
		}
	}

	ctx1 := inst.context
	wtr := ctx1.GetWriter(true)
	ctx2 := ctx1.NewIOC(MethodPost)

	err = wtr.Write(ctx2)
	if err == nil {
		ctx1.Buffer = nil
		ctx1.Cache = nil
	}

	return err
}

// Reload implements Subject.
func (inst *innerSubjectFacade) Reload() error {
	ctx := inst.context
	ctx.Cache = nil
	_, err := inst.innerGetCache()
	return err
}

// Save implements Subject.
func (inst *innerSubjectFacade) Save() error {

	ctx1 := inst.context
	ctx2 := ctx1.NewIOC(MethodFlush)
	wtr := ctx1.Writer
	buffer := ctx1.Buffer
	want := &ctx2.Want

	if buffer != nil {
		want.Properties = buffer.Properties
	}

	err := wtr.Write(ctx2)
	if err == nil {
		ctx1.Buffer = nil
		ctx1.Cache = nil
	}

	return err
}

// // GetProperties implements Subject.
// func (inst *innerSubjectFacade) GetProperties() (properties.Table, error) {

// 	cache, err := inst.innerGetCache()
// 	if err != nil {
// 		return nil, err
// 	}
// 	pt := cache.Properties
// 	if pt == nil {
// 		pt = properties.NewTable(nil)
// 		cache.Properties = pt
// 	}
// 	return pt, nil
// }

func (inst *innerSubjectFacade) innerGetCache() (*Cache, error) {
	ctx := inst.context
	ch := ctx.GetCacheHolder(true)
	return ch.GetCache(ctx)
}

// func (inst *innerSubjectFacade) innerLoadCache() (*Cache, error) {

// 	ctx1 := inst.context
// 	ctx2 := ctx1.NewIOC(MethodGet)
// 	ctx2.Context.Cache = nil

// 	err := inst.context.Reader.Read(ctx2)
// 	if err != nil {
// 		return nil, err
// 	}

// 	cache := ctx2.Context.Cache
// 	if cache == nil {
// 		return nil, fmt.Errorf("innerLoadCache() : cache is nil")
// 	}

// 	return cache, err
// }

func (inst *innerSubjectFacade) _impl() Subject {
	return inst
}

////////////////////////////////////////////////////////////////////////////////
