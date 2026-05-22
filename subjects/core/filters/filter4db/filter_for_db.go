package filter4db

import (
	"net/http"

	"github.com/starter-go/application/properties"
	"github.com/starter-go/rbac"
	"github.com/starter-go/v0/subjects"
	"github.com/starter-go/v0/subjects/core/classes/sessions"
)

type Filter4db struct {

	//starter:component

	_as func(subjects.FilterRegistry) //starter:as(".")

	Service sessions.Service //starter:inject("#")
}

// Read implements subjects.ReadFilter.
func (inst *Filter4db) Read(c *subjects.IOC, next subjects.ReadFilterChain) error {
	method := c.Want.Method
	switch method {
	case subjects.MethodGet:
		return inst.innerDoGet(c, next)
	}
	return next.Read(c)
}

// Write implements subjects.WriteFilter.
func (inst *Filter4db) Write(c *subjects.IOC, next subjects.WriteFilterChain) error {

	method := c.Want.Method
	switch method {

	case subjects.MethodPut:
		return inst.innerDoPut(c, next)

	case subjects.MethodPost:
		return inst.innerDoPost(c, next)

	case subjects.MethodDelete:
		return inst.innerDoDelete(c, next)

	default:
		// NOP
	}
	return next.Write(c)
}

func (inst *Filter4db) innerDoGet(c *subjects.IOC, next subjects.ReadFilterChain) error {

	ser := inst.Service
	ctx := c.CC
	id := inst.innerGetSessionID(c)

	o1, err := ser.Find(ctx, id)
	if err != nil {
		return err
	}

	have, err := inst.innerMakeNewResult(o1)
	if err != nil {
		return err
	}

	c.Have = *have
	return nil
}

func (inst *Filter4db) innerGetSessionID(c *subjects.IOC) rbac.SessionID {
	return c.Want.SessionID
}

func (inst *Filter4db) innerPrepareWriteSessionDTO(c *subjects.IOC) (*rbac.SessionDTO, error) {

	want := &c.Want
	src := want.Properties
	dst := new(rbac.SessionDTO)

	err := sessions.ConvertP2D(src, dst)
	if err != nil {
		return nil, err
	}

	dst.ID = want.SessionID
	dst.UUID = want.SessionUUID

	return dst, nil
}

func (inst *Filter4db) innerMakeNewResult(src *rbac.SessionDTO) (*subjects.Have, error) {

	dst := new(subjects.Have)
	pt := properties.NewTable(nil)

	err := sessions.ConvertD2P(src, pt)
	if err != nil {
		return nil, err
	}

	dst.SessionID = src.ID
	dst.SessionUUID = src.UUID
	dst.Properties = pt
	dst.Status = http.StatusOK

	return dst, nil
}

func (inst *Filter4db) innerDoPost(c *subjects.IOC, next subjects.WriteFilterChain) error {

	ser := inst.Service
	ctx := c.CC

	o1, err := inst.innerPrepareWriteSessionDTO(c)
	if err != nil {
		return err
	}

	o2, err := ser.Insert(ctx, o1)
	if err != nil {
		return err
	}

	have, err := inst.innerMakeNewResult(o2)
	if err != nil {
		return err
	}

	c.Have = *have
	return nil
}

func (inst *Filter4db) innerDoPut(c *subjects.IOC, next subjects.WriteFilterChain) error {

	ser := inst.Service
	ctx := c.CC
	id := inst.innerGetSessionID(c)

	o1, err := inst.innerPrepareWriteSessionDTO(c)
	if err != nil {
		return err
	}

	o2, err := ser.Update(ctx, id, o1)
	if err != nil {
		return err
	}

	have, err := inst.innerMakeNewResult(o2)
	if err != nil {
		return err
	}

	c.Have = *have
	return nil
}

func (inst *Filter4db) innerDoDelete(c *subjects.IOC, next subjects.WriteFilterChain) error {
	ser := inst.Service
	ctx := c.CC
	id := inst.innerGetSessionID(c)
	return ser.Remove(ctx, id)
}

// GetRegistrationList implements subjects.FilterRegistry.
func (inst *Filter4db) GetRegistrationList() []*subjects.FilterRegistration {

	r1 := &subjects.FilterRegistration{
		Name:     "Filter4db",
		Enabled:  true,
		Priority: subjects.FilterPriorityDB,
		Writer:   inst,
		Reader:   inst,
	}

	return []*subjects.FilterRegistration{r1}
}

func (inst *Filter4db) _impl() (subjects.FilterRegistry, subjects.WriteFilter, subjects.ReadFilter) {
	return inst, inst, inst
}
