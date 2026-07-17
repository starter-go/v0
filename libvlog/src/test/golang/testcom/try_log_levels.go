package testcom

import (
	"context"
	"fmt"

	"github.com/starter-go/units"
	"github.com/starter-go/vlog"
)

type TryLogLevels struct {

	//starter:component

	_as func(units.Unit) //starter:as(".")

}

// ListRegistrations implements units.Unit.
func (inst *TryLogLevels) ListRegistrations(list []*units.Registration) []*units.Registration {

	ur1 := &units.Registration{
		Do:      inst.runLogWithLevels,
		Name:    "try-log-levels",
		Enabled: true,
	}

	ur2 := &units.Registration{
		Do:      inst.runListLevels,
		Name:    "list-levels",
		Enabled: true,
	}

	list = append(list, ur1, ur2)
	return list
}

func (inst *TryLogLevels) runLogWithLevels(cc context.Context) error {

	vlog.Fatal("[log level:'%s' ]", vlog.FATAL.String())
	vlog.Error("[log level:'%s' ]", vlog.ERROR.String())
	vlog.Warn("[log level:'%s' ]", vlog.WARN.String())

	vlog.Info("[log level:'%s' ]", vlog.INFO.String())

	vlog.Debug("[log level:'%s' ]", vlog.DEBUG.String())
	vlog.Trace("[log level:'%s' ]", vlog.TRACE.String())

	return nil
}

func (inst *TryLogLevels) runListLevels(cc context.Context) error {

	list := []vlog.Level{
		vlog.FATAL,
		vlog.ERROR,
		vlog.WARN,
		vlog.INFO,
		vlog.DEBUG,
		vlog.TRACE,
	}

	fmt.Println("vlog.levels.list - begin")
	for _, lv := range list {
		en := false
		name := lv.String()
		switch lv {
		case vlog.FATAL:
			en = vlog.IsFatalEnabled()
		case vlog.ERROR:
			en = vlog.IsErrorEnabled()
		case vlog.WARN:
			en = vlog.IsWarnEnabled()
		case vlog.INFO:
			en = vlog.IsInfoEnabled()
		case vlog.DEBUG:
			en = vlog.IsDebugEnabled()
		case vlog.TRACE:
			en = vlog.IsTraceEnabled()
		}
		fmt.Printf("[vlog level:%s enabled:%v ]\n", name, en)
	}
	fmt.Println("vlog.levels.list - end")

	return nil
}

func (inst *TryLogLevels) _impl() units.Unit {
	return inst
}
