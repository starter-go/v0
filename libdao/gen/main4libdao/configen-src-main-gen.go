package main4libdao
import (
    p880ae68c3 "github.com/starter-go/v0/libdao/src/main/golang/libdaocom"
     "github.com/starter-go/application"
)

// type p880ae68c3.Foo in package:github.com/starter-go/v0/libdao/src/main/golang/libdaocom
//
// id:com-880ae68c313107a3-libdaocom-Foo
// class:
// alias:
// scope:singleton
//
type p880ae68c31_libdaocom_Foo struct {
}

func (inst* p880ae68c31_libdaocom_Foo) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-880ae68c313107a3-libdaocom-Foo"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p880ae68c31_libdaocom_Foo) new() any {
    return &p880ae68c3.Foo{}
}

func (inst* p880ae68c31_libdaocom_Foo) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p880ae68c3.Foo)
	nop(ie, com)

	


    return nil
}


