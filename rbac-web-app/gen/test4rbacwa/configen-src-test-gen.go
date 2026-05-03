package test4rbacwa
import (
    pd1a916a20 "github.com/starter-go/libgin"
    p45dbe3237 "github.com/starter-go/v0/rbac-web-app/app/classes/subjects"
    p2d258e798 "github.com/starter-go/v0/rbac-web-app/app/classes/users"
    pfabbd91b2 "github.com/starter-go/v0/rbac-web-app/src/test/golang/unittestcases"
    p78d6a2ce1 "github.com/starter-go/v0/rbac-web-app/src/test/golang/web_unit/wuclient"
    p90f4aa59a "github.com/starter-go/v0/rbac-web-app/src/test/golang/web_unit/wuserver"
     "github.com/starter-go/application"
)

// type pfabbd91b2.ExampleCase in package:github.com/starter-go/v0/rbac-web-app/src/test/golang/unittestcases
//
// id:com-fabbd91b2d06fe21-unittestcases-ExampleCase
// class:
// alias:
// scope:singleton
//
type pfabbd91b2d_unittestcases_ExampleCase struct {
}

func (inst* pfabbd91b2d_unittestcases_ExampleCase) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-fabbd91b2d06fe21-unittestcases-ExampleCase"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pfabbd91b2d_unittestcases_ExampleCase) new() any {
    return &pfabbd91b2.ExampleCase{}
}

func (inst* pfabbd91b2d_unittestcases_ExampleCase) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pfabbd91b2.ExampleCase)
	nop(ie, com)

	


    return nil
}



// type pfabbd91b2.CaseTrySubject in package:github.com/starter-go/v0/rbac-web-app/src/test/golang/unittestcases
//
// id:com-fabbd91b2d06fe21-unittestcases-CaseTrySubject
// class:class-0ef6f2938681e99da4b0c19ce3d3fb4f-Lifecycle
// alias:
// scope:singleton
//
type pfabbd91b2d_unittestcases_CaseTrySubject struct {
}

func (inst* pfabbd91b2d_unittestcases_CaseTrySubject) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-fabbd91b2d06fe21-unittestcases-CaseTrySubject"
	r.Classes = "class-0ef6f2938681e99da4b0c19ce3d3fb4f-Lifecycle"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pfabbd91b2d_unittestcases_CaseTrySubject) new() any {
    return &pfabbd91b2.CaseTrySubject{}
}

func (inst* pfabbd91b2d_unittestcases_CaseTrySubject) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pfabbd91b2.CaseTrySubject)
	nop(ie, com)

	
    com.Subjects = inst.getSubjects(ie)


    return nil
}


func (inst*pfabbd91b2d_unittestcases_CaseTrySubject) getSubjects(ie application.InjectionExt)p45dbe3237.Service{
    return ie.GetComponent("#alias-45dbe3237457ab5b761a911af615329f-Service").(p45dbe3237.Service)
}



// type p78d6a2ce1.WebUnitTestSessionRunner in package:github.com/starter-go/v0/rbac-web-app/src/test/golang/web_unit/wuclient
//
// id:com-78d6a2ce14acc73b-wuclient-WebUnitTestSessionRunner
// class:
// alias:
// scope:singleton
//
type p78d6a2ce14_wuclient_WebUnitTestSessionRunner struct {
}

func (inst* p78d6a2ce14_wuclient_WebUnitTestSessionRunner) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-78d6a2ce14acc73b-wuclient-WebUnitTestSessionRunner"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p78d6a2ce14_wuclient_WebUnitTestSessionRunner) new() any {
    return &p78d6a2ce1.WebUnitTestSessionRunner{}
}

func (inst* p78d6a2ce14_wuclient_WebUnitTestSessionRunner) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p78d6a2ce1.WebUnitTestSessionRunner)
	nop(ie, com)

	


    return nil
}



// type p90f4aa59a.WebUnitTestSessionController in package:github.com/starter-go/v0/rbac-web-app/src/test/golang/web_unit/wuserver
//
// id:com-90f4aa59a9e2a713-wuserver-WebUnitTestSessionController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p90f4aa59a9_wuserver_WebUnitTestSessionController struct {
}

func (inst* p90f4aa59a9_wuserver_WebUnitTestSessionController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-90f4aa59a9e2a713-wuserver-WebUnitTestSessionController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p90f4aa59a9_wuserver_WebUnitTestSessionController) new() any {
    return &p90f4aa59a.WebUnitTestSessionController{}
}

func (inst* p90f4aa59a9_wuserver_WebUnitTestSessionController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p90f4aa59a.WebUnitTestSessionController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p90f4aa59a9_wuserver_WebUnitTestSessionController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p90f4aa59a9_wuserver_WebUnitTestSessionController) getService(ie application.InjectionExt)p2d258e798.Service{
    return ie.GetComponent("#alias-2d258e79845135c43979b78bd9f1f74e-Service").(p2d258e798.Service)
}


