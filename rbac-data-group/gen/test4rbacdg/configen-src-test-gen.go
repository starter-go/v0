package test4rbacdg
import (
    p982dc5d72 "github.com/starter-go/v0/rbac-data-group/src/test/golang/testcom"
     "github.com/starter-go/application"
)

// type p982dc5d72.Example4t in package:github.com/starter-go/v0/rbac-data-group/src/test/golang/testcom
//
// id:com-982dc5d72ece47b7-testcom-Example4t
// class:
// alias:
// scope:singleton
//
type p982dc5d72e_testcom_Example4t struct {
}

func (inst* p982dc5d72e_testcom_Example4t) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-982dc5d72ece47b7-testcom-Example4t"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p982dc5d72e_testcom_Example4t) new() any {
    return &p982dc5d72.Example4t{}
}

func (inst* p982dc5d72e_testcom_Example4t) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p982dc5d72.Example4t)
	nop(ie, com)

	


    return nil
}


