package testwebclient4rbacwa

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

    
    inst.register(&p48be8f24a5_cases_CaseTryExample{})
    inst.register(&p48be8f24a5_cases_CaseTrySubjects{})
    inst.register(&p48be8f24a5_cases_CaseTryUseAuth{})
    inst.register(&p48be8f24a5_cases_CaseTryUseSessions{})


    return nil
}
