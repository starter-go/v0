package testcase

import (
	"github.com/starter-go/units"
)

type CaseTryExample struct {

	//starter:component

}

// ListRegistrations implements units.Unit.
func (inst *CaseTryExample) ListRegistrations(list []*units.Registration) []*units.Registration {

	r1 := &units.Registration{
		Name:    "unit-example-t",
		Enabled: true,
		Do:      inst.run,
	}

	list = append(list, r1)
	return list
}

func (inst *CaseTryExample) run() error {
	return nil
}

func (inst *CaseTryExample) _impl() units.Unit {
	return inst
}
