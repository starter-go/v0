package cases

import (
	"net/http"

	"github.com/starter-go/units"
	"github.com/starter-go/v0/htttest"
)

type CaseTryExample struct {

	//starter:component

	HTTTestService htttest.Service //starter:inject("#")

}

// ListRegistrations implements units.Unit.
func (inst *CaseTryExample) ListRegistrations(list []*units.Registration) []*units.Registration {
	r1 := &units.Registration{
		Name:    "unit-example-cl",
		Enabled: true,
		Do:      inst.run,
	}
	list = append(list, r1)
	return list
}

func (inst *CaseTryExample) run() error {

	ser := inst.HTTTestService
	agent := ser.GetUserAgent()
	tran := ser.NewTransaction(http.MethodGet, "/a/b/c")

	return agent.Execute(tran)
}

func (inst *CaseTryExample) _impl() units.Unit {
	return inst
}
