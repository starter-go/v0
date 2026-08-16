package keys

import (
	"github.com/starter-go/application"
	"github.com/starter-go/starter"
	"github.com/starter-go/units/modules/units"
	"github.com/starter-go/v0/keys"
	"github.com/starter-go/v0/keys/gen/main4keys"
	"github.com/starter-go/v0/keys/gen/test4keys"
)

func ModuleLib() application.Module {

	mb := keys.BuildModuleForLib()

	mb.Components(main4keys.ExportComponents)

	mb.Depend(starter.Module())

	return mb.Create()
}

func ModuleTest() application.Module {

	mb := keys.BuildModuleForTest()

	mb.Components(test4keys.ExportComponents)

	mb.Depend(starter.Module())
	mb.Depend(units.Module())
	mb.Depend(ModuleLib())

	return mb.Create()
}
