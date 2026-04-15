package test4httest
import (
    p1da300815 "github.com/starter-go/v0/httest/src/test/golang/unittestcases"
     "github.com/starter-go/application"
)

// type p1da300815.ExampleCase in package:github.com/starter-go/v0/httest/src/test/golang/unittestcases
//
// id:com-1da30081568d0f45-unittestcases-ExampleCase
// class:
// alias:
// scope:singleton
//
type p1da3008156_unittestcases_ExampleCase struct {
}

func (inst* p1da3008156_unittestcases_ExampleCase) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1da30081568d0f45-unittestcases-ExampleCase"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1da3008156_unittestcases_ExampleCase) new() any {
    return &p1da300815.ExampleCase{}
}

func (inst* p1da3008156_unittestcases_ExampleCase) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1da300815.ExampleCase)
	nop(ie, com)

	


    return nil
}


