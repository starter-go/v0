package security

import (
	"github.com/starter-go/application"
	"github.com/starter-go/starter"
	"github.com/starter-go/v0/security"
	"github.com/starter-go/v0/security/gen/main4security"
	"github.com/starter-go/v0/security/gen/test4security"
)

func Module() application.Module {

	mb := security.NewMainModuleBuilder()

	mb.Components(main4security.ExportComponents)
	mb.Depend(starter.Module())

	return mb.Create()
}

func ModuleForTest() application.Module {

	mb := security.NewTestModuleBuilder()

	mb.Components(test4security.ExportComponents)
	mb.Depend(Module())

	return mb.Create()
}
