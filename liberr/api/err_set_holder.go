package api

import (
	"fmt"
	"sync"

	"github.com/starter-go/base/lang"
)

var theCoreErrorSet LibCoreErrorMaker

type ErrorSetHolder struct {
	loader ErrorSetLoader
	// es     ErrorSet
	mu    sync.Mutex
	table map[lang.URI]*Registration
}

func (inst *ErrorSetHolder) innerGetTable() (map[lang.URI]*Registration, error) {
	tab := inst.table
	if tab == nil {
		// do load
		t2, err := inst.innerLoadTable()
		if err != nil {
			return nil, err
		}
		tab = t2
		inst.table = tab
	}
	return tab, nil
}

func (inst *ErrorSetHolder) innerLoadTable() (map[lang.URI]*Registration, error) {

	ldr := inst.loader
	if ldr == nil {
		return nil, fmt.Errorf("liberr:ErrorSetHolder: ErrorSet loader is nil")
	}

	es := ldr.Load()
	src := es.ListErrors()
	dst := make(map[lang.URI]*Registration)

	for _, it := range src {
		uri := it.URI
		dst[uri] = it
	}

	return dst, nil
}

func (inst *ErrorSetHolder) GetRegistration(uri lang.URI) (*Registration, error) {

	m := &inst.mu
	m.Lock()
	defer m.Unlock()

	tab, err := inst.innerGetTable()
	if err != nil {
		return nil, err
	}

	item := tab[uri]
	if item == nil {
		return nil, fmt.Errorf("liberr:ErrorSetHolder: no HyperError with URI: '%s'", uri)
	}

	return item, nil
}

func (inst *ErrorSetHolder) From(l ErrorSetLoader) *ErrorSetHolder {
	inst.loader = l
	return inst
}

func (inst *ErrorSetHolder) ErrorNS(ns Namespace, name Name, args ...any) error {
	uri := ComputeErrorURI(ns, name)
	return inst.Error(uri, args...)
}

func (inst *ErrorSetHolder) Error(uri lang.URI, args ...any) error {

	reg, err := inst.GetRegistration(uri)
	if err != nil {
		return err
	}

	ff := reg.GetFormatter()
	if ff == nil {
		ff = DefaultFormatter()
	}

	return ff.Format(args...)
}
