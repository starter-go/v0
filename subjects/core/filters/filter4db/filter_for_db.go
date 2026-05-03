package filter4db

import (
	"github.com/starter-go/v0/subjects"
	"github.com/starter-go/v0/subjects/core/classes/sessions"
)

type Filter4db struct {

	//starter:component

	_as func(subjects.FilterRegistry) //starter:as(".")

	Service sessions.Service //starter:inject("#")
}

// Read implements subjects.ReadFilter.
func (inst *Filter4db) Read(c *subjects.Context, next subjects.ReadFilterChain) error {
	return next.Read(c)
}

// Write implements subjects.WriteFilter.
func (inst *Filter4db) Write(c *subjects.Context, next subjects.WriteFilterChain) error {

	return next.Write(c)
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
