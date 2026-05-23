package filter4log

import (
	"github.com/starter-go/v0/subjects"
	"github.com/starter-go/vlog"
)

type Filter4Log struct {

	//starter:component

	_as func(subjects.FilterRegistry) //starter:as(".")

}

// Read implements subjects.ReadFilter.
func (inst *Filter4Log) Read(c *subjects.IOC, next subjects.ReadFilterChain) error {

	const tag = "Filter4Log-read-"
	inst.log(c, tag+"begin")
	defer func() {
		inst.log(c, tag+"end")
	}()

	return next.Read(c)
}

// Write implements subjects.WriteFilter.
func (inst *Filter4Log) Write(c *subjects.IOC, next subjects.WriteFilterChain) error {

	const tag = "Filter4Log-write-"
	inst.log(c, tag+"begin")
	defer func() {
		inst.log(c, tag+"end")
	}()

	return next.Write(c)
}

func (inst *Filter4Log) log(c *subjects.IOC, tag string) error {

	want := &c.Want

	vlog.Debug("[%s method:%s se_id:%d se_uuid:%s ]", tag, want.Method, want.SessionID, want.SessionUUID)

	return nil
}

// GetRegistrationList implements subjects.FilterRegistry.
func (inst *Filter4Log) GetRegistrationList() []*subjects.FilterRegistration {

	r1 := &subjects.FilterRegistration{
		Name:     "Filter4Log",
		Enabled:  true,
		Priority: subjects.FilterPriorityLog,
		Writer:   inst,
		Reader:   inst,
	}

	return []*subjects.FilterRegistration{r1}
}

func (inst *Filter4Log) _impl() (subjects.FilterRegistry, subjects.WriteFilter, subjects.ReadFilter) {
	return inst, inst, inst
}
