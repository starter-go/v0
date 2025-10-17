package test4security
import (
    p151f31cd9 "github.com/starter-go/v0/security/src/test/golang/unittestcases"
     "github.com/starter-go/application"
)

// type p151f31cd9.ExampleCase in package:github.com/starter-go/v0/security/src/test/golang/unittestcases
//
// id:com-151f31cd96a3f5ec-unittestcases-ExampleCase
// class:
// alias:
// scope:singleton
//
type p151f31cd96_unittestcases_ExampleCase struct {
}

func (inst* p151f31cd96_unittestcases_ExampleCase) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-151f31cd96a3f5ec-unittestcases-ExampleCase"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p151f31cd96_unittestcases_ExampleCase) new() any {
    return &p151f31cd9.ExampleCase{}
}

func (inst* p151f31cd96_unittestcases_ExampleCase) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p151f31cd9.ExampleCase)
	nop(ie, com)

	


    return nil
}


