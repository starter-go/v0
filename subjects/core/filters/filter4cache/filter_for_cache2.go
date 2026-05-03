package filter4cache

import "github.com/starter-go/v0/subjects"

// Filter4Cache2:
// the layer-2-cache filter,
// 使用 Redis 实现的缓存
type Filter4Cache2 struct {

	//starter:component

	_as func(subjects.FilterRegistry) //starter:as(".")

}

// Read implements subjects.ReadFilter.
func (inst *Filter4Cache2) Read(c *subjects.Context, next subjects.ReadFilterChain) error {

	return next.Read(c)
}

// Write implements subjects.WriteFilter.
func (inst *Filter4Cache2) Write(c *subjects.Context, next subjects.WriteFilterChain) error {

	return next.Write(c)
}

// GetRegistrationList implements subjects.FilterRegistry.
func (inst *Filter4Cache2) GetRegistrationList() []*subjects.FilterRegistration {

	r1 := &subjects.FilterRegistration{
		Name:     "Filter4Cache2",
		Enabled:  true,
		Priority: subjects.FilterPriorityCache2,
		Writer:   inst,
		Reader:   inst,
	}

	return []*subjects.FilterRegistration{r1}
}

func (inst *Filter4Cache2) _impl() (subjects.FilterRegistry, subjects.WriteFilter, subjects.ReadFilter) {
	return inst, inst, inst
}
