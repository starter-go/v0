package filter4cache

import (
	"net/http"

	"github.com/starter-go/application/properties"
	"github.com/starter-go/v0/subjects"
)

// Filter4CacheEmpty:
// the layer-0-cache filter,
// 使用 empty-object 实现的缓存
type Filter4CacheEmpty struct {

	//starter:component

	_as func(subjects.FilterRegistry) //starter:as(".")

}

// Read implements subjects.ReadFilter.
func (inst *Filter4CacheEmpty) Read(c *subjects.IOC, next subjects.ReadFilterChain) error {

	err := next.Read(c)
	if err != nil {
		inst.innerMakeEmptyFetchResult(c)
	}
	return nil
}

// Write implements subjects.WriteFilter.
func (inst *Filter4CacheEmpty) Write(c *subjects.IOC, next subjects.WriteFilterChain) error {
	return next.Write(c)
}

func (inst *Filter4CacheEmpty) innerMakeEmptyFetchResult(c *subjects.IOC) error {

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

	have.Properties = pt
	have.Status = http.StatusOK

	ctx.Cache = cache

	return nil
}

// GetRegistrationList implements subjects.FilterRegistry.
func (inst *Filter4CacheEmpty) GetRegistrationList() []*subjects.FilterRegistration {

	r1 := &subjects.FilterRegistration{
		Name:     "Filter4CacheEmpty",
		Enabled:  true,
		Priority: subjects.FilterPriorityCacheEmpty,
		Writer:   nil,
		Reader:   inst,
	}

	return []*subjects.FilterRegistration{r1}
}

func (inst *Filter4CacheEmpty) _impl() (subjects.FilterRegistry, subjects.WriteFilter, subjects.ReadFilter) {
	return inst, inst, inst
}
