package unitcom

import (
	"context"

	"github.com/starter-go/units"
)

type Example1 struct {

	//starter:component
}

// ListRegistrations implements units.Unit.
func (inst *Example1) ListRegistrations(list []*units.Registration) []*units.Registration {

	u1 := &units.Registration{
		Name:    "example1",
		Enabled: true,
		Do:      inst.run,
	}

	list = append(list, u1)
	return list
}

func (inst *Example1) run(cc context.Context) error {
	return nil
}

func (inst *Example1) _impl() units.Unit {
	return inst
}
