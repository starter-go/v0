package unittestcases

import (
	"net/http"
	"time"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/units"
	"github.com/starter-go/v0/htttest"
)

type CaseTryPostAuth struct {

	//starter:component

	Service htttest.Service //starter:inject("#")

}

// ListRegistrations implements units.Unit.
func (inst *CaseTryPostAuth) ListRegistrations(list []*units.Registration) []*units.Registration {

	const id = "case-try-post-auth"

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

func (inst *CaseTryPostAuth) run() error {

	ser := inst.Service
	agent := ser.GetUserAgent()
	tr := ser.NewTransaction(http.MethodPost, "/api/v1/auth")

	tr.Want.Params = map[string]string{
		"api.version": "v11",
	}
	tr.Want.Query = map[string]string{
		"service": "demo-auth-rest-api",
	}
	tr.Want.Head = map[string]string{
		"x-foo-bar": "12345678-x",
	}

	data1 := new(InnerAuthVO)
	data2 := new(InnerAuthVO)
	tr.Want.Body = &htttest.Content{
		Type: "application/json",
		Data: data1,
	}
	tr.Have.Body = &htttest.Content{
		Data: data2,
	}

	return agent.Execute(tr)
}

func (inst *CaseTryPostAuth) _impl() units.Unit {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

type InnerAuthVO struct {
	Status    int
	Message   string
	Error     string
	Time      time.Time
	Timestamp lang.Time
}

////////////////////////////////////////////////////////////////////////////////
