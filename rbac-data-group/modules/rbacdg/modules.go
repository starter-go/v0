package rbacdg

import (
	"github.com/starter-go/application"
	"github.com/starter-go/libgorm/modules/libgorm"
	"github.com/starter-go/module-gorm-mysql/modules/mysql4libgorm"
	"github.com/starter-go/rbac/modules/librbac"
	"github.com/starter-go/security/modules/security"
	"github.com/starter-go/starter"
	"github.com/starter-go/units/modules/units"
	rbacdatagroup "github.com/starter-go/v0/rbac-data-group"
	"github.com/starter-go/v0/rbac-data-group/gen/main4rbacdg"
	"github.com/starter-go/v0/rbac-data-group/gen/test4rbacdg"
)

// main (default) :
func Module() application.Module {
	return ModuleForLib()
}

// Lib :
func ModuleForLib() application.Module {

	mb := rbacdatagroup.BuildModuleForLib()

	mb.Components(main4rbacdg.ExportComponents)

	mb.Depend(starter.Module())
	mb.Depend(libgorm.Module())
	mb.Depend(librbac.Module())
	mb.Depend(security.Module())

	return mb.Create()
}

// test :
func ModuleForTest() application.Module {

	mb := rbacdatagroup.BuildModuleForTest()

	mb.Components(test4rbacdg.ExportComponents)

	mb.Depend(Module())
	mb.Depend(units.Module())
	mb.Depend(mysql4libgorm.Module())

	return mb.Create()
}
