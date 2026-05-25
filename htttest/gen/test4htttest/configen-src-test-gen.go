package test4htttest
import (
    p581c1dbb5 "github.com/starter-go/v0/htttest"
    p9ef195e5c "github.com/starter-go/v0/htttest/src/test/golang/unittestcases"
     "github.com/starter-go/application"
)

// type p9ef195e5c.ExampleCase in package:github.com/starter-go/v0/htttest/src/test/golang/unittestcases
//
// id:com-9ef195e5c6598a6b-unittestcases-ExampleCase
// class:
// alias:
// scope:singleton
//
type p9ef195e5c6_unittestcases_ExampleCase struct {
}

func (inst* p9ef195e5c6_unittestcases_ExampleCase) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-9ef195e5c6598a6b-unittestcases-ExampleCase"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p9ef195e5c6_unittestcases_ExampleCase) new() any {
    return &p9ef195e5c.ExampleCase{}
}

func (inst* p9ef195e5c6_unittestcases_ExampleCase) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p9ef195e5c.ExampleCase)
	nop(ie, com)

	
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p9ef195e5c6_unittestcases_ExampleCase) getService(ie application.InjectionExt)p581c1dbb5.Service{
    return ie.GetComponent("#alias-581c1dbb530beaf64ae39cc43b21dea9-Service").(p581c1dbb5.Service)
}



// type p9ef195e5c.CaseTryContentDecoder in package:github.com/starter-go/v0/htttest/src/test/golang/unittestcases
//
// id:com-9ef195e5c6598a6b-unittestcases-CaseTryContentDecoder
// class:
// alias:
// scope:singleton
//
type p9ef195e5c6_unittestcases_CaseTryContentDecoder struct {
}

func (inst* p9ef195e5c6_unittestcases_CaseTryContentDecoder) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-9ef195e5c6598a6b-unittestcases-CaseTryContentDecoder"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p9ef195e5c6_unittestcases_CaseTryContentDecoder) new() any {
    return &p9ef195e5c.CaseTryContentDecoder{}
}

func (inst* p9ef195e5c6_unittestcases_CaseTryContentDecoder) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p9ef195e5c.CaseTryContentDecoder)
	nop(ie, com)

	
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p9ef195e5c6_unittestcases_CaseTryContentDecoder) getService(ie application.InjectionExt)p581c1dbb5.Service{
    return ie.GetComponent("#alias-581c1dbb530beaf64ae39cc43b21dea9-Service").(p581c1dbb5.Service)
}



// type p9ef195e5c.CaseTryPostAuth in package:github.com/starter-go/v0/htttest/src/test/golang/unittestcases
//
// id:com-9ef195e5c6598a6b-unittestcases-CaseTryPostAuth
// class:
// alias:
// scope:singleton
//
type p9ef195e5c6_unittestcases_CaseTryPostAuth struct {
}

func (inst* p9ef195e5c6_unittestcases_CaseTryPostAuth) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-9ef195e5c6598a6b-unittestcases-CaseTryPostAuth"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p9ef195e5c6_unittestcases_CaseTryPostAuth) new() any {
    return &p9ef195e5c.CaseTryPostAuth{}
}

func (inst* p9ef195e5c6_unittestcases_CaseTryPostAuth) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p9ef195e5c.CaseTryPostAuth)
	nop(ie, com)

	
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p9ef195e5c6_unittestcases_CaseTryPostAuth) getService(ie application.InjectionExt)p581c1dbb5.Service{
    return ie.GetComponent("#alias-581c1dbb530beaf64ae39cc43b21dea9-Service").(p581c1dbb5.Service)
}


