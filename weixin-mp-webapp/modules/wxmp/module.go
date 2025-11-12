package wxmp

import (
	"github.com/starter-go/application"
	"github.com/starter-go/libgin/modules/libgin"
	"github.com/starter-go/libgorm/modules/libgorm"
	"github.com/starter-go/module-gorm-mysql/modules/mysql"
	"github.com/starter-go/module-gorm-sqlserver/modules/sqlserver"
	"github.com/starter-go/starter"
	weixinmpwebapp "github.com/starter-go/v0/weixin-mp-webapp"
	"github.com/starter-go/v0/weixin-mp-webapp/gen/main4wxmp"
	"github.com/starter-go/v0/weixin-mp-webapp/gen/test4wxmp"
)

func Module() application.Module {

	mb := weixinmpwebapp.NewMainModuleBuilder()

	mb.Components(main4wxmp.ExportComponents)

	mb.Depend(starter.Module())
	mb.Depend(libgin.Module())
	mb.Depend(libgorm.Module())

	mb.Depend(mysql.Module())
	mb.Depend(sqlserver.Module())

	return mb.Create()
}

func ModuleForTest() application.Module {

	mb := weixinmpwebapp.NewTestModuleBuilder()

	mb.Components(test4wxmp.ExportComponents)

	mb.Depend(Module())

	return mb.Create()
}
