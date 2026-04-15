package httestrunner

import (
	"github.com/starter-go/starter"
)

func Run(c *Context) {
	r := new(runnerOuter)
	r.init(c)
	r.run()
}

////////////////////////////////////////////////////////////////////////////////

type runnerOuter struct {
	context *Context
}

func (inst *runnerOuter) init(c *Context) {
	inst.context = c
}

func (inst *runnerOuter) run() {
	err := inst.runWithError()
	if err != nil {
		panic(err)
	}
}

func (inst *runnerOuter) runWithError() error {

	c := inst.context
	a := c.Args
	m := c.Module
	i := starter.Init(a)

	i.MainModule(m)

	name := theContextBindingName
	atts := i.GetAttributes()
	atts.SetAttribute(name, c)

	return i.WithPanic(false).Run()
}

////////////////////////////////////////////////////////////////////////////////
