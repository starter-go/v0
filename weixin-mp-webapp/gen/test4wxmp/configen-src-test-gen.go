package test4wxmp
import (
    pde0b19f7d "github.com/starter-go/v0/weixin-mp-webapp/src/test/golang/unittestcases"
     "github.com/starter-go/application"
)

// type pde0b19f7d.ExampleCase in package:github.com/starter-go/v0/weixin-mp-webapp/src/test/golang/unittestcases
//
// id:com-de0b19f7debc9066-unittestcases-ExampleCase
// class:
// alias:
// scope:singleton
//
type pde0b19f7de_unittestcases_ExampleCase struct {
}

func (inst* pde0b19f7de_unittestcases_ExampleCase) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-de0b19f7debc9066-unittestcases-ExampleCase"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pde0b19f7de_unittestcases_ExampleCase) new() any {
    return &pde0b19f7d.ExampleCase{}
}

func (inst* pde0b19f7de_unittestcases_ExampleCase) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pde0b19f7d.ExampleCase)
	nop(ie, com)

	


    return nil
}


