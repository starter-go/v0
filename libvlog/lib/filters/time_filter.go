package filters

import (
	"time"

	"github.com/starter-go/vlog"
)

type TimeFilter struct {

	//starter:component

	_as func(vlog.FilterRegistry) //starter:as(".")

}

// DoFilter implements vlog.Filter.
func (inst *TimeFilter) DoFilter(msg *vlog.Message, chain vlog.FilterChain) {

	msg.Timestamp = time.Now()

	chain.DoFilter(msg)
}

// ListLogFilterRegistration implements vlog.FilterRegistry.
func (inst *TimeFilter) ListLogFilterRegistration() []*vlog.FilterRegistration {
	mfr := &vlog.FilterRegistration{
		Order:   vlog.OrderTime,
		Name:    "TimeFilter",
		Group:   vlog.GroupMain,
		Enabled: true,
		Filter:  inst,
	}
	return []*vlog.FilterRegistration{mfr}
}

func (inst *TimeFilter) _impl() (vlog.FilterRegistry, vlog.Filter) {
	return inst, inst
}

////////////////////////////////////////////////////////////////////////////////
// EOF
