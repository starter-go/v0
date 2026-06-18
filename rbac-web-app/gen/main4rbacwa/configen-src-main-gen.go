package main4rbacwa
import (
    pd1a916a20 "github.com/starter-go/libgin"
    p512a30914 "github.com/starter-go/libgorm"
    p9621e8b71 "github.com/starter-go/security/random"
    p84e86b31b "github.com/starter-go/v0/rbac-web-app/app/classes/authx"
    p2d258e798 "github.com/starter-go/v0/rbac-web-app/app/classes/users"
    pf75fd20ca "github.com/starter-go/v0/rbac-web-app/app/data/database"
    p383a5f3ee "github.com/starter-go/v0/rbac-web-app/app/implements/iauthx"
    pabcfb7520 "github.com/starter-go/v0/rbac-web-app/app/implements/idatabase"
    p4b670e5fa "github.com/starter-go/v0/rbac-web-app/app/implements/iusers"
    p7249a1596 "github.com/starter-go/v0/rbac-web-app/app/web/controllers"
    p7f79a0bbf "github.com/starter-go/v0/rbac-web-app/app/web/controllers/admin"
    pc3ca883e2 "github.com/starter-go/v0/rbac-web-app/app/web/controllers/helper"
    p79b61f17f "github.com/starter-go/v0/rbac-web-app/app/web/controllers/home"
    pfd2c28477 "github.com/starter-go/v0/subjects"
     "github.com/starter-go/application"
)

// type pf75fd20ca.MyDataGroup in package:github.com/starter-go/v0/rbac-web-app/app/data/database
//
// id:com-f75fd20ca4f14f60-database-MyDataGroup
// class:class-512a309140d0ad99eb1c95c8dc0d02f9-GroupRegistry
// alias:
// scope:singleton
//
type pf75fd20ca4_database_MyDataGroup struct {
}

func (inst* pf75fd20ca4_database_MyDataGroup) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-f75fd20ca4f14f60-database-MyDataGroup"
	r.Classes = "class-512a309140d0ad99eb1c95c8dc0d02f9-GroupRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pf75fd20ca4_database_MyDataGroup) new() any {
    return &pf75fd20ca.MyDataGroup{}
}

func (inst* pf75fd20ca4_database_MyDataGroup) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pf75fd20ca.MyDataGroup)
	nop(ie, com)

	
    com.Enabled = inst.getEnabled(ie)
    com.Alias = inst.getAlias(ie)
    com.Prefix = inst.getPrefix(ie)
    com.Source = inst.getSource(ie)
    com.URI = inst.getURI(ie)


    return nil
}


func (inst*pf75fd20ca4_database_MyDataGroup) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${datagroup.default.enabled}")
}


func (inst*pf75fd20ca4_database_MyDataGroup) getAlias(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.default.alias}")
}


func (inst*pf75fd20ca4_database_MyDataGroup) getPrefix(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.default.table-name-prefix}")
}


func (inst*pf75fd20ca4_database_MyDataGroup) getSource(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.default.datasource}")
}


func (inst*pf75fd20ca4_database_MyDataGroup) getURI(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.default.uri}")
}



// type p383a5f3ee.ActionLoginAuthorizer in package:github.com/starter-go/v0/rbac-web-app/app/implements/iauthx
//
// id:com-383a5f3ee9f70032-iauthx-ActionLoginAuthorizer
// class:class-84e86b31b37511e994be50a820fadd04-Authorizer
// alias:
// scope:singleton
//
type p383a5f3ee9_iauthx_ActionLoginAuthorizer struct {
}

func (inst* p383a5f3ee9_iauthx_ActionLoginAuthorizer) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-383a5f3ee9f70032-iauthx-ActionLoginAuthorizer"
	r.Classes = "class-84e86b31b37511e994be50a820fadd04-Authorizer"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p383a5f3ee9_iauthx_ActionLoginAuthorizer) new() any {
    return &p383a5f3ee.ActionLoginAuthorizer{}
}

func (inst* p383a5f3ee9_iauthx_ActionLoginAuthorizer) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p383a5f3ee.ActionLoginAuthorizer)
	nop(ie, com)

	
    com.SessionSpanSecondsDef = inst.getSessionSpanSecondsDef(ie)
    com.SessionSpanSecondsMax = inst.getSessionSpanSecondsMax(ie)
    com.SessionSpanSecondsMin = inst.getSessionSpanSecondsMin(ie)


    return nil
}


func (inst*p383a5f3ee9_iauthx_ActionLoginAuthorizer) getSessionSpanSecondsDef(ie application.InjectionExt)int64{
    return ie.GetInt64("${security.session.default-age-sec}")
}


func (inst*p383a5f3ee9_iauthx_ActionLoginAuthorizer) getSessionSpanSecondsMax(ie application.InjectionExt)int64{
    return ie.GetInt64("${security.session.max-age-sec}")
}


func (inst*p383a5f3ee9_iauthx_ActionLoginAuthorizer) getSessionSpanSecondsMin(ie application.InjectionExt)int64{
    return ie.GetInt64("${security.session.min-age-sec}")
}



// type p383a5f3ee.AuthenticationServiceImpl in package:github.com/starter-go/v0/rbac-web-app/app/implements/iauthx
//
// id:com-383a5f3ee9f70032-iauthx-AuthenticationServiceImpl
// class:
// alias:alias-84e86b31b37511e994be50a820fadd04-AuthenticationService
// scope:singleton
//
type p383a5f3ee9_iauthx_AuthenticationServiceImpl struct {
}

func (inst* p383a5f3ee9_iauthx_AuthenticationServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-383a5f3ee9f70032-iauthx-AuthenticationServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-84e86b31b37511e994be50a820fadd04-AuthenticationService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p383a5f3ee9_iauthx_AuthenticationServiceImpl) new() any {
    return &p383a5f3ee.AuthenticationServiceImpl{}
}

func (inst* p383a5f3ee9_iauthx_AuthenticationServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p383a5f3ee.AuthenticationServiceImpl)
	nop(ie, com)

	
    com.RawProviderList = inst.getRawProviderList(ie)


    return nil
}


func (inst*p383a5f3ee9_iauthx_AuthenticationServiceImpl) getRawProviderList(ie application.InjectionExt)[]p84e86b31b.Authenticator{
    dst := make([]p84e86b31b.Authenticator, 0)
    src := ie.ListComponents(".class-84e86b31b37511e994be50a820fadd04-Authenticator")
    for _, item1 := range src {
        item2 := item1.(p84e86b31b.Authenticator)
        dst = append(dst, item2)
    }
    return dst
}



// type p383a5f3ee.AuthorizationServiceImpl in package:github.com/starter-go/v0/rbac-web-app/app/implements/iauthx
//
// id:com-383a5f3ee9f70032-iauthx-AuthorizationServiceImpl
// class:
// alias:alias-84e86b31b37511e994be50a820fadd04-AuthorizationService
// scope:singleton
//
type p383a5f3ee9_iauthx_AuthorizationServiceImpl struct {
}

func (inst* p383a5f3ee9_iauthx_AuthorizationServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-383a5f3ee9f70032-iauthx-AuthorizationServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-84e86b31b37511e994be50a820fadd04-AuthorizationService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p383a5f3ee9_iauthx_AuthorizationServiceImpl) new() any {
    return &p383a5f3ee.AuthorizationServiceImpl{}
}

func (inst* p383a5f3ee9_iauthx_AuthorizationServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p383a5f3ee.AuthorizationServiceImpl)
	nop(ie, com)

	
    com.RawProviderList = inst.getRawProviderList(ie)


    return nil
}


func (inst*p383a5f3ee9_iauthx_AuthorizationServiceImpl) getRawProviderList(ie application.InjectionExt)[]p84e86b31b.Authorizer{
    dst := make([]p84e86b31b.Authorizer, 0)
    src := ie.ListComponents(".class-84e86b31b37511e994be50a820fadd04-Authorizer")
    for _, item1 := range src {
        item2 := item1.(p84e86b31b.Authorizer)
        dst = append(dst, item2)
    }
    return dst
}



// type p383a5f3ee.AuthxServiceImpl in package:github.com/starter-go/v0/rbac-web-app/app/implements/iauthx
//
// id:com-383a5f3ee9f70032-iauthx-AuthxServiceImpl
// class:
// alias:alias-84e86b31b37511e994be50a820fadd04-Service
// scope:singleton
//
type p383a5f3ee9_iauthx_AuthxServiceImpl struct {
}

func (inst* p383a5f3ee9_iauthx_AuthxServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-383a5f3ee9f70032-iauthx-AuthxServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-84e86b31b37511e994be50a820fadd04-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p383a5f3ee9_iauthx_AuthxServiceImpl) new() any {
    return &p383a5f3ee.AuthxServiceImpl{}
}

func (inst* p383a5f3ee9_iauthx_AuthxServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p383a5f3ee.AuthxServiceImpl)
	nop(ie, com)

	
    com.Auth1 = inst.getAuth1(ie)
    com.Auth2 = inst.getAuth2(ie)


    return nil
}


func (inst*p383a5f3ee9_iauthx_AuthxServiceImpl) getAuth1(ie application.InjectionExt)p84e86b31b.AuthenticationService{
    return ie.GetComponent("#alias-84e86b31b37511e994be50a820fadd04-AuthenticationService").(p84e86b31b.AuthenticationService)
}


func (inst*p383a5f3ee9_iauthx_AuthxServiceImpl) getAuth2(ie application.InjectionExt)p84e86b31b.AuthorizationService{
    return ie.GetComponent("#alias-84e86b31b37511e994be50a820fadd04-AuthorizationService").(p84e86b31b.AuthorizationService)
}



// type p383a5f3ee.HTTPBasicAuthenticator in package:github.com/starter-go/v0/rbac-web-app/app/implements/iauthx
//
// id:com-383a5f3ee9f70032-iauthx-HTTPBasicAuthenticator
// class:class-84e86b31b37511e994be50a820fadd04-Authenticator
// alias:
// scope:singleton
//
type p383a5f3ee9_iauthx_HTTPBasicAuthenticator struct {
}

func (inst* p383a5f3ee9_iauthx_HTTPBasicAuthenticator) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-383a5f3ee9f70032-iauthx-HTTPBasicAuthenticator"
	r.Classes = "class-84e86b31b37511e994be50a820fadd04-Authenticator"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p383a5f3ee9_iauthx_HTTPBasicAuthenticator) new() any {
    return &p383a5f3ee.HTTPBasicAuthenticator{}
}

func (inst* p383a5f3ee9_iauthx_HTTPBasicAuthenticator) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p383a5f3ee.HTTPBasicAuthenticator)
	nop(ie, com)

	
    com.UserAuthNameService = inst.getUserAuthNameService(ie)
    com.DebugEnabled = inst.getDebugEnabled(ie)


    return nil
}


func (inst*p383a5f3ee9_iauthx_HTTPBasicAuthenticator) getUserAuthNameService(ie application.InjectionExt)p84e86b31b.UserAuthNameService{
    return ie.GetComponent("#alias-84e86b31b37511e994be50a820fadd04-UserAuthNameService").(p84e86b31b.UserAuthNameService)
}


func (inst*p383a5f3ee9_iauthx_HTTPBasicAuthenticator) getDebugEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${debug.enabled}")
}



// type p383a5f3ee.UserAuthNameServiceImpl in package:github.com/starter-go/v0/rbac-web-app/app/implements/iauthx
//
// id:com-383a5f3ee9f70032-iauthx-UserAuthNameServiceImpl
// class:
// alias:alias-84e86b31b37511e994be50a820fadd04-UserAuthNameService
// scope:singleton
//
type p383a5f3ee9_iauthx_UserAuthNameServiceImpl struct {
}

func (inst* p383a5f3ee9_iauthx_UserAuthNameServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-383a5f3ee9f70032-iauthx-UserAuthNameServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-84e86b31b37511e994be50a820fadd04-UserAuthNameService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p383a5f3ee9_iauthx_UserAuthNameServiceImpl) new() any {
    return &p383a5f3ee.UserAuthNameServiceImpl{}
}

func (inst* p383a5f3ee9_iauthx_UserAuthNameServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p383a5f3ee.UserAuthNameServiceImpl)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*p383a5f3ee9_iauthx_UserAuthNameServiceImpl) getDao(ie application.InjectionExt)p2d258e798.DAO{
    return ie.GetComponent("#alias-2d258e79845135c43979b78bd9f1f74e-DAO").(p2d258e798.DAO)
}



// type pabcfb7520.MyDatabaseAgentImpl in package:github.com/starter-go/v0/rbac-web-app/app/implements/idatabase
//
// id:com-abcfb752071056b5-idatabase-MyDatabaseAgentImpl
// class:
// alias:alias-f75fd20ca4f14f60caace00a3f949b98-Agent
// scope:singleton
//
type pabcfb75207_idatabase_MyDatabaseAgentImpl struct {
}

func (inst* pabcfb75207_idatabase_MyDatabaseAgentImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-abcfb752071056b5-idatabase-MyDatabaseAgentImpl"
	r.Classes = ""
	r.Aliases = "alias-f75fd20ca4f14f60caace00a3f949b98-Agent"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pabcfb75207_idatabase_MyDatabaseAgentImpl) new() any {
    return &pabcfb7520.MyDatabaseAgentImpl{}
}

func (inst* pabcfb75207_idatabase_MyDatabaseAgentImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pabcfb7520.MyDatabaseAgentImpl)
	nop(ie, com)

	
    com.DSManager = inst.getDSManager(ie)


    return nil
}


func (inst*pabcfb75207_idatabase_MyDatabaseAgentImpl) getDSManager(ie application.InjectionExt)p512a30914.DataSourceManager{
    return ie.GetComponent("#alias-512a309140d0ad99eb1c95c8dc0d02f9-DataSourceManager").(p512a30914.DataSourceManager)
}



// type p4b670e5fa.UserDaoImpl in package:github.com/starter-go/v0/rbac-web-app/app/implements/iusers
//
// id:com-4b670e5facdc70b8-iusers-UserDaoImpl
// class:
// alias:alias-2d258e79845135c43979b78bd9f1f74e-DAO
// scope:singleton
//
type p4b670e5fac_iusers_UserDaoImpl struct {
}

func (inst* p4b670e5fac_iusers_UserDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-4b670e5facdc70b8-iusers-UserDaoImpl"
	r.Classes = ""
	r.Aliases = "alias-2d258e79845135c43979b78bd9f1f74e-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p4b670e5fac_iusers_UserDaoImpl) new() any {
    return &p4b670e5fa.UserDaoImpl{}
}

func (inst* p4b670e5fac_iusers_UserDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p4b670e5fa.UserDaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDSer = inst.getUUIDSer(ie)


    return nil
}


func (inst*p4b670e5fac_iusers_UserDaoImpl) getAgent(ie application.InjectionExt)pf75fd20ca.Agent{
    return ie.GetComponent("#alias-f75fd20ca4f14f60caace00a3f949b98-Agent").(pf75fd20ca.Agent)
}


func (inst*p4b670e5fac_iusers_UserDaoImpl) getUUIDSer(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type p4b670e5fa.UserServiceImpl in package:github.com/starter-go/v0/rbac-web-app/app/implements/iusers
//
// id:com-4b670e5facdc70b8-iusers-UserServiceImpl
// class:
// alias:alias-2d258e79845135c43979b78bd9f1f74e-Service
// scope:singleton
//
type p4b670e5fac_iusers_UserServiceImpl struct {
}

func (inst* p4b670e5fac_iusers_UserServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-4b670e5facdc70b8-iusers-UserServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-2d258e79845135c43979b78bd9f1f74e-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p4b670e5fac_iusers_UserServiceImpl) new() any {
    return &p4b670e5fa.UserServiceImpl{}
}

func (inst* p4b670e5fac_iusers_UserServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p4b670e5fa.UserServiceImpl)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*p4b670e5fac_iusers_UserServiceImpl) getDao(ie application.InjectionExt)p2d258e798.DAO{
    return ie.GetComponent("#alias-2d258e79845135c43979b78bd9f1f74e-DAO").(p2d258e798.DAO)
}



// type p7249a1596.ExampleController in package:github.com/starter-go/v0/rbac-web-app/app/web/controllers
//
// id:com-7249a1596a817e5d-controllers-ExampleController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p7249a1596a_controllers_ExampleController struct {
}

func (inst* p7249a1596a_controllers_ExampleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-7249a1596a817e5d-controllers-ExampleController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p7249a1596a_controllers_ExampleController) new() any {
    return &p7249a1596.ExampleController{}
}

func (inst* p7249a1596a_controllers_ExampleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p7249a1596.ExampleController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p7249a1596a_controllers_ExampleController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p7249a1596a_controllers_ExampleController) getService(ie application.InjectionExt)p2d258e798.Service{
    return ie.GetComponent("#alias-2d258e79845135c43979b78bd9f1f74e-Service").(p2d258e798.Service)
}



// type p7f79a0bbf.UsersController in package:github.com/starter-go/v0/rbac-web-app/app/web/controllers/admin
//
// id:com-7f79a0bbff8317e1-admin-UsersController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p7f79a0bbff_admin_UsersController struct {
}

func (inst* p7f79a0bbff_admin_UsersController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-7f79a0bbff8317e1-admin-UsersController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p7f79a0bbff_admin_UsersController) new() any {
    return &p7f79a0bbf.UsersController{}
}

func (inst* p7f79a0bbff_admin_UsersController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p7f79a0bbf.UsersController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p7f79a0bbff_admin_UsersController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p7f79a0bbff_admin_UsersController) getService(ie application.InjectionExt)p2d258e798.Service{
    return ie.GetComponent("#alias-2d258e79845135c43979b78bd9f1f74e-Service").(p2d258e798.Service)
}



// type pc3ca883e2.JWTokenController in package:github.com/starter-go/v0/rbac-web-app/app/web/controllers/helper
//
// id:com-c3ca883e2525cff7-helper-JWTokenController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type pc3ca883e25_helper_JWTokenController struct {
}

func (inst* pc3ca883e25_helper_JWTokenController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-c3ca883e2525cff7-helper-JWTokenController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pc3ca883e25_helper_JWTokenController) new() any {
    return &pc3ca883e2.JWTokenController{}
}

func (inst* pc3ca883e25_helper_JWTokenController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pc3ca883e2.JWTokenController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.ChainHolder = inst.getChainHolder(ie)


    return nil
}


func (inst*pc3ca883e25_helper_JWTokenController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*pc3ca883e25_helper_JWTokenController) getChainHolder(ie application.InjectionExt)pfd2c28477.FilterChainHolder{
    return ie.GetComponent("#alias-fd2c28477d8555ea1fa4190037afa453-FilterChainHolder").(pfd2c28477.FilterChainHolder)
}



// type pc3ca883e2.GinLibjwtAdapter in package:github.com/starter-go/v0/rbac-web-app/app/web/controllers/helper
//
// id:com-c3ca883e2525cff7-helper-GinLibjwtAdapter
// class:class-40146ff2506e686679e1be8b766901e1-AdapterProvider
// alias:
// scope:singleton
//
type pc3ca883e25_helper_GinLibjwtAdapter struct {
}

func (inst* pc3ca883e25_helper_GinLibjwtAdapter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-c3ca883e2525cff7-helper-GinLibjwtAdapter"
	r.Classes = "class-40146ff2506e686679e1be8b766901e1-AdapterProvider"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pc3ca883e25_helper_GinLibjwtAdapter) new() any {
    return &pc3ca883e2.GinLibjwtAdapter{}
}

func (inst* pc3ca883e25_helper_GinLibjwtAdapter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pc3ca883e2.GinLibjwtAdapter)
	nop(ie, com)

	
    com.MyName = inst.getMyName(ie)
    com.MyPriority = inst.getMyPriority(ie)
    com.MyEnabled = inst.getMyEnabled(ie)
    com.MyUseCookie = inst.getMyUseCookie(ie)
    com.MyUseHeader = inst.getMyUseHeader(ie)
    com.MyUseQuery = inst.getMyUseQuery(ie)
    com.MyCookieName = inst.getMyCookieName(ie)
    com.MyCookieMaxAge = inst.getMyCookieMaxAge(ie)
    com.MyCookiePath = inst.getMyCookiePath(ie)
    com.MyCookieDomain = inst.getMyCookieDomain(ie)
    com.MyCookieSecure = inst.getMyCookieSecure(ie)
    com.MyCookieHttpOnly = inst.getMyCookieHttpOnly(ie)


    return nil
}


func (inst*pc3ca883e25_helper_GinLibjwtAdapter) getMyName(ie application.InjectionExt)string{
    return ie.GetString("${jwt-adapter.libgin.name}")
}


func (inst*pc3ca883e25_helper_GinLibjwtAdapter) getMyPriority(ie application.InjectionExt)int{
    return ie.GetInt("${jwt-adapter.libgin.priority}")
}


func (inst*pc3ca883e25_helper_GinLibjwtAdapter) getMyEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${jwt-adapter.libgin.enabled}")
}


func (inst*pc3ca883e25_helper_GinLibjwtAdapter) getMyUseCookie(ie application.InjectionExt)bool{
    return ie.GetBool("${jwt-adapter.libgin.use-http-cookie}")
}


func (inst*pc3ca883e25_helper_GinLibjwtAdapter) getMyUseHeader(ie application.InjectionExt)bool{
    return ie.GetBool("${jwt-adapter.libgin.use-http-header}")
}


func (inst*pc3ca883e25_helper_GinLibjwtAdapter) getMyUseQuery(ie application.InjectionExt)bool{
    return ie.GetBool("${jwt-adapter.libgin.use-http-query}")
}


func (inst*pc3ca883e25_helper_GinLibjwtAdapter) getMyCookieName(ie application.InjectionExt)string{
    return ie.GetString("${jwt-adapter.libgin.cookie.name}")
}


func (inst*pc3ca883e25_helper_GinLibjwtAdapter) getMyCookieMaxAge(ie application.InjectionExt)int{
    return ie.GetInt("${jwt-adapter.libgin.cookie.max-age}")
}


func (inst*pc3ca883e25_helper_GinLibjwtAdapter) getMyCookiePath(ie application.InjectionExt)string{
    return ie.GetString("${jwt-adapter.libgin.cookie.path}")
}


func (inst*pc3ca883e25_helper_GinLibjwtAdapter) getMyCookieDomain(ie application.InjectionExt)string{
    return ie.GetString("${jwt-adapter.libgin.cookie.domain}")
}


func (inst*pc3ca883e25_helper_GinLibjwtAdapter) getMyCookieSecure(ie application.InjectionExt)bool{
    return ie.GetBool("${jwt-adapter.libgin.cookie.secure}")
}


func (inst*pc3ca883e25_helper_GinLibjwtAdapter) getMyCookieHttpOnly(ie application.InjectionExt)bool{
    return ie.GetBool("${jwt-adapter.libgin.cookie.http-only}")
}



// type p79b61f17f.AuthxController in package:github.com/starter-go/v0/rbac-web-app/app/web/controllers/home
//
// id:com-79b61f17f767d154-home-AuthxController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p79b61f17f7_home_AuthxController struct {
}

func (inst* p79b61f17f7_home_AuthxController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-79b61f17f767d154-home-AuthxController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p79b61f17f7_home_AuthxController) new() any {
    return &p79b61f17f.AuthxController{}
}

func (inst* p79b61f17f7_home_AuthxController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p79b61f17f.AuthxController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p79b61f17f7_home_AuthxController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p79b61f17f7_home_AuthxController) getService(ie application.InjectionExt)p84e86b31b.Service{
    return ie.GetComponent("#alias-84e86b31b37511e994be50a820fadd04-Service").(p84e86b31b.Service)
}



// type p79b61f17f.SessionController in package:github.com/starter-go/v0/rbac-web-app/app/web/controllers/home
//
// id:com-79b61f17f767d154-home-SessionController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p79b61f17f7_home_SessionController struct {
}

func (inst* p79b61f17f7_home_SessionController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-79b61f17f767d154-home-SessionController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p79b61f17f7_home_SessionController) new() any {
    return &p79b61f17f.SessionController{}
}

func (inst* p79b61f17f7_home_SessionController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p79b61f17f.SessionController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)


    return nil
}


func (inst*p79b61f17f7_home_SessionController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


