package cases

import (
	"fmt"
	"net/http"

	"github.com/starter-go/application/properties"
	"github.com/starter-go/rbac"
	"github.com/starter-go/units"
	"github.com/starter-go/v0/htttest"
)

type CaseTryUseAuth struct {

	//starter:component

	HTTTestService htttest.Service //starter:inject("#")

}

// ListRegistrations implements units.Unit.
func (inst *CaseTryUseAuth) ListRegistrations(list []*units.Registration) []*units.Registration {
	r1 := &units.Registration{
		Name:     "unit-auth",
		Enabled:  true,
		Priority: 99,
		Do:       inst.run,
	}
	list = append(list, r1)
	return list
}

func (inst *CaseTryUseAuth) run() error {

	ser := inst.HTTTestService
	agent := ser.GetUserAgent()
	tran := ser.NewTransaction(http.MethodPost, "/api/v1/auth")

	head1 := make(map[string]string)
	head1["Authorization"] = "Basic dGVzdDoxMjPCow=="
	tran.Want.Head = head1

	vo1 := new(rbac.AuthVO)
	vo2 := new(rbac.AuthVO)
	body1 := &htttest.Content{
		Type: "application/json",
		Data: vo1,
	}
	body2 := &htttest.Content{
		Type: "application/json",
		Data: vo2,
	}
	tran.Want.Body = body1
	tran.Have.Body = body2

	defer func() {
		if tran.Have.Error == nil {
			inst.innerSaveJwtFromCookie(tran)
			inst.innerSaveJwtFromHeader(tran)
		}
	}()

	return agent.Execute(tran)
}

func (inst *CaseTryUseAuth) innerSaveJwtFromCookie(tr *htttest.Transaction) error {
	return fmt.Errorf("no impl")
}

func (inst *CaseTryUseAuth) innerSaveJwtFromHeader(tr *htttest.Transaction) error {
	const (
		name1 = "x-jwt"
		name2 = "x-set-jwt"
	)

	src := tr.Have.Head
	dst := tr.AC.RequestHeaders
	if src == nil {
		return nil
	}
	if dst == nil {
		dst = properties.NewTable(nil)
		tr.AC.RequestHeaders = dst
	}

	value := src[name2]
	if value == "" {
		return nil
	}
	dst.SetProperty(name1, value)
	return nil
}

func (inst *CaseTryUseAuth) _impl() units.Unit {
	return inst
}
