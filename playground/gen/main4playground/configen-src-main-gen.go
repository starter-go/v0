package main4playground
import (
    pcd7c597fc "github.com/starter-go/v0/playground/src/main/golang/app/demo"
     "github.com/starter-go/application"
)

// type pcd7c597fc.ExampleController in package:github.com/starter-go/v0/playground/src/main/golang/app/demo
//
// id:com-cd7c597fc1e18cf7-demo-ExampleController
// class:
// alias:
// scope:singleton
//
type pcd7c597fc1_demo_ExampleController struct {
}

func (inst* pcd7c597fc1_demo_ExampleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-cd7c597fc1e18cf7-demo-ExampleController"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pcd7c597fc1_demo_ExampleController) new() any {
    return &pcd7c597fc.ExampleController{}
}

func (inst* pcd7c597fc1_demo_ExampleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pcd7c597fc.ExampleController)
	nop(ie, com)

	


    return nil
}


