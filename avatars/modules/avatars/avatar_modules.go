package avatars

import (
	"github.com/starter-go/application"
	"github.com/starter-go/libgin/modules/libgin"

	"github.com/starter-go/mimetypes/modules/mimetypes"
	"github.com/starter-go/starter"
	"github.com/starter-go/units/modules/units"
	"github.com/starter-go/v0/avatars"
	"github.com/starter-go/v0/avatars/gen/main4avatars"
	"github.com/starter-go/v0/avatars/gen/test4avatars"
)

func Module() application.Module {
	return ModuleForMain()
}

func ModuleForMain() application.Module {

	mb := avatars.BuildModuleForMain()

	mb.Components(main4avatars.ExportComponents)

	mb.Depend(starter.Module())
	mb.Depend(libgin.Module())
	mb.Depend(mimetypes.ModuleForCommon())

	return mb.Create()
}

func ModuleForTest() application.Module {

	mb := avatars.BuildModuleForTest()

	mb.Components(test4avatars.ExportComponents)

	mb.Depend(ModuleForMain())
	mb.Depend(units.Module())

	return mb.Create()
}
