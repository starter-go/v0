package main4wat
import (
    pd1a916a20 "github.com/starter-go/libgin"
    p3456572df "github.com/starter-go/v0/web-app-template/app/web/controllers"
     "github.com/starter-go/application"
)

// type p3456572df.ExampleController in package:github.com/starter-go/v0/web-app-template/app/web/controllers
//
// id:com-3456572df3d3078b-controllers-ExampleController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p3456572df3_controllers_ExampleController struct {
}

func (inst* p3456572df3_controllers_ExampleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-3456572df3d3078b-controllers-ExampleController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p3456572df3_controllers_ExampleController) new() any {
    return &p3456572df.ExampleController{}
}

func (inst* p3456572df3_controllers_ExampleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p3456572df.ExampleController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)


    return nil
}


func (inst*p3456572df3_controllers_ExampleController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


