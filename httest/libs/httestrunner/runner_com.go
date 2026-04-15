package httestrunner

import (
	"fmt"

	"github.com/starter-go/application"
	"github.com/starter-go/v0/httest/libs/httestcore"
)

////////////////////////////////////////////////////////////////////////////////

type RunnerCom struct {

	//starter:component

	AC application.Context //starter:inject("context")

	rb *httestcore.RequestBuilder
}

func (inst *RunnerCom) _impl() application.Lifecycle {
	return inst
}

func (inst *RunnerCom) Life() *application.Life {
	l := new(application.Life)

	l.OnCreate = inst.onCreate
	l.OnStart = inst.onStart
	l.OnLoop = inst.onLoop
	l.OnStop = inst.onStop
	l.OnDestroy = inst.onDestroy

	return l
}

func (inst *RunnerCom) onCreate() error {
	// vlog.Debug("RunnerCom:onCreate()")

	rb, err := inst.innerLoadRequestBuilder()
	inst.rb = rb
	return err
}

func (inst *RunnerCom) onStart() error {
	// vlog.Debug("RunnerCom:onStart()")
	return nil
}

func (inst *RunnerCom) onLoop() error {

	// vlog.Debug("RunnerCom:onLoop()")

	ctx, err := inst.innerGetContext()
	if err != nil {
		return err
	}

	inner := new(runnerInner)
	inner.init(ctx)
	return inner.run()
}

func (inst *RunnerCom) onStop() error {
	// vlog.Debug("RunnerCom:onStop()")
	return nil
}

func (inst *RunnerCom) onDestroy() error {
	// vlog.Debug("RunnerCom:onDestroy()")
	return nil
}

func (inst *RunnerCom) innerLoadRequestBuilder() (*httestcore.RequestBuilder, error) {

	rb := inst.rb
	if rb != nil {
		return rb, nil
	}

	ctx, err := inst.innerGetContext()
	if err != nil {
		return nil, err
	}

	rb = ctx.RequestBuilder
	if rb == nil {
		return nil, fmt.Errorf("no httestcore.RequestBuilder in the context")
	}

	inst.rb = rb
	return rb, nil
}

func (inst *RunnerCom) innerGetContext() (*Context, error) {

	name := theContextBindingName
	o1 := inst.AC.GetAttributes().GetAttribute(name)
	ctx := o1.(*Context)
	return ctx, nil
}

////////////////////////////////////////////////////////////////////////////////
