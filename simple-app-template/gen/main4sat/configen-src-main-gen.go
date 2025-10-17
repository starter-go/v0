package main4sat
import (
    pde88e5193 "github.com/starter-go/v0/simple-app-template/app/demo"
     "github.com/starter-go/application"
)

// type pde88e5193.ExampleController in package:github.com/starter-go/v0/simple-app-template/app/demo
//
// id:com-de88e5193f842f88-demo-ExampleController
// class:
// alias:
// scope:singleton
//
type pde88e5193f_demo_ExampleController struct {
}

func (inst* pde88e5193f_demo_ExampleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-de88e5193f842f88-demo-ExampleController"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pde88e5193f_demo_ExampleController) new() any {
    return &pde88e5193.ExampleController{}
}

func (inst* pde88e5193f_demo_ExampleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pde88e5193.ExampleController)
	nop(ie, com)

	


    return nil
}


