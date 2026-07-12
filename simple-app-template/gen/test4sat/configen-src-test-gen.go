package test4sat
import (
    p5df72d263 "github.com/starter-go/v0/simple-app-template/src/test/golang/testcom"
     "github.com/starter-go/application"
)

// type p5df72d263.Example4t in package:github.com/starter-go/v0/simple-app-template/src/test/golang/testcom
//
// id:com-5df72d2631b7f514-testcom-Example4t
// class:
// alias:
// scope:singleton
//
type p5df72d2631_testcom_Example4t struct {
}

func (inst* p5df72d2631_testcom_Example4t) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-5df72d2631b7f514-testcom-Example4t"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p5df72d2631_testcom_Example4t) new() any {
    return &p5df72d263.Example4t{}
}

func (inst* p5df72d2631_testcom_Example4t) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p5df72d263.Example4t)
	nop(ie, com)

	


    return nil
}


