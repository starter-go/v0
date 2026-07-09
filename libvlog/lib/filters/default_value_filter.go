package filters

import (
	"net/http"

	"github.com/starter-go/vlog"
)

type DefaultValueFilter struct {

	//starter:component

	_as func(vlog.FilterRegistry) //starter:as(".")

}

// DoFilter implements vlog.Filter.
func (inst *DefaultValueFilter) DoFilter(msg *vlog.Message, chain vlog.FilterChain) {

	met := msg.Method
	lv := msg.Level
	sender := msg.Sender

	if met == "" {
		met = http.MethodPost
		msg.Method = met
	}

	if lv == 0 {
		lv = vlog.INFO
		msg.Level = lv
	}

	if sender == nil {
		sender = inst
		msg.Sender = sender
	}

	chain.DoFilter(msg)
}

// ListLogFilterRegistration implements vlog.FilterRegistry.
func (inst *DefaultValueFilter) ListLogFilterRegistration() []*vlog.FilterRegistration {
	mfr := &vlog.FilterRegistration{
		Order:   1,
		Name:    "DefaultValueFilter",
		Group:   vlog.GroupMain,
		Enabled: true,
		Filter:  inst,
	}
	return []*vlog.FilterRegistration{mfr}
}

func (inst *DefaultValueFilter) _impl() (vlog.FilterRegistry, vlog.Filter) {
	return inst, inst
}

////////////////////////////////////////////////////////////////////////////////
// EOF
