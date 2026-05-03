package subjects

import (
	"fmt"

	"github.com/starter-go/application/properties"
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

// Create implements Subject.
func (inst *innerSubjectFacade) Create() error {
	panic("unimplemented")
}

// Reload implements Subject.
func (inst *innerSubjectFacade) Reload() error {
	panic("unimplemented")
}

// Save implements Subject.
func (inst *innerSubjectFacade) Save() error {
	panic("unimplemented")
}

// GetProperties implements Subject.
func (inst *innerSubjectFacade) GetProperties() (properties.Table, error) {

	cache, err := inst.innerGetCache()
	if err != nil {
		return nil, err
	}
	return cache.Properties, nil
}

func (inst *innerSubjectFacade) innerGetCache() (*Cache, error) {

	cache := inst.context.Cache
	if cache != nil {
		if cache.Ready() {
			return cache, nil
		}
	}

	// load
	cache, err := inst.innerLoadCache()
	if err != nil {
		return nil, err
	}

	inst.context.Cache = cache
	return cache, nil
}

func (inst *innerSubjectFacade) innerLoadCache() (*Cache, error) {

	ctx1 := inst.context
	ctx2 := ctx1.Clone()

	ctx2.Method = MethodGet
	err := inst.context.Reader.Read(ctx2)
	cache := ctx2.Cache

	if cache == nil && err == nil {
		err = fmt.Errorf("innerSubjectFacade.innerLoadCache() : cache is nil")
	}

	return cache, err
}

func (inst *innerSubjectFacade) _impl() Subject {
	return inst
}

////////////////////////////////////////////////////////////////////////////////
