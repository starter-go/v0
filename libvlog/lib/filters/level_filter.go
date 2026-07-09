package filters

import (
	"net/http"

	"github.com/starter-go/vlog"
)

type LevelFilter struct {

	//starter:component

	_as func(vlog.FilterRegistry) //starter:as(".")

	Level string //starter:inject("${vlog.level}")

	cache *innerLevelFilterCache
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
		c = new(innerLevelFilterCache)
		c.load(inst)
		inst.cache = c
	}
	return c.limit
}

func (inst *LevelFilter) _impl() (vlog.FilterRegistry, vlog.Filter) {
	return inst, inst
}

////////////////////////////////////////////////////////////////////////////////

type innerLevelFilterCache struct {
	limit vlog.Level
}

func (inst *innerLevelFilterCache) load(f *LevelFilter) error {
	str := f.Level
	l, err := vlog.ParseLevel(str)
	inst.limit = l
	return err
}

////////////////////////////////////////////////////////////////////////////////
// EOF
