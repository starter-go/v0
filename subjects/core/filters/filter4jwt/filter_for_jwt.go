package filter4jwt

import "github.com/starter-go/v0/subjects"

type Filter4jwt struct {

	//starter:component

	_as func(subjects.FilterRegistry) //starter:as(".")

}

// Read implements subjects.ReadFilter.
func (inst *Filter4jwt) Read(c *subjects.Context, next subjects.ReadFilterChain) error {

	return next.Read(c)
}

// Write implements subjects.WriteFilter.
func (inst *Filter4jwt) Write(c *subjects.Context, next subjects.WriteFilterChain) error {

	return next.Write(c)
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
