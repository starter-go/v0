package test4subjects
import (
    pfd2c28477 "github.com/starter-go/v0/subjects"
    p85899ef78 "github.com/starter-go/v0/subjects/core/classes/sessions"
    p8158f311c "github.com/starter-go/v0/subjects/src/test/golang/unittestcases"
     "github.com/starter-go/application"
)

// type p8158f311c.ExampleCase in package:github.com/starter-go/v0/subjects/src/test/golang/unittestcases
//
// id:com-8158f311c02d7656-unittestcases-ExampleCase
// class:
// alias:
// scope:singleton
//
type p8158f311c0_unittestcases_ExampleCase struct {
}

func (inst* p8158f311c0_unittestcases_ExampleCase) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-8158f311c02d7656-unittestcases-ExampleCase"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p8158f311c0_unittestcases_ExampleCase) new() any {
    return &p8158f311c.ExampleCase{}
}

func (inst* p8158f311c0_unittestcases_ExampleCase) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p8158f311c.ExampleCase)
	nop(ie, com)

	
    com.Enabled = inst.getEnabled(ie)


    return nil
}


func (inst*p8158f311c0_unittestcases_ExampleCase) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${test-case.example.enabled}")
}



// type p8158f311c.TrySessionService in package:github.com/starter-go/v0/subjects/src/test/golang/unittestcases
//
// id:com-8158f311c02d7656-unittestcases-TrySessionService
// class:
// alias:
// scope:singleton
//
type p8158f311c0_unittestcases_TrySessionService struct {
}

func (inst* p8158f311c0_unittestcases_TrySessionService) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-8158f311c02d7656-unittestcases-TrySessionService"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p8158f311c0_unittestcases_TrySessionService) new() any {
    return &p8158f311c.TrySessionService{}
}

func (inst* p8158f311c0_unittestcases_TrySessionService) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p8158f311c.TrySessionService)
	nop(ie, com)

	
    com.Service = inst.getService(ie)
    com.Enabled = inst.getEnabled(ie)


    return nil
}


func (inst*p8158f311c0_unittestcases_TrySessionService) getService(ie application.InjectionExt)p85899ef78.Service{
    return ie.GetComponent("#alias-85899ef785c3f033c4d8293618ff61de-Service").(p85899ef78.Service)
}


func (inst*p8158f311c0_unittestcases_TrySessionService) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${test-case.session-service.enabled}")
}



// type p8158f311c.CaseToTrySubject in package:github.com/starter-go/v0/subjects/src/test/golang/unittestcases
//
// id:com-8158f311c02d7656-unittestcases-CaseToTrySubject
// class:
// alias:
// scope:singleton
//
type p8158f311c0_unittestcases_CaseToTrySubject struct {
}

func (inst* p8158f311c0_unittestcases_CaseToTrySubject) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-8158f311c02d7656-unittestcases-CaseToTrySubject"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p8158f311c0_unittestcases_CaseToTrySubject) new() any {
    return &p8158f311c.CaseToTrySubject{}
}

func (inst* p8158f311c0_unittestcases_CaseToTrySubject) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p8158f311c.CaseToTrySubject)
	nop(ie, com)

	
    com.ChainHolder = inst.getChainHolder(ie)
    com.Enabled = inst.getEnabled(ie)


    return nil
}


func (inst*p8158f311c0_unittestcases_CaseToTrySubject) getChainHolder(ie application.InjectionExt)pfd2c28477.FilterChainHolder{
    return ie.GetComponent("#alias-fd2c28477d8555ea1fa4190037afa453-FilterChainHolder").(pfd2c28477.FilterChainHolder)
}


func (inst*p8158f311c0_unittestcases_CaseToTrySubject) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${test-case.subjects.enabled}")
}


