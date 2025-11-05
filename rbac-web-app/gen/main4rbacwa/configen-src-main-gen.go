package main4rbacwa
import (
    pd1a916a20 "github.com/starter-go/libgin"
    p512a30914 "github.com/starter-go/libgorm"
    p9621e8b71 "github.com/starter-go/security/random"
    p84e86b31b "github.com/starter-go/v0/rbac-web-app/app/classes/authx"
    p63b715cf9 "github.com/starter-go/v0/rbac-web-app/app/classes/sessions"
    p441b1814d "github.com/starter-go/v0/rbac-web-app/app/classes/tokens"
    p2d258e798 "github.com/starter-go/v0/rbac-web-app/app/classes/users"
    pf75fd20ca "github.com/starter-go/v0/rbac-web-app/app/data/database"
    p383a5f3ee "github.com/starter-go/v0/rbac-web-app/app/implements/iauthx"
    pabcfb7520 "github.com/starter-go/v0/rbac-web-app/app/implements/idatabase"
    p8841b2383 "github.com/starter-go/v0/rbac-web-app/app/implements/isessions"
    p3821d5b0f "github.com/starter-go/v0/rbac-web-app/app/implements/itokens"
    p4b670e5fa "github.com/starter-go/v0/rbac-web-app/app/implements/iusers"
    p7249a1596 "github.com/starter-go/v0/rbac-web-app/app/web/controllers"
    p7f79a0bbf "github.com/starter-go/v0/rbac-web-app/app/web/controllers/admin"
    p0fcabc88b "github.com/starter-go/v0/rbac-web-app/app/web/controllers/debug"
    p79b61f17f "github.com/starter-go/v0/rbac-web-app/app/web/controllers/home"
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



// type p8841b2383.SessionDaoImpl in package:github.com/starter-go/v0/rbac-web-app/app/implements/isessions
//
// id:com-8841b23836065757-isessions-SessionDaoImpl
// class:
// alias:alias-63b715cf9c89c4b9295921d813a75e67-DAO
// scope:singleton
//
type p8841b23836_isessions_SessionDaoImpl struct {
}

func (inst* p8841b23836_isessions_SessionDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-8841b23836065757-isessions-SessionDaoImpl"
	r.Classes = ""
	r.Aliases = "alias-63b715cf9c89c4b9295921d813a75e67-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p8841b23836_isessions_SessionDaoImpl) new() any {
    return &p8841b2383.SessionDaoImpl{}
}

func (inst* p8841b23836_isessions_SessionDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p8841b2383.SessionDaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDSer = inst.getUUIDSer(ie)


    return nil
}


func (inst*p8841b23836_isessions_SessionDaoImpl) getAgent(ie application.InjectionExt)pf75fd20ca.Agent{
    return ie.GetComponent("#alias-f75fd20ca4f14f60caace00a3f949b98-Agent").(pf75fd20ca.Agent)
}


func (inst*p8841b23836_isessions_SessionDaoImpl) getUUIDSer(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type p8841b2383.SessionServiceImpl in package:github.com/starter-go/v0/rbac-web-app/app/implements/isessions
//
// id:com-8841b23836065757-isessions-SessionServiceImpl
// class:
// alias:alias-63b715cf9c89c4b9295921d813a75e67-Service
// scope:singleton
//
type p8841b23836_isessions_SessionServiceImpl struct {
}

func (inst* p8841b23836_isessions_SessionServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-8841b23836065757-isessions-SessionServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-63b715cf9c89c4b9295921d813a75e67-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p8841b23836_isessions_SessionServiceImpl) new() any {
    return &p8841b2383.SessionServiceImpl{}
}

func (inst* p8841b23836_isessions_SessionServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p8841b2383.SessionServiceImpl)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*p8841b23836_isessions_SessionServiceImpl) getDao(ie application.InjectionExt)p63b715cf9.DAO{
    return ie.GetComponent("#alias-63b715cf9c89c4b9295921d813a75e67-DAO").(p63b715cf9.DAO)
}



// type p3821d5b0f.JWTokenServiceImpl in package:github.com/starter-go/v0/rbac-web-app/app/implements/itokens
//
// id:com-3821d5b0f6716e9c-itokens-JWTokenServiceImpl
// class:
// alias:alias-441b1814dcb3d885c070bf59d5773c63-JWTokenService
// scope:singleton
//
type p3821d5b0f6_itokens_JWTokenServiceImpl struct {
}

func (inst* p3821d5b0f6_itokens_JWTokenServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-3821d5b0f6716e9c-itokens-JWTokenServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-441b1814dcb3d885c070bf59d5773c63-JWTokenService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p3821d5b0f6_itokens_JWTokenServiceImpl) new() any {
    return &p3821d5b0f.JWTokenServiceImpl{}
}

func (inst* p3821d5b0f6_itokens_JWTokenServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p3821d5b0f.JWTokenServiceImpl)
	nop(ie, com)

	
    com.ConfigUseCookie = inst.getConfigUseCookie(ie)
    com.ConfigUseHeader = inst.getConfigUseHeader(ie)
    com.ConfigKeySecret = inst.getConfigKeySecret(ie)
    com.ConfigJWTokenMaxAgeSec = inst.getConfigJWTokenMaxAgeSec(ie)
    com.ConfigDoVerify = inst.getConfigDoVerify(ie)


    return nil
}


func (inst*p3821d5b0f6_itokens_JWTokenServiceImpl) getConfigUseCookie(ie application.InjectionExt)bool{
    return ie.GetBool("${server.www.jwt.use-cookie}")
}


func (inst*p3821d5b0f6_itokens_JWTokenServiceImpl) getConfigUseHeader(ie application.InjectionExt)bool{
    return ie.GetBool("${server.www.jwt.use-header}")
}


func (inst*p3821d5b0f6_itokens_JWTokenServiceImpl) getConfigKeySecret(ie application.InjectionExt)string{
    return ie.GetString("${server.www.jwt.key-secret}")
}


func (inst*p3821d5b0f6_itokens_JWTokenServiceImpl) getConfigJWTokenMaxAgeSec(ie application.InjectionExt)int{
    return ie.GetInt("${server.www.jwt.max-age-in-sec}")
}


func (inst*p3821d5b0f6_itokens_JWTokenServiceImpl) getConfigDoVerify(ie application.InjectionExt)bool{
    return ie.GetBool("${server.www.jwt.do-verify}")
}



// type p3821d5b0f.TokenServiceImpl in package:github.com/starter-go/v0/rbac-web-app/app/implements/itokens
//
// id:com-3821d5b0f6716e9c-itokens-TokenServiceImpl
// class:
// alias:alias-441b1814dcb3d885c070bf59d5773c63-Service
// scope:singleton
//
type p3821d5b0f6_itokens_TokenServiceImpl struct {
}

func (inst* p3821d5b0f6_itokens_TokenServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-3821d5b0f6716e9c-itokens-TokenServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-441b1814dcb3d885c070bf59d5773c63-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p3821d5b0f6_itokens_TokenServiceImpl) new() any {
    return &p3821d5b0f.TokenServiceImpl{}
}

func (inst* p3821d5b0f6_itokens_TokenServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p3821d5b0f.TokenServiceImpl)
	nop(ie, com)

	
    com.JWTser = inst.getJWTser(ie)


    return nil
}


func (inst*p3821d5b0f6_itokens_TokenServiceImpl) getJWTser(ie application.InjectionExt)p441b1814d.JWTokenService{
    return ie.GetComponent("#alias-441b1814dcb3d885c070bf59d5773c63-JWTokenService").(p441b1814d.JWTokenService)
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



// type p0fcabc88b.MockController in package:github.com/starter-go/v0/rbac-web-app/app/web/controllers/debug
//
// id:com-0fcabc88b927ca96-debug-MockController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p0fcabc88b9_debug_MockController struct {
}

func (inst* p0fcabc88b9_debug_MockController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-0fcabc88b927ca96-debug-MockController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p0fcabc88b9_debug_MockController) new() any {
    return &p0fcabc88b.MockController{}
}

func (inst* p0fcabc88b9_debug_MockController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p0fcabc88b.MockController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.TokenService = inst.getTokenService(ie)
    com.SessionService = inst.getSessionService(ie)
    com.UserService = inst.getUserService(ie)


    return nil
}


func (inst*p0fcabc88b9_debug_MockController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p0fcabc88b9_debug_MockController) getTokenService(ie application.InjectionExt)p441b1814d.Service{
    return ie.GetComponent("#alias-441b1814dcb3d885c070bf59d5773c63-Service").(p441b1814d.Service)
}


func (inst*p0fcabc88b9_debug_MockController) getSessionService(ie application.InjectionExt)p63b715cf9.Service{
    return ie.GetComponent("#alias-63b715cf9c89c4b9295921d813a75e67-Service").(p63b715cf9.Service)
}


func (inst*p0fcabc88b9_debug_MockController) getUserService(ie application.InjectionExt)p2d258e798.Service{
    return ie.GetComponent("#alias-2d258e79845135c43979b78bd9f1f74e-Service").(p2d258e798.Service)
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



// type p79b61f17f.JWTokenController in package:github.com/starter-go/v0/rbac-web-app/app/web/controllers/home
//
// id:com-79b61f17f767d154-home-JWTokenController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p79b61f17f7_home_JWTokenController struct {
}

func (inst* p79b61f17f7_home_JWTokenController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-79b61f17f767d154-home-JWTokenController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p79b61f17f7_home_JWTokenController) new() any {
    return &p79b61f17f.JWTokenController{}
}

func (inst* p79b61f17f7_home_JWTokenController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p79b61f17f.JWTokenController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.TokenService = inst.getTokenService(ie)
    com.SessionService = inst.getSessionService(ie)
    com.UserService = inst.getUserService(ie)


    return nil
}


func (inst*p79b61f17f7_home_JWTokenController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p79b61f17f7_home_JWTokenController) getTokenService(ie application.InjectionExt)p441b1814d.Service{
    return ie.GetComponent("#alias-441b1814dcb3d885c070bf59d5773c63-Service").(p441b1814d.Service)
}


func (inst*p79b61f17f7_home_JWTokenController) getSessionService(ie application.InjectionExt)p63b715cf9.Service{
    return ie.GetComponent("#alias-63b715cf9c89c4b9295921d813a75e67-Service").(p63b715cf9.Service)
}


func (inst*p79b61f17f7_home_JWTokenController) getUserService(ie application.InjectionExt)p2d258e798.Service{
    return ie.GetComponent("#alias-2d258e79845135c43979b78bd9f1f74e-Service").(p2d258e798.Service)
}


