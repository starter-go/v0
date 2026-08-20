package test4avatars
import (
    p48bcc947c "github.com/starter-go/v0/avatars/src/test/golang/unitcom"
     "github.com/starter-go/application"
)

// type p48bcc947c.Example1 in package:github.com/starter-go/v0/avatars/src/test/golang/unitcom
//
// id:com-48bcc947c98c904b-unitcom-Example1
// class:
// alias:
// scope:singleton
//
type p48bcc947c9_unitcom_Example1 struct {
}

func (inst* p48bcc947c9_unitcom_Example1) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-48bcc947c98c904b-unitcom-Example1"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p48bcc947c9_unitcom_Example1) new() any {
    return &p48bcc947c.Example1{}
}

func (inst* p48bcc947c9_unitcom_Example1) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p48bcc947c.Example1)
	nop(ie, com)

	


    return nil
}


