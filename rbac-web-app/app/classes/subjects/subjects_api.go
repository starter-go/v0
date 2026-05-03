package subjects

import (
	"context"
	"fmt"

	"github.com/starter-go/application/properties"
	"github.com/starter-go/v0/rbac-web-app/app/classes/statestores"
	"github.com/starter-go/v0/rbac-web-app/app/classes/webcontexts"
	"github.com/starter-go/v0/rbac-web-app/app/web/dto"
)

////////////////////////////////////////////////////////////////////////////////
// api

const (
	theSubjectAttrKey = "github.com/starter-go/v0/rbac-web-app/app/classes/subjects#Subject"
)

type Subject interface {
	GetToken() (*dto.Token, error)

	GetSession() (*dto.Session, error)

	GetUser() (*dto.User, error)

	GetContext() *Context

	GetProperties() properties.Table

	SetProperty(name, value string)

	Save() error

	Create() error

	Delete() error
}

func Current(cc context.Context) (Subject, error) {

	const key = theSubjectAttrKey

	acc, err := webcontexts.GetAccess(cc)
	if err != nil {
		return nil, err
	}

	value := acc.GetAttr(key)
	if value != nil {
		sub, ok := value.(Subject)
		if ok && (sub != nil) {
			return sub, nil
		}
	}

	ctx := NewContext()
	sub := ctx.Subject
	acc.SetAttr(key, sub)
	return sub, nil
}

////////////////////////////////////////////////////////////////////////////////
// facade

type innerSubjectFacade struct {
	context *Context
}

// Delete implements Subject.
func (inst *innerSubjectFacade) Delete() error {

	cache, err := inst.innerGetCache(false)
	if err != nil {
		return err
	}

	ses := cache.Session
	if ses == nil {
		return fmt.Errorf("no session")
	}

	st := inst.context.Store
	sc := new(statestores.Context)

	sc.CC = inst.context.Parent
	sc.Method = statestores.MethodDelete
	sc.SessionID = ses.ID
	sc.SessionUUID = ses.UUID

	return st.Handle(sc)
}

// Create implements Subject.
func (inst *innerSubjectFacade) Create() error {

	cache, err := inst.innerGetCache(true)
	if err != nil {
		return err
	}

	st := inst.context.Store
	sc := new(statestores.Context)

	sc.CC = inst.context.Parent
	sc.Method = statestores.MethodPost

	sc.Properties = cache.Properties
	sc.Session = cache.Session
	sc.Token = cache.Token
	sc.User = cache.User

	err = st.Handle(sc)
	if err == nil {
		inst.context.Cache = cache
	}
	return err
}

// GetProperties implements Subject.
func (inst *innerSubjectFacade) GetProperties() properties.Table {

	cache, err := inst.innerGetCache(true)
	if err != nil {
		panic(err)
	}
	return cache.Properties
}

// Save implements Subject.
func (inst *innerSubjectFacade) Save() error {

	cache, err := inst.innerGetCache(true)
	if err != nil {
		return err
	}

	if !cache.Dirty {
		return nil
	}

	st := inst.context.Store
	sc := new(statestores.Context)

	sc.CC = inst.context.Parent
	sc.Method = statestores.MethodPut

	sc.Properties = cache.Properties
	sc.Session = cache.Session
	sc.Token = cache.Token
	sc.User = cache.User

	err = st.Handle(sc)
	if err == nil {
		cache.Dirty = false
	}
	return err
}

// SetProperty implements Subject.
func (inst *innerSubjectFacade) SetProperty(name string, value string) {

	t := inst.GetProperties()
	t.SetProperty(name, value)
	inst.context.Cache.Dirty = true
}

func (inst *innerSubjectFacade) _impl() Subject {
	return inst
}

func (inst *innerSubjectFacade) innerMakeNewCache() *Cache {

	cache := new(Cache)

	cache.Properties = properties.NewTable(nil)
	cache.Session = new(dto.Session)
	cache.Token = new(dto.Token)
	cache.User = new(dto.User)

	return cache
}

func (inst *innerSubjectFacade) innerGetCache(safe bool) (*Cache, error) {

	cache := inst.context.Cache
	if cache != nil {
		return cache, nil
	}

	// load
	cache, err := inst.innerLoadCache()
	if err != nil {
		if safe {
			cache = inst.innerMakeNewCache()
			err = nil
		} else {
			return nil, err
		}
	}

	inst.context.Cache = cache
	return cache, err
}

func (inst *innerSubjectFacade) innerLoadCache() (*Cache, error) {

	cache := new(Cache)
	sc := new(statestores.Context)
	st := inst.context.Store

	sc.CC = inst.context.Parent
	sc.Method = statestores.MethodGet

	err := st.Handle(sc)
	if err == nil {
		cache.Dirty = false
		cache.Properties = sc.Properties
		cache.Session = sc.Session
		cache.Token = sc.Token
		cache.User = sc.User
	}

	return cache, err
}

func (inst *innerSubjectFacade) GetToken() (*dto.Token, error) {

	cache, err := inst.innerGetCache(false)
	if err != nil {
		return nil, err
	}
	return cache.Token, nil
}

func (inst *innerSubjectFacade) GetSession() (*dto.Session, error) {

	cache, err := inst.innerGetCache(false)
	if err != nil {
		return nil, err
	}
	return cache.Session, nil
}

func (inst *innerSubjectFacade) GetUser() (*dto.User, error) {

	cache, err := inst.innerGetCache(false)
	if err != nil {
		return nil, err
	}
	return cache.User, nil
}

func (inst *innerSubjectFacade) GetContext() *Context {
	return inst.context
}

////////////////////////////////////////////////////////////////////////////////
// EOF
