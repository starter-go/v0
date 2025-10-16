package test4wat
import (
    p891b3af5b "github.com/starter-go/v0/web-app-template/src/test/golang/unittestcases"
     "github.com/starter-go/application"
)

// type p891b3af5b.ExampleCase in package:github.com/starter-go/v0/web-app-template/src/test/golang/unittestcases
//
// id:com-891b3af5b83599e0-unittestcases-ExampleCase
// class:
// alias:
// scope:singleton
//
type p891b3af5b8_unittestcases_ExampleCase struct {
}

func (inst* p891b3af5b8_unittestcases_ExampleCase) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-891b3af5b83599e0-unittestcases-ExampleCase"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p891b3af5b8_unittestcases_ExampleCase) new() any {
    return &p891b3af5b.ExampleCase{}
}

func (inst* p891b3af5b8_unittestcases_ExampleCase) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p891b3af5b.ExampleCase)
	nop(ie, com)

	


    return nil
}


