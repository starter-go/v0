package pack2

import (
	"context"

	"github.com/starter-go/units"
)

type Demo2units struct {

	//starter:component

	////// _as func() //starter:as(".")

}

// ListRegistrations implements [units.Unit].
func (inst *Demo2units) ListRegistrations(list []*units.Registration) []*units.Registration {

	u1 := &units.Registration{
		Name:    "Demo2_unit_1",
		Enabled: true,
		Do:      inst.runUnit1,
	}

	list = append(list, u1)
	return list
}

func (inst *Demo2units) runUnit1(cc context.Context) error {
	return nil
}

func (inst *Demo2units) _impl() units.Unit {
	return inst
}
