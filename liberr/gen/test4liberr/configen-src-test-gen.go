package test4liberr
import (
    pa4d825a70 "github.com/starter-go/v0/liberr/src/test/golang/unittestcases"
     "github.com/starter-go/application"
)

// type pa4d825a70.ExampleCase in package:github.com/starter-go/v0/liberr/src/test/golang/unittestcases
//
// id:com-a4d825a70bc96e1c-unittestcases-ExampleCase
// class:
// alias:
// scope:singleton
//
type pa4d825a70b_unittestcases_ExampleCase struct {
}

func (inst* pa4d825a70b_unittestcases_ExampleCase) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a4d825a70bc96e1c-unittestcases-ExampleCase"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa4d825a70b_unittestcases_ExampleCase) new() any {
    return &pa4d825a70.ExampleCase{}
}

func (inst* pa4d825a70b_unittestcases_ExampleCase) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa4d825a70.ExampleCase)
	nop(ie, com)

	


    return nil
}


