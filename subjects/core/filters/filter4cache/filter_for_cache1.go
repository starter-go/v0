package filter4cache

import "github.com/starter-go/v0/subjects"

// Filter4Cache1:
// the layer-1-cache filter,
// 直接使用内存实现的 cache
type Filter4Cache1 struct {

	//starter:component

	_as func(subjects.FilterRegistry) //starter:as(".")

}

// Read implements subjects.ReadFilter.
func (inst *Filter4Cache1) Read(c *subjects.Context, next subjects.ReadFilterChain) error {
	return next.Read(c)
}

// Write implements subjects.WriteFilter.
func (inst *Filter4Cache1) Write(c *subjects.Context, next subjects.WriteFilterChain) error {

	return next.Write(c)
}

// GetRegistrationList implements subjects.FilterRegistry.
func (inst *Filter4Cache1) GetRegistrationList() []*subjects.FilterRegistration {

	r1 := &subjects.FilterRegistration{
		Name:     "Filter4Cache1",
		Enabled:  true,
		Priority: subjects.FilterPriorityCache1,
		Writer:   inst,
		Reader:   inst,
	}

	return []*subjects.FilterRegistration{r1}
}

func (inst *Filter4Cache1) _impl() (subjects.FilterRegistry, subjects.WriteFilter, subjects.ReadFilter) {
	return inst, inst, inst
}
