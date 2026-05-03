package main4subjects

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

    
    inst.register(&p0d8344d2dc_isessions_SessionDaoImpl{})
    inst.register(&p0d8344d2dc_isessions_SessionServiceImpl{})
    inst.register(&p1e478b5d42_filter4db_Filter4db{})
    inst.register(&p29a9720d88_isubjectdb_SubjectDataBaseAgent{})
    inst.register(&p29a9720d88_isubjectdb_SubjectDataGroup{})
    inst.register(&p4270aadd3a_filter4jwt_Filter4jwt{})
    inst.register(&pd00479541f_filter4log_Filter4Log{})
    inst.register(&pd3bd430b1e_filter4cache_Filter4Cache1{})
    inst.register(&pd3bd430b1e_filter4cache_Filter4Cache2{})
    inst.register(&pdb4050fdbd_autoloaders_DefaultFilterChainHolder{})
    inst.register(&pdb4050fdbd_autoloaders_DefaultFilterChainLoader{})


    return nil
}
