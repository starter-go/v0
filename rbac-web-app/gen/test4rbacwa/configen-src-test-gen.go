package test4rbacwa
import (
    pfabbd91b2 "github.com/starter-go/v0/rbac-web-app/src/test/golang/unittestcases"
     "github.com/starter-go/application"
)

// type pfabbd91b2.ExampleCase in package:github.com/starter-go/v0/rbac-web-app/src/test/golang/unittestcases
//
// id:com-fabbd91b2d06fe21-unittestcases-ExampleCase
// class:
// alias:
// scope:singleton
//
type pfabbd91b2d_unittestcases_ExampleCase struct {
}

func (inst* pfabbd91b2d_unittestcases_ExampleCase) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-fabbd91b2d06fe21-unittestcases-ExampleCase"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pfabbd91b2d_unittestcases_ExampleCase) new() any {
    return &pfabbd91b2.ExampleCase{}
}

func (inst* pfabbd91b2d_unittestcases_ExampleCase) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pfabbd91b2.ExampleCase)
	nop(ie, com)

	


    return nil
}


