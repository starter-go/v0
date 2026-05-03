package test4rbacwa

import "github.com/starter-go/application"

func nop(a ... any) {    
}

func registerComponents(cr application.ComponentRegistry) error {
    ac:=&autoRegistrar{}
    ac.init(cr)
    return ac.addAll()
}

type comFactory interface {
    register(cr application.ComponentRegistry) error
}

type autoRegistrar struct {
    cr application.ComponentRegistry
}

func (inst *autoRegistrar) init(cr application.ComponentRegistry) {
	inst.cr = cr
}

func (inst *autoRegistrar) register(factory comFactory) error {
	return factory.register(inst.cr)
}

func (inst*autoRegistrar) addAll() error {

    
    inst.register(&p78d6a2ce14_wuclient_WebUnitTestSessionRunner{})
    inst.register(&p90f4aa59a9_wuserver_WebUnitTestSessionController{})
    inst.register(&pfabbd91b2d_unittestcases_CaseTrySubject{})
    inst.register(&pfabbd91b2d_unittestcases_ExampleCase{})


    return nil
}
