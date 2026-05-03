package subjects

import (
	"github.com/starter-go/application"
	"github.com/starter-go/libgorm/modules/libgorm"
	"github.com/starter-go/module-gorm-mysql/modules/mysql"
	"github.com/starter-go/security/modules/security"
	"github.com/starter-go/starter"
	"github.com/starter-go/units/modules/units"
	"github.com/starter-go/v0/subjects"
	"github.com/starter-go/v0/subjects/gen/main4subjects"
	"github.com/starter-go/v0/subjects/gen/test4subjects"
)

func ModuleForLib() application.Module {

	mb := subjects.BuildModuleForLib()
	mb.Components(main4subjects.ExportComponents)

	mb.Depend(starter.Module())
	mb.Depend(security.Module())
	mb.Depend(libgorm.Module())
	mb.Depend(mysql.Module())

	return mb.Create()
}

func ModuleForTest() application.Module {

	mb := subjects.BuildModuleForTest()
	mb.Components(test4subjects.ExportComponents)

	mb.Depend(ModuleForLib())
	mb.Depend(units.Module())

	return mb.Create()
}
