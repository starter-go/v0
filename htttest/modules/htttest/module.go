package htttest

import (
	"github.com/starter-go/application"
	"github.com/starter-go/starter"
	"github.com/starter-go/units/modules/units"
	"github.com/starter-go/v0/htttest"
	"github.com/starter-go/v0/htttest/gen/main4htttest"
	"github.com/starter-go/v0/htttest/gen/test4htttest"
)

func Module() application.Module {

	mb := htttest.BuildModuleForLib()
	mb.Components(main4htttest.ExportComponents)

	mb.Depend(starter.Module())
	mb.Depend(units.Module())

	return mb.Create()
}

func ModuleForTest() application.Module {

	mb := htttest.BuildModuleForTest()
	mb.Components(test4htttest.ExportComponents)

	mb.Depend(Module())

	return mb.Create()
}
