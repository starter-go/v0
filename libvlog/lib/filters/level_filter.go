package filters

import (
	"net/http"

	"github.com/starter-go/v0/libvlog/api/config"
	"github.com/starter-go/vlog"
)

type LevelFilter struct {

	//starter:component

	_as func(vlog.FilterRegistry) //starter:as(".")

	ConfigService config.Service //starter:inject("#")

	cache *config.Configuration
}

// DoFilter implements vlog.Filter.
func (inst *LevelFilter) DoFilter(msg *vlog.Message, chain vlog.FilterChain) {

	lv := msg.Level
	limit := inst.innerGetLimit()
	met := msg.Method

	if met == http.MethodGet {
		msg.Level = limit
		msg.Status = http.StatusOK
		return
	}

	if lv < limit {
		return
	}

	chain.DoFilter(msg)
}

// ListLogFilterRegistration implements vlog.FilterRegistry.
func (inst *LevelFilter) ListLogFilterRegistration() []*vlog.FilterRegistration {
	mfr := &vlog.FilterRegistration{
		Order:   vlog.OrderLevel,
		Name:    "LevelFilter",
		Group:   vlog.GroupMain,
		Enabled: true,
		Filter:  inst,
	}
	return []*vlog.FilterRegistration{mfr}
}

func (inst *LevelFilter) innerGetLimit() vlog.Level {
	c := inst.cache
	if c == nil {
		ser := inst.ConfigService
		c2, _ := ser.GetConfiguration()
		if c2 == nil {
			c2 = new(config.Configuration)
			ser.LoadDefault(c2)
		}
		c = c2
		inst.cache = c2
	}
	return c.Level
}

func (inst *LevelFilter) _impl() (vlog.FilterRegistry, vlog.Filter) {
	return inst, inst
}

////////////////////////////////////////////////////////////////////////////////
// EOF
