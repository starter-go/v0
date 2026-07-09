package filters

import (
	"fmt"

	"github.com/starter-go/vlog"
)

type FormatterFilter struct {

	//starter:component

	_as func(vlog.FilterRegistry) //starter:as(".")

	HeadFormat string //starter:inject("${vlog.formatters.default.format}")

	cachedFmt *vlog.MessageHeadFormatter
}

// DoFilter implements vlog.Filter.
func (inst *FormatterFilter) DoFilter(msg *vlog.Message, chain vlog.FilterChain) {

	txt1 := inst.innerFormatMsgHeadText(msg)
	txt2 := inst.innerFormatMsgBodyText(msg)
	msg.Text = txt1 + " " + txt2

	chain.DoFilter(msg)
}

// ListLogFilterRegistration implements vlog.FilterRegistry.
func (inst *FormatterFilter) ListLogFilterRegistration() []*vlog.FilterRegistration {
	mfr := &vlog.FilterRegistration{
		Order:   vlog.OrderFormat,
		Name:    "FormatterFilter",
		Group:   vlog.GroupMain,
		Enabled: true,
		Filter:  inst,
	}
	return []*vlog.FilterRegistration{mfr}
}

func (inst *FormatterFilter) innerGetMsgHeadFormatter() *vlog.MessageHeadFormatter {
	ff := inst.cachedFmt
	if ff == nil {
		ff = new(vlog.MessageHeadFormatter)
		ff.Init(inst.HeadFormat)
		inst.cachedFmt = ff
	}
	return ff
}

func (inst *FormatterFilter) innerFormatMsgHeadText(msg *vlog.Message) string {
	f := inst.innerGetMsgHeadFormatter()
	return f.Format(msg)
}

func (inst *FormatterFilter) innerFormatMsgBodyText(msg *vlog.Message) string {
	f := msg.Format
	a := msg.Arguments
	return fmt.Sprintf(f, a...)
}

func (inst *FormatterFilter) _impl() (vlog.FilterRegistry, vlog.Filter) {
	return inst, inst
}

////////////////////////////////////////////////////////////////////////////////
// EOF
