package test4wat
import (
    p4876c9c6d "github.com/starter-go/v0/security-web-app/src/test/golang/unittestcases"
     "github.com/starter-go/application"
)

// type p4876c9c6d.ExampleCase in package:github.com/starter-go/v0/security-web-app/src/test/golang/unittestcases
//
// id:com-4876c9c6d67f4da3-unittestcases-ExampleCase
// class:
// alias:
// scope:singleton
//
type p4876c9c6d6_unittestcases_ExampleCase struct {
}

func (inst* p4876c9c6d6_unittestcases_ExampleCase) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-4876c9c6d67f4da3-unittestcases-ExampleCase"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p4876c9c6d6_unittestcases_ExampleCase) new() any {
    return &p4876c9c6d.ExampleCase{}
}

func (inst* p4876c9c6d6_unittestcases_ExampleCase) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p4876c9c6d.ExampleCase)
	nop(ie, com)

	


    return nil
}


