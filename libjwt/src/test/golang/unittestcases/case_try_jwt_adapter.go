package unittestcases

import (
	"context"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/units"
	"github.com/starter-go/v0/libjwt/api/jwt"
	"github.com/starter-go/vlog"
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

func (inst *CaseTryAdapter) innerNewToken() *jwt.Token {

	now := lang.Now()
	t := new(jwt.Token)
	const span = 1000 * 300

	t.NotBefore = now - span
	t.NotAfter = now + span
	t.SessionID = 123456
	t.SessionUUID = "f1385427-0572-720f-ccee-82164dbd7f5c"

	t.NotBefore = 0
	t.NotAfter = 0

	return t
}

func (inst *CaseTryAdapter) run() error {

	ctx := context.Background()
	ada := inst.Service.GetAdapter()
	a1 := new(jwt.Access)
	a2 := new(jwt.Access)
	token1 := inst.innerNewToken()

	a1.Context = ctx
	a1.Service = inst.Service
	a1.Token = token1

	a2.Context = ctx
	a2.Service = inst.Service

	err := ada.SetToken(a1)
	if err != nil {
		return err
	}

	err = ada.GetToken(a2)
	if err != nil {
		return err
	}

	vlog.Debug("access[1].token = %v", a1.Token)
	vlog.Debug("access[2].text  = %s", a2.Text)
	vlog.Debug("access[2].token = %v", a2.Token)

	return nil
}

func (inst *CaseTryAdapter) _impl() units.Unit {
	return inst
}
