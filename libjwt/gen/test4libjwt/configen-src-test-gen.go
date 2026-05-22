package test4libjwt
import (
    p40146ff25 "github.com/starter-go/v0/libjwt/api/jwt"
    p390ccebeb "github.com/starter-go/v0/libjwt/src/test/golang/unittestcases"
     "github.com/starter-go/application"
)

// type p390ccebeb.ExampleCase in package:github.com/starter-go/v0/libjwt/src/test/golang/unittestcases
//
// id:com-390ccebeb6a00052-unittestcases-ExampleCase
// class:
// alias:
// scope:singleton
//
type p390ccebeb6_unittestcases_ExampleCase struct {
}

func (inst* p390ccebeb6_unittestcases_ExampleCase) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-390ccebeb6a00052-unittestcases-ExampleCase"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p390ccebeb6_unittestcases_ExampleCase) new() any {
    return &p390ccebeb.ExampleCase{}
}

func (inst* p390ccebeb6_unittestcases_ExampleCase) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p390ccebeb.ExampleCase)
	nop(ie, com)

	
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p390ccebeb6_unittestcases_ExampleCase) getService(ie application.InjectionExt)p40146ff25.Service{
    return ie.GetComponent("#alias-40146ff2506e686679e1be8b766901e1-Service").(p40146ff25.Service)
}



// type p390ccebeb.CaseTryAdapter in package:github.com/starter-go/v0/libjwt/src/test/golang/unittestcases
//
// id:com-390ccebeb6a00052-unittestcases-CaseTryAdapter
// class:
// alias:
// scope:singleton
//
type p390ccebeb6_unittestcases_CaseTryAdapter struct {
}

func (inst* p390ccebeb6_unittestcases_CaseTryAdapter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-390ccebeb6a00052-unittestcases-CaseTryAdapter"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p390ccebeb6_unittestcases_CaseTryAdapter) new() any {
    return &p390ccebeb.CaseTryAdapter{}
}

func (inst* p390ccebeb6_unittestcases_CaseTryAdapter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p390ccebeb.CaseTryAdapter)
	nop(ie, com)

	
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p390ccebeb6_unittestcases_CaseTryAdapter) getService(ie application.InjectionExt)p40146ff25.Service{
    return ie.GetComponent("#alias-40146ff2506e686679e1be8b766901e1-Service").(p40146ff25.Service)
}



// type p390ccebeb.CaseTryJwtCodec in package:github.com/starter-go/v0/libjwt/src/test/golang/unittestcases
//
// id:com-390ccebeb6a00052-unittestcases-CaseTryJwtCodec
// class:
// alias:
// scope:singleton
//
type p390ccebeb6_unittestcases_CaseTryJwtCodec struct {
}

func (inst* p390ccebeb6_unittestcases_CaseTryJwtCodec) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-390ccebeb6a00052-unittestcases-CaseTryJwtCodec"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p390ccebeb6_unittestcases_CaseTryJwtCodec) new() any {
    return &p390ccebeb.CaseTryJwtCodec{}
}

func (inst* p390ccebeb6_unittestcases_CaseTryJwtCodec) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p390ccebeb.CaseTryJwtCodec)
	nop(ie, com)

	
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p390ccebeb6_unittestcases_CaseTryJwtCodec) getService(ie application.InjectionExt)p40146ff25.Service{
    return ie.GetComponent("#alias-40146ff2506e686679e1be8b766901e1-Service").(p40146ff25.Service)
}



// type p390ccebeb.CaseTryKeyLoader in package:github.com/starter-go/v0/libjwt/src/test/golang/unittestcases
//
// id:com-390ccebeb6a00052-unittestcases-CaseTryKeyLoader
// class:
// alias:
// scope:singleton
//
type p390ccebeb6_unittestcases_CaseTryKeyLoader struct {
}

func (inst* p390ccebeb6_unittestcases_CaseTryKeyLoader) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-390ccebeb6a00052-unittestcases-CaseTryKeyLoader"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p390ccebeb6_unittestcases_CaseTryKeyLoader) new() any {
    return &p390ccebeb.CaseTryKeyLoader{}
}

func (inst* p390ccebeb6_unittestcases_CaseTryKeyLoader) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p390ccebeb.CaseTryKeyLoader)
	nop(ie, com)

	
    com.Loader = inst.getLoader(ie)


    return nil
}


func (inst*p390ccebeb6_unittestcases_CaseTryKeyLoader) getLoader(ie application.InjectionExt)p40146ff25.KeyLoader{
    return ie.GetComponent("#alias-40146ff2506e686679e1be8b766901e1-KeyLoader").(p40146ff25.KeyLoader)
}


