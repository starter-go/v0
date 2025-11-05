package rbacwebapp

import (
	"github.com/starter-go/application"
	"github.com/starter-go/libgin/modules/libgin"
	"github.com/starter-go/libgorm/modules/libgorm"
	"github.com/starter-go/module-gorm-mysql/modules/mysql"
	"github.com/starter-go/module-gorm-sqlserver/modules/sqlserver"
	"github.com/starter-go/security/modules/security"
	rbacwebapp "github.com/starter-go/v0/rbac-web-app"
	"github.com/starter-go/v0/rbac-web-app/gen/main4rbacwa"
)

// Module aka. ModuleForMain
func Module() application.Module {

	mb := rbacwebapp.NewMainModuleBuilder()

	mb.Components(main4rbacwa.ExportComponents)

	mb.Depend(libgin.Module())
	mb.Depend(libgorm.Module())
	mb.Depend(security.Module())

	mb.Depend(mysql.Module())
	mb.Depend(sqlserver.Module())

	return mb.Create()
}

func ModuleForTest() application.Module {

	mb := rbacwebapp.NewTestModuleBuilder()

	mb.Components(main4rbacwa.ExportComponents)

	mb.Depend(Module())

	return mb.Create()
}
