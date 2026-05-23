package filter4cache

import (
	"fmt"
	"net/http"

	"github.com/starter-go/application/properties"
	"github.com/starter-go/v0/subjects"
)

type Filter4GoodResult struct {

	//starter:component

	_as func(subjects.FilterRegistry) //starter:as(".")

}

// Read implements subjects.ReadFilter.
func (inst *Filter4GoodResult) Read(c *subjects.IOC, next subjects.ReadFilterChain) error {

	err := next.Read(c)
	return inst.innerHandleResult(c, err)
}

// Write implements subjects.WriteFilter.
func (inst *Filter4GoodResult) Write(c *subjects.IOC, next subjects.WriteFilterChain) error {

	err := next.Write(c)
	if err != nil {
		return err
	}

	ctx := c.Context
	ctx.Cache = nil
	ctx.Buffer = nil

	return nil
}

func (inst *Filter4GoodResult) innerHandleResult(c *subjects.IOC, e2 error) error {

	if e2 == nil {
		e2 = inst.innerCheckResult(c)
		if e2 == nil {
			return nil
		}
	}

	return inst.innerMakeGoodResult(c)
}

func (inst *Filter4GoodResult) innerCheckResult(c *subjects.IOC) error {

	have := &c.Have
	cache := have.Cache
	if cache == nil {
		return fmt.Errorf("subjects.IOC.result: cache is nil")
	}

	pt1 := have.Properties
	pt2 := cache.Properties

	if pt1 == nil || pt2 == nil {
		return fmt.Errorf("subjects.IOC.result: properties is nil")
	}

	return nil
}

func (inst *Filter4GoodResult) innerMakeGoodResult(c *subjects.IOC) error {

	const uuid = "00000000-0000-0000-0000-000000000000"

	have := &c.Have
	ctx := c.Context
	cache := ctx.Cache
	pt := have.Properties

	if cache == nil {
		cache = new(subjects.Cache)
	}

	if pt == nil {
		pt = properties.NewTable(nil)
	}

	cache.Properties = pt
	cache.Loaded = true
	cache.SessionUUID = uuid

	have.Properties = pt
	have.Status = http.StatusOK
	have.Cache = cache
	have.SessionUUID = uuid

	ctx.Cache = cache

	return nil
}

// GetRegistrationList implements subjects.FilterRegistry.
func (inst *Filter4GoodResult) GetRegistrationList() []*subjects.FilterRegistration {

	r1 := &subjects.FilterRegistration{
		Name:     "Filter4GoodResult",
		Enabled:  true,
		Priority: subjects.FilterPriorityGoodResult,
		Writer:   inst,
		Reader:   inst,
	}

	return []*subjects.FilterRegistration{r1}
}

func (inst *Filter4GoodResult) _impl() (subjects.FilterRegistry, subjects.WriteFilter, subjects.ReadFilter) {
	return inst, inst, inst
}
