package filter4cache

import (
	"fmt"
	"net/http"

	"github.com/starter-go/v0/subjects"
)

// Filter4Cache1:
// the layer-1-cache filter,
// 直接使用内存实现的 cache
type Filter4Cache1 struct {

	//starter:component

	_as func(subjects.FilterRegistry) //starter:as(".")

}

// Read implements subjects.ReadFilter.
func (inst *Filter4Cache1) Read(c *subjects.IOC, next subjects.ReadFilterChain) error {

	// 先检查 context.cache
	ctx := c.Context
	cache := ctx.Cache
	if cache != nil {
		if cache.Ready() {
			return inst.innerMakeResultFromCache(c, cache)
		}
	}

	err := next.Read(c)
	if err != nil {
		return err
	}

	// read 成功后, 保存新的 cache
	return inst.innerKeepNewCache(c)
}

func (inst *Filter4Cache1) innerMakeResultFromCache(c *subjects.IOC, cache *subjects.Cache) error {

	have := &c.Have

	have.SessionID = cache.SessionID
	have.SessionUUID = cache.SessionUUID
	have.Properties = cache.Properties
	have.Status = http.StatusOK
	have.Cache = cache

	return nil
}

func (inst *Filter4Cache1) innerKeepNewCache(c *subjects.IOC) error {

	ctx := c.Context
	cache := c.Have.Cache
	ctx.Cache = cache

	if cache == nil {
		return fmt.Errorf("Filter4Cache1: result cache is nil")
	}

	return nil
}

// Write implements subjects.WriteFilter.
func (inst *Filter4Cache1) Write(c *subjects.IOC, next subjects.WriteFilterChain) error {

	err := next.Write(c)
	if err != nil {
		return err
	}

	// 写成功后, 清除  cache
	ctx := c.Context
	ctx.Cache = nil

	return nil
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
