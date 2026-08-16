package test4libdao

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

    
    inst.register(&p8d3a42c23b_testcom_Example4t{})
    inst.register(&p8d3a42c23b_testcom_MockDao1{})
    inst.register(&p8d3a42c23b_testcom_MockDao2{})
    inst.register(&p8d3a42c23b_testcom_MockDao3{})
    inst.register(&p8d3a42c23b_testcom_MockDaoMain{})
    inst.register(&p8d3a42c23b_testcom_UnitForDaoProxy{})


    return nil
}
