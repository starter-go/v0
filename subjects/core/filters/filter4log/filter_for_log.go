package filter4log

import "github.com/starter-go/v0/subjects"

type Filter4Log struct {

	//starter:component

	_as func(subjects.FilterRegistry) //starter:as(".")

}

// Read implements subjects.ReadFilter.
func (inst *Filter4Log) Read(c *subjects.Context, next subjects.ReadFilterChain) error {

	return next.Read(c)
}

// Write implements subjects.WriteFilter.
func (inst *Filter4Log) Write(c *subjects.Context, next subjects.WriteFilterChain) error {

	return next.Write(c)
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
