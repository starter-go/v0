package testwebserver4rbacwa
import (
    pd1a916a20 "github.com/starter-go/libgin"
    p2d258e798 "github.com/starter-go/v0/rbac-web-app/app/classes/users"
    p819ff5697 "github.com/starter-go/v0/rbac-web-app/src/web-test/golang/server/controllers"
     "github.com/starter-go/application"
)

// type p819ff5697.WebUnitTestSessionController in package:github.com/starter-go/v0/rbac-web-app/src/web-test/golang/server/controllers
//
// id:com-819ff569706f8b38-controllers-WebUnitTestSessionController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p819ff56970_controllers_WebUnitTestSessionController struct {
}

func (inst* p819ff56970_controllers_WebUnitTestSessionController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-819ff569706f8b38-controllers-WebUnitTestSessionController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p819ff56970_controllers_WebUnitTestSessionController) new() any {
    return &p819ff5697.WebUnitTestSessionController{}
}

func (inst* p819ff56970_controllers_WebUnitTestSessionController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p819ff5697.WebUnitTestSessionController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p819ff56970_controllers_WebUnitTestSessionController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p819ff56970_controllers_WebUnitTestSessionController) getService(ie application.InjectionExt)p2d258e798.Service{
    return ie.GetComponent("#alias-2d258e79845135c43979b78bd9f1f74e-Service").(p2d258e798.Service)
}


