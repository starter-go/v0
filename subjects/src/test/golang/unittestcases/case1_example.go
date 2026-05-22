package unittestcases

import (
	"github.com/starter-go/units"
)

type ExampleCase struct {

	//starter:component

	Enabled bool //starter:inject("${unit.test-example.enabled}")

}

// ListRegistrations implements units.Unit.
func (inst *ExampleCase) ListRegistrations(list []*units.Registration) []*units.Registration {

	u1 := &units.Registration{
		Name:    "ExampleCase",
		Enabled: inst.Enabled,
		Do:      inst.run,
	}

	list = append(list, u1)
	return list
}

func (inst *ExampleCase) _impl() units.Unit {
	return inst
}

func (inst *ExampleCase) run() error {
	return nil
}
