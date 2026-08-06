package liberr

import (
	"github.com/starter-go/application"
	"github.com/starter-go/i18n/modules/i18n"
	"github.com/starter-go/starter"
	"github.com/starter-go/units/modules/units"
	"github.com/starter-go/v0/liberr"
	"github.com/starter-go/v0/liberr/gen/main4liberr"
	"github.com/starter-go/v0/liberr/gen/test4liberr"
)

func Module() application.Module {
	mb := liberr.NewModuleForLib()

	mb.Components(main4liberr.ExportComponents)

	mb.Depend(starter.Module())
	mb.Depend(i18n.Module())

	return mb.Create()
}

func ModuleForTest() application.Module {
	mb := liberr.NewModuleForTest()

	mb.Components(test4liberr.ExportComponents)

	mb.Depend(Module())
	mb.Depend(units.Module())

	return mb.Create()
}
