package test4playground
import (
    p8fa0d9f5f "github.com/starter-go/v0/playground/src/test/golang/testcom"
     "github.com/starter-go/application"
)

// type p8fa0d9f5f.Example4t in package:github.com/starter-go/v0/playground/src/test/golang/testcom
//
// id:com-8fa0d9f5ff34c34a-testcom-Example4t
// class:
// alias:
// scope:singleton
//
type p8fa0d9f5ff_testcom_Example4t struct {
}

func (inst* p8fa0d9f5ff_testcom_Example4t) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-8fa0d9f5ff34c34a-testcom-Example4t"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p8fa0d9f5ff_testcom_Example4t) new() any {
    return &p8fa0d9f5f.Example4t{}
}

func (inst* p8fa0d9f5ff_testcom_Example4t) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p8fa0d9f5f.Example4t)
	nop(ie, com)

	


    return nil
}



// type p8fa0d9f5f.TryInterfaceReflect in package:github.com/starter-go/v0/playground/src/test/golang/testcom
//
// id:com-8fa0d9f5ff34c34a-testcom-TryInterfaceReflect
// class:
// alias:
// scope:singleton
//
type p8fa0d9f5ff_testcom_TryInterfaceReflect struct {
}

func (inst* p8fa0d9f5ff_testcom_TryInterfaceReflect) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-8fa0d9f5ff34c34a-testcom-TryInterfaceReflect"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p8fa0d9f5ff_testcom_TryInterfaceReflect) new() any {
    return &p8fa0d9f5f.TryInterfaceReflect{}
}

func (inst* p8fa0d9f5ff_testcom_TryInterfaceReflect) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p8fa0d9f5f.TryInterfaceReflect)
	nop(ie, com)

	


    return nil
}


