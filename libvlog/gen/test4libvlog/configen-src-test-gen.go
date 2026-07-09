package test4libvlog
import (
    p17727fc70 "github.com/starter-go/v0/libvlog/src/test/golang/testcom"
     "github.com/starter-go/application"
)

// type p17727fc70.Example1 in package:github.com/starter-go/v0/libvlog/src/test/golang/testcom
//
// id:com-17727fc7053d8691-testcom-Example1
// class:class-0dc072ed44b3563882bff4e657a52e62-Unit
// alias:
// scope:singleton
//
type p17727fc705_testcom_Example1 struct {
}

func (inst* p17727fc705_testcom_Example1) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-17727fc7053d8691-testcom-Example1"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Unit"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p17727fc705_testcom_Example1) new() any {
    return &p17727fc70.Example1{}
}

func (inst* p17727fc705_testcom_Example1) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p17727fc70.Example1)
	nop(ie, com)

	


    return nil
}



// type p17727fc70.TryLogLevels in package:github.com/starter-go/v0/libvlog/src/test/golang/testcom
//
// id:com-17727fc7053d8691-testcom-TryLogLevels
// class:class-0dc072ed44b3563882bff4e657a52e62-Unit
// alias:
// scope:singleton
//
type p17727fc705_testcom_TryLogLevels struct {
}

func (inst* p17727fc705_testcom_TryLogLevels) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-17727fc7053d8691-testcom-TryLogLevels"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Unit"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p17727fc705_testcom_TryLogLevels) new() any {
    return &p17727fc70.TryLogLevels{}
}

func (inst* p17727fc705_testcom_TryLogLevels) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p17727fc70.TryLogLevels)
	nop(ie, com)

	


    return nil
}


