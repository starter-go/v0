package helpfilters

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac"
	"github.com/starter-go/v0/subjects"
)

const (
	attrNameSessionID   = "session.id"
	attrNameSessionUUID = "session.uuid"
)

////////////////////////////////////////////////////////////////////////////////

type MockTokenFilter struct {

	//starter:component

	_as func(subjects.FilterRegistry) //starter:as(".")

}

func (inst *MockTokenFilter) innerLoadParams(c2 *subjects.IOC) {

	c1 := c2.Context
	att := c1.GetAttributes(true)

	// cache := c.GetCache(true)

	id := att.GetAttribute(attrNameSessionID)
	uuid := att.GetAttribute(attrNameSessionUUID)

	if id != nil {
		c2.Want.SessionID = id.(rbac.SessionID)
	}

	if uuid != nil {
		u2 := uuid.(lang.UUID)
		c2.Want.SessionUUID = u2
	}

}

func (inst *MockTokenFilter) innerSaveResult(c2 *subjects.IOC) {

	c1 := c2.Context
	att := c1.GetAttributes(true)

	uuid := c2.Have.SessionUUID
	id := c2.Have.SessionID

	att.SetAttribute(attrNameSessionID, id)
	att.SetAttribute(attrNameSessionUUID, uuid)
}

// Read implements subjects.ReadFilter.
func (inst *MockTokenFilter) Read(c *subjects.IOC, next subjects.ReadFilterChain) error {

	inst.innerLoadParams(c)

	return next.Read(c)
}

// Write implements subjects.WriteFilter.
func (inst *MockTokenFilter) Write(c *subjects.IOC, next subjects.WriteFilterChain) error {

	inst.innerLoadParams(c)

	err := next.Write(c)
	if err != nil {
		return err
	}

	inst.innerSaveResult(c)

	return nil
}

// GetRegistrationList implements subjects.FilterRegistry.
func (inst *MockTokenFilter) GetRegistrationList() []*subjects.FilterRegistration {

	r1 := &subjects.FilterRegistration{
		Name:     "MockTokenFilter",
		Enabled:  true,
		Priority: subjects.FilterPriorityJWT,
		Writer:   inst,
		Reader:   inst,
	}

	return []*subjects.FilterRegistration{r1}
}

func (inst *MockTokenFilter) _impl() (subjects.FilterRegistry, subjects.WriteFilter, subjects.ReadFilter) {
	return inst, inst, inst
}

////////////////////////////////////////////////////////////////////////////////
