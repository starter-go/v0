package test4rbacwa
import (
    pb51fe7a59 "github.com/starter-go/v0/rbac-web-app/src/test/golang/testcase"
     "github.com/starter-go/application"
)

// type pb51fe7a59.CaseTryExample in package:github.com/starter-go/v0/rbac-web-app/src/test/golang/testcase
//
// id:com-b51fe7a597a81919-testcase-CaseTryExample
// class:
// alias:
// scope:singleton
//
type pb51fe7a597_testcase_CaseTryExample struct {
}

func (inst* pb51fe7a597_testcase_CaseTryExample) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-b51fe7a597a81919-testcase-CaseTryExample"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pb51fe7a597_testcase_CaseTryExample) new() any {
    return &pb51fe7a59.CaseTryExample{}
}

func (inst* pb51fe7a597_testcase_CaseTryExample) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pb51fe7a59.CaseTryExample)
	nop(ie, com)

	


    return nil
}


