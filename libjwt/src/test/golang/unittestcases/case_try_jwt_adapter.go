package unittestcases

import (
	"github.com/starter-go/units"
	"github.com/starter-go/v0/libjwt/api/jwt"
)

type CaseTryAdapter struct {

	//starter:component

	Service jwt.Service //starter:inject("#")

}

// ListRegistrations implements units.Unit.
func (inst *CaseTryAdapter) ListRegistrations(list []*units.Registration) []*units.Registration {

	const id = "try-jwt-adapter"

	ur1 := &units.Registration{

		ID:          id,
		Name:        id,
		Class:       "github.com/starter-go/v0/libjwt/src/test/golang/unittestcases#CaseTryAdapter",
		Description: "",

		Enabled:  true,
		Priority: 0,
		OnError:  units.OnErrorError,

		Do: inst.run,
	}

	list = append(list, ur1)
	return list
}

func (inst *CaseTryAdapter) run() error {
	return nil
}

func (inst *CaseTryAdapter) _impl() units.Unit {
	return inst
}
