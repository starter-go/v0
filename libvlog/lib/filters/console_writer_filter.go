package filters

import (
	"net/http"
	"os"

	"github.com/starter-go/vlog"
)

type ConsoleWriterFilter struct {

	//starter:component

	_as func(vlog.FilterRegistry) //starter:as(".")

}

// DoFilter implements vlog.Filter.
func (inst *ConsoleWriterFilter) DoFilter(msg *vlog.Message, chain vlog.FilterChain) {

	const nl = "\n"
	const gate = vlog.WARN
	lv := msg.Level
	out := os.Stdout
	met := msg.Method

	if lv >= gate {
		out = os.Stderr
	}

	if met == http.MethodPost || met == http.MethodPut {
		text := msg.Text + nl
		out.WriteString(text)
	}

	chain.DoFilter(msg)
}

// ListLogFilterRegistration implements vlog.FilterRegistry.
func (inst *ConsoleWriterFilter) ListLogFilterRegistration() []*vlog.FilterRegistration {
	mfr := &vlog.FilterRegistration{
		Order:   vlog.OrderWrite,
		Name:    "ConsoleWriterFilter",
		Group:   vlog.GroupMain,
		Enabled: true,
		Filter:  inst,
	}
	return []*vlog.FilterRegistration{mfr}
}

func (inst *ConsoleWriterFilter) _impl() (vlog.FilterRegistry, vlog.Filter) {
	return inst, inst
}

////////////////////////////////////////////////////////////////////////////////
// EOF
