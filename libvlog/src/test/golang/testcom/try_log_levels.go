package testcom

import (
	"context"

	"github.com/starter-go/units"
	"github.com/starter-go/vlog"
)

type TryLogLevels struct {

	//starter:component

	_as func(units.Unit) //starter:as(".")

}

// ListRegistrations implements units.Unit.
func (inst *TryLogLevels) ListRegistrations(list []*units.Registration) []*units.Registration {

	ur := &units.Registration{
		Do:      inst.run,
		Name:    "try-log-levels",
		Enabled: true,
	}

	list = append(list, ur)
	return list
}

func (inst *TryLogLevels) run(cc context.Context) error {

	vlog.Fatal("[log level:'%s' ]", vlog.FATAL.String())
	vlog.Error("[log level:'%s' ]", vlog.ERROR.String())
	vlog.Warn("[log level:'%s' ]", vlog.WARN.String())

	vlog.Info("[log level:'%s' ]", vlog.INFO.String())

	vlog.Debug("[log level:'%s' ]", vlog.DEBUG.String())
	vlog.Trace("[log level:'%s' ]", vlog.TRACE.String())

	return nil
}

func (inst *TryLogLevels) _impl() units.Unit {
	return inst
}
