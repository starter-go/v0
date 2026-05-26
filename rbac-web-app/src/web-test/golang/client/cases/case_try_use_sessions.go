package cases

import (
	"net/http"

	"github.com/starter-go/rbac"
	"github.com/starter-go/units"
	"github.com/starter-go/v0/htttest"
	"github.com/starter-go/vlog"
)

type CaseTryUseSessions struct {

	//starter:component

	HTTTestService htttest.Service //starter:inject("#")

}

// ListRegistrations implements units.Unit.
func (inst *CaseTryUseSessions) ListRegistrations(list []*units.Registration) []*units.Registration {
	r1 := &units.Registration{
		Name:     "unit-sessions",
		Enabled:  true,
		Priority: 5,
		Do:       inst.run,
	}
	list = append(list, r1)
	return list
}

func (inst *CaseTryUseSessions) run() error {

	ser := inst.HTTTestService
	agent := ser.GetUserAgent()
	tran := ser.NewTransaction(http.MethodGet, "/api/v1/sessions/current")

	vo2 := new(rbac.SessionVO)
	body2 := &htttest.Content{
		Data: vo2,
	}
	tran.Have.Body = body2

	defer func() {
		vlog.Debug("resp.body = %v", body2)
		vlog.Debug("resp.vo   = %v", vo2)
		vlog.Debug("tran      = %v", tran)
	}()

	return agent.Execute(tran)
}

func (inst *CaseTryUseSessions) _impl() units.Unit {
	return inst
}
