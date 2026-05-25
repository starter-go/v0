package testwebclient4rbacwa
import (
    p581c1dbb5 "github.com/starter-go/v0/htttest"
    p48be8f24a "github.com/starter-go/v0/rbac-web-app/src/web-test/golang/client/cases"
     "github.com/starter-go/application"
)

// type p48be8f24a.CaseTryExample in package:github.com/starter-go/v0/rbac-web-app/src/web-test/golang/client/cases
//
// id:com-48be8f24a5e6d95d-cases-CaseTryExample
// class:
// alias:
// scope:singleton
//
type p48be8f24a5_cases_CaseTryExample struct {
}

func (inst* p48be8f24a5_cases_CaseTryExample) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-48be8f24a5e6d95d-cases-CaseTryExample"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p48be8f24a5_cases_CaseTryExample) new() any {
    return &p48be8f24a.CaseTryExample{}
}

func (inst* p48be8f24a5_cases_CaseTryExample) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p48be8f24a.CaseTryExample)
	nop(ie, com)

	
    com.HTTTestService = inst.getHTTTestService(ie)


    return nil
}


func (inst*p48be8f24a5_cases_CaseTryExample) getHTTTestService(ie application.InjectionExt)p581c1dbb5.Service{
    return ie.GetComponent("#alias-581c1dbb530beaf64ae39cc43b21dea9-Service").(p581c1dbb5.Service)
}



// type p48be8f24a.CaseTrySubjects in package:github.com/starter-go/v0/rbac-web-app/src/web-test/golang/client/cases
//
// id:com-48be8f24a5e6d95d-cases-CaseTrySubjects
// class:class-0ef6f2938681e99da4b0c19ce3d3fb4f-Lifecycle
// alias:
// scope:singleton
//
type p48be8f24a5_cases_CaseTrySubjects struct {
}

func (inst* p48be8f24a5_cases_CaseTrySubjects) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-48be8f24a5e6d95d-cases-CaseTrySubjects"
	r.Classes = "class-0ef6f2938681e99da4b0c19ce3d3fb4f-Lifecycle"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p48be8f24a5_cases_CaseTrySubjects) new() any {
    return &p48be8f24a.CaseTrySubjects{}
}

func (inst* p48be8f24a5_cases_CaseTrySubjects) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p48be8f24a.CaseTrySubjects)
	nop(ie, com)

	
    com.HTTTestService = inst.getHTTTestService(ie)


    return nil
}


func (inst*p48be8f24a5_cases_CaseTrySubjects) getHTTTestService(ie application.InjectionExt)p581c1dbb5.Service{
    return ie.GetComponent("#alias-581c1dbb530beaf64ae39cc43b21dea9-Service").(p581c1dbb5.Service)
}



// type p48be8f24a.CaseTryUseAuth in package:github.com/starter-go/v0/rbac-web-app/src/web-test/golang/client/cases
//
// id:com-48be8f24a5e6d95d-cases-CaseTryUseAuth
// class:
// alias:
// scope:singleton
//
type p48be8f24a5_cases_CaseTryUseAuth struct {
}

func (inst* p48be8f24a5_cases_CaseTryUseAuth) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-48be8f24a5e6d95d-cases-CaseTryUseAuth"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p48be8f24a5_cases_CaseTryUseAuth) new() any {
    return &p48be8f24a.CaseTryUseAuth{}
}

func (inst* p48be8f24a5_cases_CaseTryUseAuth) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p48be8f24a.CaseTryUseAuth)
	nop(ie, com)

	
    com.HTTTestService = inst.getHTTTestService(ie)


    return nil
}


func (inst*p48be8f24a5_cases_CaseTryUseAuth) getHTTTestService(ie application.InjectionExt)p581c1dbb5.Service{
    return ie.GetComponent("#alias-581c1dbb530beaf64ae39cc43b21dea9-Service").(p581c1dbb5.Service)
}



// type p48be8f24a.CaseTryUseSessions in package:github.com/starter-go/v0/rbac-web-app/src/web-test/golang/client/cases
//
// id:com-48be8f24a5e6d95d-cases-CaseTryUseSessions
// class:
// alias:
// scope:singleton
//
type p48be8f24a5_cases_CaseTryUseSessions struct {
}

func (inst* p48be8f24a5_cases_CaseTryUseSessions) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-48be8f24a5e6d95d-cases-CaseTryUseSessions"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p48be8f24a5_cases_CaseTryUseSessions) new() any {
    return &p48be8f24a.CaseTryUseSessions{}
}

func (inst* p48be8f24a5_cases_CaseTryUseSessions) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p48be8f24a.CaseTryUseSessions)
	nop(ie, com)

	
    com.HTTTestService = inst.getHTTTestService(ie)


    return nil
}


func (inst*p48be8f24a5_cases_CaseTryUseSessions) getHTTTestService(ie application.InjectionExt)p581c1dbb5.Service{
    return ie.GetComponent("#alias-581c1dbb530beaf64ae39cc43b21dea9-Service").(p581c1dbb5.Service)
}


