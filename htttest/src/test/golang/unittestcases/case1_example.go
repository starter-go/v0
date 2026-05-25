package unittestcases

import (
	"net/http"

	"github.com/starter-go/units"
	"github.com/starter-go/v0/htttest"
)

type ExampleCase struct {

	//starter:component

	Service htttest.Service //starter:inject("#")

}

// ListRegistrations implements units.Unit.
func (inst *ExampleCase) ListRegistrations(list []*units.Registration) []*units.Registration {

	const id = "example"

	r1 := &units.Registration{

		Name:  id,
		ID:    id,
		Class: "",

		Enabled:  true,
		Priority: 0,

		Do: inst.run,
	}

	list = append(list, r1)

	return list
}

func (inst *ExampleCase) run() error {

	ser := inst.Service
	agent := ser.GetUserAgent()
	reqctx := ser.NewTransaction(http.MethodGet, "/a/b/c")

	return agent.Execute(reqctx)
}

func (inst *ExampleCase) _impl() units.Unit {
	return inst
}
