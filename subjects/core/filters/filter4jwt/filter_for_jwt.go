package filter4jwt

import (
	"fmt"

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

	defer inst.onReadDone(c, token)

	return next.Read(c)
}

func (inst *Filter4jwt) onReadDone(c *subjects.IOC, tk *jwt.Token) error {

	if c == nil || tk == nil {
		return fmt.Errorf("subjects:Filter4jwt.onReadDone() : param(s) is nil")
	}

	have := &c.Have
	cache := have.Cache
	if cache == nil {
		return fmt.Errorf("subjects: ioc.have.cache is nil")
	}

	pt := cache.Properties
	if pt == nil {
		return fmt.Errorf("subjects: cache.Properties is nil")
	}

	const (
		keyTokenNotAfter    string = string(subjects.PNameTokenNotAfter)
		keyTokenNotBefore   string = string(subjects.PNameTokenNotBefore)
		keyTokenSessionID   string = string(subjects.PNameTokenReferID)
		keyTokenSessionUUID string = string(subjects.PNameTokenReferUUID)
	)

	valueNotAfter := tk.NotAfter
	valueNotBefore := tk.NotBefore
	valueSessionID := tk.SessionID
	valueSessionUUID := tk.SessionUUID

	sett := pt.Setter()
	sett.SetInt64(keyTokenNotAfter, int64(valueNotAfter))
	sett.SetInt64(keyTokenNotBefore, int64(valueNotBefore))
	sett.SetInt64(keyTokenSessionID, int64(valueSessionID))
	sett.SetString(keyTokenSessionUUID, valueSessionUUID.String())

	return nil
}

func (inst *Filter4jwt) onWritePre(ioc *subjects.IOC) error {

	if ioc == nil {
		return nil
	}

	subCtx := ioc.Context
	want := &ioc.Want

	if subCtx == nil {
		return nil
	}

	cache := subCtx.Cache
	if cache == nil {
		return nil
	}

	want.SessionID = cache.SessionID
	want.SessionUUID = cache.SessionUUID

	return nil
}

// Write implements subjects.WriteFilter.
func (inst *Filter4jwt) Write(c *subjects.IOC, next subjects.WriteFilterChain) error {

	err := inst.onWritePre(c)
	if err != nil {
		return err
	}

	err = next.Write(c)
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
