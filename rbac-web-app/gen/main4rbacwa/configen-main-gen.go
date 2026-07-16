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

    
    inst.register(&p383a5f3ee9_iauthx_ActionLoginAuthorizer{})
    inst.register(&p383a5f3ee9_iauthx_AuthenticationServiceImpl{})
    inst.register(&p383a5f3ee9_iauthx_AuthorizationServiceImpl{})
    inst.register(&p383a5f3ee9_iauthx_AuthxServiceImpl{})
    inst.register(&p383a5f3ee9_iauthx_HTTPBasicAuthenticator{})
    inst.register(&p383a5f3ee9_iauthx_UserAuthNameServiceImpl{})
    inst.register(&p4b670e5fac_iusers_UserDaoImpl{})
    inst.register(&p4b670e5fac_iusers_UserServiceImpl{})
    inst.register(&p590373c3a2_iexamples_ExampleDaoImpl{})
    inst.register(&p590373c3a2_iexamples_ExampleServiceImpl{})
    inst.register(&p7249a1596a_controllers_ExampleController{})
    inst.register(&p7f79a0bbff_admin_AdminPermissionController{})
    inst.register(&p7f79a0bbff_admin_AdminRoleController{})
    inst.register(&p7f79a0bbff_admin_AdminSessionController{})
    inst.register(&p7f79a0bbff_admin_ExampleController{})
    inst.register(&p7f79a0bbff_admin_UsersController{})
    inst.register(&p95d8708df4_iroles_RoleDaoImpl{})
    inst.register(&p95d8708df4_iroles_RoleServiceImpl{})
    inst.register(&pabcfb75207_idatabase_MyDatabaseAgentImpl{})
    inst.register(&pc3ca883e25_helper_GinLibjwtAdapter{})
    inst.register(&pc3ca883e25_helper_JWTokenFilter{})
    inst.register(&pc3ca883e25_helper_RbacCheckerFilter{})
    inst.register(&pe475f399ce_my_AuthxController{})
    inst.register(&pe475f399ce_my_SessionController{})
    inst.register(&pea0a25a6a9_ipermissions_PermissionDaoImpl{})
    inst.register(&pea0a25a6a9_ipermissions_PermissionServiceImpl{})
    inst.register(&pf75fd20ca4_database_MyDataGroup{})


    return nil
}
