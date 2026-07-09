package testcom

import (
	"context"

	"github.com/starter-go/units"
)

type Example1 struct {

	//starter:component

	_as func(units.Unit) //starter:as(".")

}

// ListRegistrations implements units.Unit.
func (inst *Example1) ListRegistrations(list []*units.Registration) []*units.Registration {

	ur := &units.Registration{
		Do:      inst.run,
		Name:    "Example1",
		Enabled: true,
	}

	list = append(list, ur)
	return list
}

func (inst *Example1) run(cc context.Context) error {

	return nil
}

func (inst *Example1) _impl() units.Unit {
	return inst
}
