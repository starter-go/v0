package test4sat
import (
    p7acdadc5b "github.com/starter-go/v0/simple-app-template/src/test/golang/unittestcases"
     "github.com/starter-go/application"
)

// type p7acdadc5b.ExampleCase in package:github.com/starter-go/v0/simple-app-template/src/test/golang/unittestcases
//
// id:com-7acdadc5ba06833f-unittestcases-ExampleCase
// class:
// alias:
// scope:singleton
//
type p7acdadc5ba_unittestcases_ExampleCase struct {
}

func (inst* p7acdadc5ba_unittestcases_ExampleCase) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-7acdadc5ba06833f-unittestcases-ExampleCase"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p7acdadc5ba_unittestcases_ExampleCase) new() any {
    return &p7acdadc5b.ExampleCase{}
}

func (inst* p7acdadc5ba_unittestcases_ExampleCase) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p7acdadc5b.ExampleCase)
	nop(ie, com)

	


    return nil
}


