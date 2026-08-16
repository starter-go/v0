package testcom

import (
	"context"

	"github.com/starter-go/units"
)

type Example4t struct {

	//starter:component
}

// ListRegistrations implements units.Unit.
func (inst *Example4t) ListRegistrations(list []*units.Registration) []*units.Registration {

	u1 := &units.Registration{
		Name:    "example4t",
		Enabled: true,
		Do:      inst.run,
	}

	list = append(list, u1)
	return list
}

func (inst *Example4t) run(cc context.Context) error {
	return nil
}

func (inst *Example4t) _impl() units.Unit {
	return inst
}
