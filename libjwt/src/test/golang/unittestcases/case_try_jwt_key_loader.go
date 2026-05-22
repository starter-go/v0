package unittestcases

import (
	"github.com/starter-go/units"
	"github.com/starter-go/v0/libjwt/api/jwt"
	"github.com/starter-go/vlog"
)

type CaseTryKeyLoader struct {

	//starter:component

	Loader jwt.KeyLoader //starter:inject("#")

}

// ListRegistrations implements units.Unit.
func (inst *CaseTryKeyLoader) ListRegistrations(list []*units.Registration) []*units.Registration {

	const id = "try-key-loader"

	ur1 := &units.Registration{

		ID:          id,
		Class:       "github.com/starter-go/v0/libjwt/src/test/golang/unittestcases#CaseTryKeyLoader",
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

func (inst *CaseTryKeyLoader) run() error {

	k, err := inst.Loader.Load()
	if err != nil {
		return err
	}

	bin := k.Bytes()
	vlog.Debug(" load key raw-data, bytes = %v", bin)

	return nil
}

func (inst *CaseTryKeyLoader) _impl() units.Unit {
	return inst
}
