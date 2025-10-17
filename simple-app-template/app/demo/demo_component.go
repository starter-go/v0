package demo

import (
	"github.com/starter-go/application"
	"github.com/starter-go/vlog"
)

////////////////////////////////////////////////////////////////////////////////

type ExampleController struct {

	//starter:component

}

func (inst *ExampleController) _impl() application.Lifecycle {
	return inst
}

func (inst *ExampleController) Life() *application.Life {
	l := new(application.Life)

	l.OnCreate = inst.onCreate
	l.OnStart = inst.onStart
	l.OnLoop = inst.onLoop
	l.OnStop = inst.onStop
	l.OnDestroy = inst.onDestroy

	return l
}

func (inst *ExampleController) onCreate() error {
	vlog.Debug("ExampleController:onCreate()")
	return nil
}

func (inst *ExampleController) onStart() error {
	vlog.Debug("ExampleController:onStart()")
	return nil
}

func (inst *ExampleController) onLoop() error {
	vlog.Debug("ExampleController:onLoop()")
	return nil
}

func (inst *ExampleController) onStop() error {
	vlog.Debug("ExampleController:onStop()")
	return nil
}

func (inst *ExampleController) onDestroy() error {
	vlog.Debug("ExampleController:onDestroy()")
	return nil
}

////////////////////////////////////////////////////////////////////////////////
