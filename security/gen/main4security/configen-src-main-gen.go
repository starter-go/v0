package main4security
import (
    p030415bff "github.com/starter-go/v0/security/app/web/controllers"
     "github.com/starter-go/application"
)

// type p030415bff.ExampleController in package:github.com/starter-go/v0/security/app/web/controllers
//
// id:com-030415bffd75d59f-controllers-ExampleController
// class:
// alias:
// scope:singleton
//
type p030415bffd_controllers_ExampleController struct {
}

func (inst* p030415bffd_controllers_ExampleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-030415bffd75d59f-controllers-ExampleController"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p030415bffd_controllers_ExampleController) new() any {
    return &p030415bff.ExampleController{}
}

func (inst* p030415bffd_controllers_ExampleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p030415bff.ExampleController)
	nop(ie, com)

	


    return nil
}


