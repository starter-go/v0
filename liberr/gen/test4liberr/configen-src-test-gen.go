package test4liberr
import (
    pc12bc7456 "github.com/starter-go/v0/liberr/src/test/golang/pack1"
    pb412d31c6 "github.com/starter-go/v0/liberr/src/test/golang/pack2"
     "github.com/starter-go/application"
)

// type pc12bc7456.Demo1ErrorSet in package:github.com/starter-go/v0/liberr/src/test/golang/pack1
//
// id:com-c12bc74566deeab5-pack1-Demo1ErrorSet
// class:class-430c7b94c7c73bc48c254d5d9d9f6a6c-ErrorSet
// alias:
// scope:singleton
//
type pc12bc74566_pack1_Demo1ErrorSet struct {
}

func (inst* pc12bc74566_pack1_Demo1ErrorSet) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-c12bc74566deeab5-pack1-Demo1ErrorSet"
	r.Classes = "class-430c7b94c7c73bc48c254d5d9d9f6a6c-ErrorSet"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pc12bc74566_pack1_Demo1ErrorSet) new() any {
    return &pc12bc7456.Demo1ErrorSet{}
}

func (inst* pc12bc74566_pack1_Demo1ErrorSet) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pc12bc7456.Demo1ErrorSet)
	nop(ie, com)

	


    return nil
}



// type pb412d31c6.Demo2ErrorSet in package:github.com/starter-go/v0/liberr/src/test/golang/pack2
//
// id:com-b412d31c6d2a3ad3-pack2-Demo2ErrorSet
// class:class-430c7b94c7c73bc48c254d5d9d9f6a6c-ErrorSet
// alias:
// scope:singleton
//
type pb412d31c6d_pack2_Demo2ErrorSet struct {
}

func (inst* pb412d31c6d_pack2_Demo2ErrorSet) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-b412d31c6d2a3ad3-pack2-Demo2ErrorSet"
	r.Classes = "class-430c7b94c7c73bc48c254d5d9d9f6a6c-ErrorSet"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pb412d31c6d_pack2_Demo2ErrorSet) new() any {
    return &pb412d31c6.Demo2ErrorSet{}
}

func (inst* pb412d31c6d_pack2_Demo2ErrorSet) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pb412d31c6.Demo2ErrorSet)
	nop(ie, com)

	


    return nil
}



// type pb412d31c6.Demo2units in package:github.com/starter-go/v0/liberr/src/test/golang/pack2
//
// id:com-b412d31c6d2a3ad3-pack2-Demo2units
// class:
// alias:
// scope:singleton
//
type pb412d31c6d_pack2_Demo2units struct {
}

func (inst* pb412d31c6d_pack2_Demo2units) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-b412d31c6d2a3ad3-pack2-Demo2units"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pb412d31c6d_pack2_Demo2units) new() any {
    return &pb412d31c6.Demo2units{}
}

func (inst* pb412d31c6d_pack2_Demo2units) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pb412d31c6.Demo2units)
	nop(ie, com)

	


    return nil
}


