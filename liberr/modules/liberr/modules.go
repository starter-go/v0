package liberr

import (
	"github.com/starter-go/application"
	"github.com/starter-go/starter"
	"github.com/starter-go/v0/liberr"
	"github.com/starter-go/v0/liberr/gen/main4liberr"
	"github.com/starter-go/v0/liberr/gen/test4liberr"
)

func Module() application.Module {
	mb := liberr.NewModuleForLib()

	mb.Components(main4liberr.ExportComponents)

	mb.Depend(starter.Module())

	return mb.Create()
}

func ModuleForTest() application.Module {
	mb := liberr.NewModuleForTest()

	mb.Components(test4liberr.ExportComponents)

	mb.Depend(Module())

	return mb.Create()
}
