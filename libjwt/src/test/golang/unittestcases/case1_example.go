package unittestcases

import (
	"github.com/starter-go/units"
	"github.com/starter-go/v0/libjwt/api/jwt"
)

type ExampleCase struct {

	//starter:component

	Service jwt.Service //starter:inject("#")

}

// ListRegistrations implements units.Unit.
func (inst *ExampleCase) ListRegistrations(list []*units.Registration) []*units.Registration {

	const id = "example"

	ur1 := &units.Registration{

		ID:          id,
		Class:       "github.com/starter-go/v0/libjwt/src/test/golang/unittestcases#ExampleCase",
		Name:        id,
		Description: "",

		Enabled:  true,
		Priority: 0,
		OnError:  units.OnErrorError,

		Do: inst.run,
	}

	list = append(list, ur1)
	return list
}

func (inst *ExampleCase) run() error {
	return nil
}

func (inst *ExampleCase) _impl() units.Unit {
	return inst
}
