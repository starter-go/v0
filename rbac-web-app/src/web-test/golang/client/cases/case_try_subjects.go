package cases

import (
	"net/http"

	"github.com/starter-go/application"
	"github.com/starter-go/units"
	"github.com/starter-go/v0/htttest"
)

type CaseTrySubjects struct {

	//starter:component

	_as func(application.Lifecycle) //starter:as(".")

	HTTTestService htttest.Service //starter:inject("#")

}

// ListRegistrations implements units.Unit.
func (inst *CaseTrySubjects) ListRegistrations(list []*units.Registration) []*units.Registration {
	r1 := &units.Registration{
		Name:     "unit-subjects",
		Enabled:  true,
		Priority: 6,
		Do:       inst.run,
	}
	list = append(list, r1)
	return list
}

func (inst *CaseTrySubjects) run() error {

	ser := inst.HTTTestService
	agent := ser.GetUserAgent()
	tran := ser.NewTransaction(http.MethodGet, "/a/b/c")

	return agent.Execute(tran)
}

func (inst *CaseTrySubjects) _impl() units.Unit {
	return inst
}
