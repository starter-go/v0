package main4rbacwa

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

    
    inst.register(&p0fcabc88b9_debug_MockController{})
    inst.register(&p3821d5b0f6_itokens_JWTokenServiceImpl{})
    inst.register(&p3821d5b0f6_itokens_TokenServiceImpl{})
    inst.register(&p383a5f3ee9_iauthx_AuthenticationServiceImpl{})
    inst.register(&p383a5f3ee9_iauthx_AuthorizationServiceImpl{})
    inst.register(&p383a5f3ee9_iauthx_AuthxServiceImpl{})
    inst.register(&p4b670e5fac_iusers_UserDaoImpl{})
    inst.register(&p4b670e5fac_iusers_UserServiceImpl{})
    inst.register(&p7249a1596a_controllers_ExampleController{})
    inst.register(&p79b61f17f7_home_AuthxController{})
    inst.register(&p79b61f17f7_home_JWTokenController{})
    inst.register(&p7f79a0bbff_admin_UsersController{})
    inst.register(&p8841b23836_isessions_SessionDaoImpl{})
    inst.register(&p8841b23836_isessions_SessionServiceImpl{})
    inst.register(&pabcfb75207_idatabase_MyDatabaseAgentImpl{})
    inst.register(&pf75fd20ca4_database_MyDataGroup{})


    return nil
}
