package libdao

import (
	"github.com/starter-go/application"
	"github.com/starter-go/starter"
	"github.com/starter-go/units/modules/units"
	"github.com/starter-go/v0/libdao"
	"github.com/starter-go/v0/libdao/gen/main4libdao"
	"github.com/starter-go/v0/libdao/gen/test4libdao"
)

func Module() application.Module {
	return ModuleForMain()
}

func ModuleForMain() application.Module {

	mb := libdao.BuildModuleForMain()
	mb.Components(main4libdao.ExportComponents)

	mb.Depend(starter.Module())

	return mb.Create()
}

func ModuleForTest() application.Module {

	mb := libdao.BuildModuleForTest()
	mb.Components(test4libdao.ExportComponents)

	mb.Depend(ModuleForMain())
	mb.Depend(units.Module())

	return mb.Create()
}
