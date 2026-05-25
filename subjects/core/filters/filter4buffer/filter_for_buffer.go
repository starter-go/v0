package filter4buffer

import "github.com/starter-go/v0/subjects"

type Filter4Buffer struct {

	//starter:component

	_as func(subjects.FilterRegistry) //starter:as(".")

}

// Read implements subjects.ReadFilter.
func (inst *Filter4Buffer) Read(c *subjects.IOC, next subjects.ReadFilterChain) error {

	return next.Read(c)
}

// Write implements subjects.WriteFilter.
func (inst *Filter4Buffer) Write(c *subjects.IOC, next subjects.WriteFilterChain) error {
	method := c.Want.Method
	switch method {
	case subjects.MethodFlush:
		return inst.innerDoFlush(c, next)
	case subjects.MethodPut:
		return inst.innerDoPostOrPut(c, next, false)
	case subjects.MethodPost:
		return inst.innerDoPostOrPut(c, next, true)
	}
	return next.Write(c)
}

func (inst *Filter4Buffer) innerDoFlush(c *subjects.IOC, next subjects.WriteFilterChain) error {

	ctx := c.Context
	buffer := ctx.Buffer

	if buffer.IsEmpty() {
		return nil
	}

	want := &c.Want
	want.Properties = buffer.Properties

	// rewrite-method
	if ctx.DoCreate {
		want.Method = subjects.MethodPost
	} else {
		want.Method = subjects.MethodPut
	}

	ctx.Buffer = nil
	ctx.DoCreate = false

	return next.Write(c)
}

func (inst *Filter4Buffer) innerDoPostOrPut(c *subjects.IOC, next subjects.WriteFilterChain, do_create bool) error {

	ctx := c.Context
	buffer := ctx.Buffer
	buffer2 := buffer.Init()
	if buffer == nil {
		buffer = buffer2
		c.Context.Buffer = buffer2
	}

	src := c.Want.Properties

	if do_create {
		ctx.DoCreate = true
	}

	if src != nil {
		dst := buffer.Properties
		tmp := src.Export(nil)
		dst.Import(tmp)
	}

	return nil
}

// GetRegistrationList implements subjects.FilterRegistry.
func (inst *Filter4Buffer) GetRegistrationList() []*subjects.FilterRegistration {

	r1 := &subjects.FilterRegistration{
		Name:     "Filter4Buffer",
		Enabled:  false,
		Priority: subjects.FilterPriorityBuffer,
		Writer:   inst,
		Reader:   nil,
	}

	return []*subjects.FilterRegistration{r1}
}

func (inst *Filter4Buffer) _impl() (subjects.FilterRegistry, subjects.WriteFilter, subjects.ReadFilter) {
	return inst, inst, inst
}
