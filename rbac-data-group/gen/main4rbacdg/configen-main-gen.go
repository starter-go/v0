package main4rbacdg

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

    
    inst.register(&p01fd859558_isessions_SessionDaoImpl{})
    inst.register(&p3da2710ff4_idatagroup_StdRbacDataAgent{})
    inst.register(&p3da2710ff4_idatagroup_StdRbacDataGroup{})
    inst.register(&p651d619428_iroles_RoleDaoImpl{})
    inst.register(&pa2c28bd102_iusers_UserDaoImpl{})
    inst.register(&pb437cfd580_itables_TableDaoImpl{})
    inst.register(&pe64067b08c_ipermissions_PermissionDaoImpl{})
    inst.register(&pfa85f3231c_iauths_AuthenticationDaoImpl{})


    return nil
}
