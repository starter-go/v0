package test4wat
import (
    pb7cda72c4 "github.com/starter-go/v0/web-app-template/src/test/golang/unitcom"
     "github.com/starter-go/application"
)

// type pb7cda72c4.Example1 in package:github.com/starter-go/v0/web-app-template/src/test/golang/unitcom
//
// id:com-b7cda72c4dc1987a-unitcom-Example1
// class:
// alias:
// scope:singleton
//
type pb7cda72c4d_unitcom_Example1 struct {
}

func (inst* pb7cda72c4d_unitcom_Example1) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-b7cda72c4dc1987a-unitcom-Example1"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pb7cda72c4d_unitcom_Example1) new() any {
    return &pb7cda72c4.Example1{}
}

func (inst* pb7cda72c4d_unitcom_Example1) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pb7cda72c4.Example1)
	nop(ie, com)

	


    return nil
}


