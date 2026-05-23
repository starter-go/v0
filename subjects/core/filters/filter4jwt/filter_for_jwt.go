package filter4jwt

import (
	"github.com/starter-go/rbac"
	"github.com/starter-go/v0/libjwt/api/jwt"
	"github.com/starter-go/v0/subjects"
)

type Filter4jwt struct {

	//starter:component

	_as func(subjects.FilterRegistry) //starter:as(".")

	JWTService jwt.Service //starter:inject("#")

}

// Read implements subjects.ReadFilter.
func (inst *Filter4jwt) Read(c *subjects.IOC, next subjects.ReadFilterChain) error {

	ser := inst.JWTService
	ada := ser.GetAdapter()

	acc := &jwt.Access{
		Context: c.CC,
		Service: ser,
	}

	err := ada.GetToken(acc)
	if err != nil {
		return err
	}

	token := acc.Token
	want := &c.Want

	want.SessionID = rbac.SessionID(token.SessionID)
	want.SessionUUID = token.SessionUUID

	return next.Read(c)
}

// Write implements subjects.WriteFilter.
func (inst *Filter4jwt) Write(c *subjects.IOC, next subjects.WriteFilterChain) error {

	err := next.Write(c)
	if err != nil {
		return err
	}

	have := &c.Have
	sid := have.SessionID
	suuid := have.SessionUUID
	ser := inst.JWTService
	ada := ser.GetAdapter()

	token := &jwt.Token{
		SessionID:   int64(sid),
		SessionUUID: suuid,
	}

	acc := &jwt.Access{
		Context: c.CC,
		Service: ser,
		Token:   token,
	}

	err = ada.SetToken(acc)
	return err
}

// GetRegistrationList implements subjects.FilterRegistry.
func (inst *Filter4jwt) GetRegistrationList() []*subjects.FilterRegistration {

	r1 := &subjects.FilterRegistration{
		Name:     "Filter4jwt",
		Enabled:  true,
		Priority: subjects.FilterPriorityJWT,
		Writer:   inst,
		Reader:   inst,
	}

	return []*subjects.FilterRegistration{r1}
}

func (inst *Filter4jwt) _impl() (subjects.FilterRegistry, subjects.WriteFilter, subjects.ReadFilter) {
	return inst, inst, inst
}
