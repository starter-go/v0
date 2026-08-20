package main4avatars
import (
    pd1a916a20 "github.com/starter-go/libgin"
    p886fa7078 "github.com/starter-go/v0/avatars/app/web/controllers"
     "github.com/starter-go/application"
)

// type p886fa7078.ExampleController in package:github.com/starter-go/v0/avatars/app/web/controllers
//
// id:com-886fa70781d088a8-controllers-ExampleController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p886fa70781_controllers_ExampleController struct {
}

func (inst* p886fa70781_controllers_ExampleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-886fa70781d088a8-controllers-ExampleController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p886fa70781_controllers_ExampleController) new() any {
    return &p886fa7078.ExampleController{}
}

func (inst* p886fa70781_controllers_ExampleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p886fa7078.ExampleController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)


    return nil
}


func (inst*p886fa70781_controllers_ExampleController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


