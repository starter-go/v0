package subjects

import (
	"context"
	"fmt"

	"github.com/starter-go/base/context2"
)

const theAdapterBindingName = "github.com/starter-go/v0/subjects#Adapter.binding"

type Adapter interface {
	GetHolder(c context.Context) (*Holder, error)
}

func GetAdapter(c context.Context) (Adapter, error) {

	// do get

	const name = theAdapterBindingName
	value := c.Value(name)
	ada, ok := value.(Adapter)
	if ok && (ada != nil) {
		return ada, nil
	}

	// do load

	vtable, err := context2.GetValues(c)
	if err != nil {
		return nil, err
	}

	adapterNew := new(innerCommonAdapter)
	ada = adapterNew

	adapterNew.innerSetupHolder(vtable)
	vtable.SetValue(name, ada)

	return ada, nil
}

////////////////////////////////////////////////////////////////////////////////

type innerCommonAdapter struct {
}

// GetHolder implements Adapter.
func (inst *innerCommonAdapter) GetHolder(c context.Context) (*Holder, error) {

	key := inst.innerGetHolderKey()
	value := c.Value(key)
	h, ok := value.(*Holder)
	if !ok || (h == nil) {
		return nil, fmt.Errorf("no subject.Holder in current context")
	}
	return h, nil
}

func (inst *innerCommonAdapter) innerSetupHolder(vtable context2.Values) {

	key := inst.innerGetHolderKey()
	h := new(Holder)
	c := new(Context)
	facade := new(innerSubjectFacade)
	cc := vtable.Context()

	facade.context = c
	h.Context = c
	c.CC = cc
	c.Facade = facade

	vtable.SetValue(key, h)
}

func (inst *innerCommonAdapter) innerGetHolderKey() string {
	return "github.com/starter-go/v0/subjects#Holder.binding"
}

func (inst *innerCommonAdapter) _impl() Adapter {
	return inst
}

////////////////////////////////////////////////////////////////////////////////
