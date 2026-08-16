package testcom

import (
	"context"

	"github.com/starter-go/units"
)

type ExampleUnit struct {

	//starter:component

}

// ListRegistrations implements [units.Unit].
func (inst *ExampleUnit) ListRegistrations(list []*units.Registration) []*units.Registration {

	ur1 := &units.Registration{
		ID:       "test-example",
		Enabled:  true,
		Priority: 0,
		Do:       inst.run,
	}

	list = append(list, ur1)
	return list
}

func (inst *ExampleUnit) run(c context.Context) error {
	return nil
}

func (inst *ExampleUnit) _impl() units.Unit {
	return inst
}
