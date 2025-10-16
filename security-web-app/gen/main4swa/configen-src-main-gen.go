package main4swa

import (
	"github.com/starter-go/application"
	p45c9f8c2d "github.com/starter-go/v0/security-web-app/app/web/controllers"
)

// type p45c9f8c2d.ExampleController in package:github.com/starter-go/v0/security-web-app/app/web/controllers
//
// id:com-45c9f8c2d5988003-controllers-ExampleController
// class:
// alias:
// scope:singleton
type p45c9f8c2d5_controllers_ExampleController struct {
}

func (inst *p45c9f8c2d5_controllers_ExampleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-45c9f8c2d5988003-controllers-ExampleController"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst *p45c9f8c2d5_controllers_ExampleController) new() any {
	return &p45c9f8c2d.ExampleController{}
}

func (inst *p45c9f8c2d5_controllers_ExampleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p45c9f8c2d.ExampleController)
	nop(ie, com)

	return nil
}
