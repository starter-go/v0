package main4rbacdg
import (
    p512a30914 "github.com/starter-go/libgorm"
    p9621e8b71 "github.com/starter-go/security/random"
    p89b1fdc13 "github.com/starter-go/v0/rbac-data-group/src/main/golang/api/daos"
    pfa85f3231 "github.com/starter-go/v0/rbac-data-group/src/main/golang/lib/implementations/iauths"
    p3da2710ff "github.com/starter-go/v0/rbac-data-group/src/main/golang/lib/implementations/idatagroup"
    pe64067b08 "github.com/starter-go/v0/rbac-data-group/src/main/golang/lib/implementations/ipermissions"
    p651d61942 "github.com/starter-go/v0/rbac-data-group/src/main/golang/lib/implementations/iroles"
    p01fd85955 "github.com/starter-go/v0/rbac-data-group/src/main/golang/lib/implementations/isessions"
    pb437cfd58 "github.com/starter-go/v0/rbac-data-group/src/main/golang/lib/implementations/itables"
    pa2c28bd10 "github.com/starter-go/v0/rbac-data-group/src/main/golang/lib/implementations/iusers"
     "github.com/starter-go/application"
)

// type pfa85f3231.AuthentDaoImpl in package:github.com/starter-go/v0/rbac-data-group/src/main/golang/lib/implementations/iauths
//
// id:com-fa85f3231cb93af0-iauths-AuthentDaoImpl
// class:
// alias:alias-89b1fdc13b09035b5add53e6ef430892-IAuthenticationDao
// scope:singleton
//
type pfa85f3231c_iauths_AuthentDaoImpl struct {
}

func (inst* pfa85f3231c_iauths_AuthentDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-fa85f3231cb93af0-iauths-AuthentDaoImpl"
	r.Classes = ""
	r.Aliases = "alias-89b1fdc13b09035b5add53e6ef430892-IAuthenticationDao"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pfa85f3231c_iauths_AuthentDaoImpl) new() any {
    return &pfa85f3231.AuthentDaoImpl{}
}

func (inst* pfa85f3231c_iauths_AuthentDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pfa85f3231.AuthentDaoImpl)
	nop(ie, com)

	
    com.DBAgent = inst.getDBAgent(ie)
    com.UUIDGenSer = inst.getUUIDGenSer(ie)


    return nil
}


func (inst*pfa85f3231c_iauths_AuthentDaoImpl) getDBAgent(ie application.InjectionExt)p89b1fdc13.IDatabaseAgent{
    return ie.GetComponent("#alias-89b1fdc13b09035b5add53e6ef430892-IDatabaseAgent").(p89b1fdc13.IDatabaseAgent)
}


func (inst*pfa85f3231c_iauths_AuthentDaoImpl) getUUIDGenSer(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type p3da2710ff.StdRbacDaoSet in package:github.com/starter-go/v0/rbac-data-group/src/main/golang/lib/implementations/idatagroup
//
// id:com-3da2710ff4821723-idatagroup-StdRbacDaoSet
// class:class-24287f4589fe5add27fb48a88d706565-DaoSet
// alias:
// scope:singleton
//
type p3da2710ff4_idatagroup_StdRbacDaoSet struct {
}

func (inst* p3da2710ff4_idatagroup_StdRbacDaoSet) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-3da2710ff4821723-idatagroup-StdRbacDaoSet"
	r.Classes = "class-24287f4589fe5add27fb48a88d706565-DaoSet"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p3da2710ff4_idatagroup_StdRbacDaoSet) new() any {
    return &p3da2710ff.StdRbacDaoSet{}
}

func (inst* p3da2710ff4_idatagroup_StdRbacDaoSet) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p3da2710ff.StdRbacDaoSet)
	nop(ie, com)

	
    com.ConfigEnabled = inst.getConfigEnabled(ie)
    com.ConfigPriority = inst.getConfigPriority(ie)
    com.Authents = inst.getAuthents(ie)
    com.Permissions = inst.getPermissions(ie)
    com.Roles = inst.getRoles(ie)
    com.Sessions = inst.getSessions(ie)
    com.Tables = inst.getTables(ie)
    com.Users = inst.getUsers(ie)


    return nil
}


func (inst*p3da2710ff4_idatagroup_StdRbacDaoSet) getConfigEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${rbac-dao-set.standard.enabled}")
}


func (inst*p3da2710ff4_idatagroup_StdRbacDaoSet) getConfigPriority(ie application.InjectionExt)int{
    return ie.GetInt("${rbac-dao-set.standard.priority}")
}


func (inst*p3da2710ff4_idatagroup_StdRbacDaoSet) getAuthents(ie application.InjectionExt)p89b1fdc13.IAuthenticationDao{
    return ie.GetComponent("#alias-89b1fdc13b09035b5add53e6ef430892-IAuthenticationDao").(p89b1fdc13.IAuthenticationDao)
}


func (inst*p3da2710ff4_idatagroup_StdRbacDaoSet) getPermissions(ie application.InjectionExt)p89b1fdc13.IPermissionDao{
    return ie.GetComponent("#alias-89b1fdc13b09035b5add53e6ef430892-IPermissionDao").(p89b1fdc13.IPermissionDao)
}


func (inst*p3da2710ff4_idatagroup_StdRbacDaoSet) getRoles(ie application.InjectionExt)p89b1fdc13.IRoleDao{
    return ie.GetComponent("#alias-89b1fdc13b09035b5add53e6ef430892-IRoleDao").(p89b1fdc13.IRoleDao)
}


func (inst*p3da2710ff4_idatagroup_StdRbacDaoSet) getSessions(ie application.InjectionExt)p89b1fdc13.ISessionDao{
    return ie.GetComponent("#alias-89b1fdc13b09035b5add53e6ef430892-ISessionDao").(p89b1fdc13.ISessionDao)
}


func (inst*p3da2710ff4_idatagroup_StdRbacDaoSet) getTables(ie application.InjectionExt)p89b1fdc13.ITableDao{
    return ie.GetComponent("#alias-89b1fdc13b09035b5add53e6ef430892-ITableDao").(p89b1fdc13.ITableDao)
}


func (inst*p3da2710ff4_idatagroup_StdRbacDaoSet) getUsers(ie application.InjectionExt)p89b1fdc13.IUserDao{
    return ie.GetComponent("#alias-89b1fdc13b09035b5add53e6ef430892-IUserDao").(p89b1fdc13.IUserDao)
}



// type p3da2710ff.StdRbacDataGroup in package:github.com/starter-go/v0/rbac-data-group/src/main/golang/lib/implementations/idatagroup
//
// id:com-3da2710ff4821723-idatagroup-StdRbacDataGroup
// class:class-512a309140d0ad99eb1c95c8dc0d02f9-GroupRegistry
// alias:
// scope:singleton
//
type p3da2710ff4_idatagroup_StdRbacDataGroup struct {
}

func (inst* p3da2710ff4_idatagroup_StdRbacDataGroup) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-3da2710ff4821723-idatagroup-StdRbacDataGroup"
	r.Classes = "class-512a309140d0ad99eb1c95c8dc0d02f9-GroupRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p3da2710ff4_idatagroup_StdRbacDataGroup) new() any {
    return &p3da2710ff.StdRbacDataGroup{}
}

func (inst* p3da2710ff4_idatagroup_StdRbacDataGroup) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p3da2710ff.StdRbacDataGroup)
	nop(ie, com)

	
    com.ConfigAlias = inst.getConfigAlias(ie)
    com.ConfigEnabled = inst.getConfigEnabled(ie)
    com.ConfigPrefix = inst.getConfigPrefix(ie)
    com.ConfigSource = inst.getConfigSource(ie)
    com.ConfigURI = inst.getConfigURI(ie)


    return nil
}


func (inst*p3da2710ff4_idatagroup_StdRbacDataGroup) getConfigAlias(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.std-rbac-dg.alias}")
}


func (inst*p3da2710ff4_idatagroup_StdRbacDataGroup) getConfigEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${datagroup.std-rbac-dg.enabled}")
}


func (inst*p3da2710ff4_idatagroup_StdRbacDataGroup) getConfigPrefix(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.std-rbac-dg.table-name-prefix}")
}


func (inst*p3da2710ff4_idatagroup_StdRbacDataGroup) getConfigSource(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.std-rbac-dg.datasource}")
}


func (inst*p3da2710ff4_idatagroup_StdRbacDataGroup) getConfigURI(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.std-rbac-dg.uri}")
}



// type p3da2710ff.StdRbacDataAgent in package:github.com/starter-go/v0/rbac-data-group/src/main/golang/lib/implementations/idatagroup
//
// id:com-3da2710ff4821723-idatagroup-StdRbacDataAgent
// class:
// alias:alias-89b1fdc13b09035b5add53e6ef430892-IDatabaseAgent
// scope:singleton
//
type p3da2710ff4_idatagroup_StdRbacDataAgent struct {
}

func (inst* p3da2710ff4_idatagroup_StdRbacDataAgent) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-3da2710ff4821723-idatagroup-StdRbacDataAgent"
	r.Classes = ""
	r.Aliases = "alias-89b1fdc13b09035b5add53e6ef430892-IDatabaseAgent"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p3da2710ff4_idatagroup_StdRbacDataAgent) new() any {
    return &p3da2710ff.StdRbacDataAgent{}
}

func (inst* p3da2710ff4_idatagroup_StdRbacDataAgent) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p3da2710ff.StdRbacDataAgent)
	nop(ie, com)

	
    com.DSM = inst.getDSM(ie)
    com.ConfigAlias = inst.getConfigAlias(ie)


    return nil
}


func (inst*p3da2710ff4_idatagroup_StdRbacDataAgent) getDSM(ie application.InjectionExt)p512a30914.DataSourceManager{
    return ie.GetComponent("#alias-512a309140d0ad99eb1c95c8dc0d02f9-DataSourceManager").(p512a30914.DataSourceManager)
}


func (inst*p3da2710ff4_idatagroup_StdRbacDataAgent) getConfigAlias(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.std-rbac-dg.alias}")
}



// type pe64067b08.PermissionDaoImpl in package:github.com/starter-go/v0/rbac-data-group/src/main/golang/lib/implementations/ipermissions
//
// id:com-e64067b08ccba2a2-ipermissions-PermissionDaoImpl
// class:
// alias:alias-89b1fdc13b09035b5add53e6ef430892-IPermissionDao
// scope:singleton
//
type pe64067b08c_ipermissions_PermissionDaoImpl struct {
}

func (inst* pe64067b08c_ipermissions_PermissionDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-e64067b08ccba2a2-ipermissions-PermissionDaoImpl"
	r.Classes = ""
	r.Aliases = "alias-89b1fdc13b09035b5add53e6ef430892-IPermissionDao"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pe64067b08c_ipermissions_PermissionDaoImpl) new() any {
    return &pe64067b08.PermissionDaoImpl{}
}

func (inst* pe64067b08c_ipermissions_PermissionDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pe64067b08.PermissionDaoImpl)
	nop(ie, com)

	
    com.DBAgent = inst.getDBAgent(ie)
    com.UUIDGenSer = inst.getUUIDGenSer(ie)


    return nil
}


func (inst*pe64067b08c_ipermissions_PermissionDaoImpl) getDBAgent(ie application.InjectionExt)p89b1fdc13.IDatabaseAgent{
    return ie.GetComponent("#alias-89b1fdc13b09035b5add53e6ef430892-IDatabaseAgent").(p89b1fdc13.IDatabaseAgent)
}


func (inst*pe64067b08c_ipermissions_PermissionDaoImpl) getUUIDGenSer(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type p651d61942.RoleDaoImpl in package:github.com/starter-go/v0/rbac-data-group/src/main/golang/lib/implementations/iroles
//
// id:com-651d6194286589c1-iroles-RoleDaoImpl
// class:
// alias:alias-89b1fdc13b09035b5add53e6ef430892-IRoleDao
// scope:singleton
//
type p651d619428_iroles_RoleDaoImpl struct {
}

func (inst* p651d619428_iroles_RoleDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-651d6194286589c1-iroles-RoleDaoImpl"
	r.Classes = ""
	r.Aliases = "alias-89b1fdc13b09035b5add53e6ef430892-IRoleDao"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p651d619428_iroles_RoleDaoImpl) new() any {
    return &p651d61942.RoleDaoImpl{}
}

func (inst* p651d619428_iroles_RoleDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p651d61942.RoleDaoImpl)
	nop(ie, com)

	
    com.DBAgent = inst.getDBAgent(ie)
    com.UUIDGenSer = inst.getUUIDGenSer(ie)


    return nil
}


func (inst*p651d619428_iroles_RoleDaoImpl) getDBAgent(ie application.InjectionExt)p89b1fdc13.IDatabaseAgent{
    return ie.GetComponent("#alias-89b1fdc13b09035b5add53e6ef430892-IDatabaseAgent").(p89b1fdc13.IDatabaseAgent)
}


func (inst*p651d619428_iroles_RoleDaoImpl) getUUIDGenSer(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type p01fd85955.SessionDaoImpl in package:github.com/starter-go/v0/rbac-data-group/src/main/golang/lib/implementations/isessions
//
// id:com-01fd85955800a2c6-isessions-SessionDaoImpl
// class:
// alias:alias-89b1fdc13b09035b5add53e6ef430892-ISessionDao
// scope:singleton
//
type p01fd859558_isessions_SessionDaoImpl struct {
}

func (inst* p01fd859558_isessions_SessionDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-01fd85955800a2c6-isessions-SessionDaoImpl"
	r.Classes = ""
	r.Aliases = "alias-89b1fdc13b09035b5add53e6ef430892-ISessionDao"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p01fd859558_isessions_SessionDaoImpl) new() any {
    return &p01fd85955.SessionDaoImpl{}
}

func (inst* p01fd859558_isessions_SessionDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p01fd85955.SessionDaoImpl)
	nop(ie, com)

	
    com.DBAgent = inst.getDBAgent(ie)
    com.UUIDGenSer = inst.getUUIDGenSer(ie)


    return nil
}


func (inst*p01fd859558_isessions_SessionDaoImpl) getDBAgent(ie application.InjectionExt)p89b1fdc13.IDatabaseAgent{
    return ie.GetComponent("#alias-89b1fdc13b09035b5add53e6ef430892-IDatabaseAgent").(p89b1fdc13.IDatabaseAgent)
}


func (inst*p01fd859558_isessions_SessionDaoImpl) getUUIDGenSer(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type pb437cfd58.TableDaoImpl in package:github.com/starter-go/v0/rbac-data-group/src/main/golang/lib/implementations/itables
//
// id:com-b437cfd58063a43d-itables-TableDaoImpl
// class:
// alias:alias-89b1fdc13b09035b5add53e6ef430892-ITableDao
// scope:singleton
//
type pb437cfd580_itables_TableDaoImpl struct {
}

func (inst* pb437cfd580_itables_TableDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-b437cfd58063a43d-itables-TableDaoImpl"
	r.Classes = ""
	r.Aliases = "alias-89b1fdc13b09035b5add53e6ef430892-ITableDao"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pb437cfd580_itables_TableDaoImpl) new() any {
    return &pb437cfd58.TableDaoImpl{}
}

func (inst* pb437cfd580_itables_TableDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pb437cfd58.TableDaoImpl)
	nop(ie, com)

	
    com.DBAgent = inst.getDBAgent(ie)
    com.UUIDGenSer = inst.getUUIDGenSer(ie)


    return nil
}


func (inst*pb437cfd580_itables_TableDaoImpl) getDBAgent(ie application.InjectionExt)p89b1fdc13.IDatabaseAgent{
    return ie.GetComponent("#alias-89b1fdc13b09035b5add53e6ef430892-IDatabaseAgent").(p89b1fdc13.IDatabaseAgent)
}


func (inst*pb437cfd580_itables_TableDaoImpl) getUUIDGenSer(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type pa2c28bd10.UserDaoImpl in package:github.com/starter-go/v0/rbac-data-group/src/main/golang/lib/implementations/iusers
//
// id:com-a2c28bd102d9f26c-iusers-UserDaoImpl
// class:
// alias:alias-89b1fdc13b09035b5add53e6ef430892-IUserDao
// scope:singleton
//
type pa2c28bd102_iusers_UserDaoImpl struct {
}

func (inst* pa2c28bd102_iusers_UserDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a2c28bd102d9f26c-iusers-UserDaoImpl"
	r.Classes = ""
	r.Aliases = "alias-89b1fdc13b09035b5add53e6ef430892-IUserDao"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa2c28bd102_iusers_UserDaoImpl) new() any {
    return &pa2c28bd10.UserDaoImpl{}
}

func (inst* pa2c28bd102_iusers_UserDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa2c28bd10.UserDaoImpl)
	nop(ie, com)

	
    com.DBAgent = inst.getDBAgent(ie)
    com.UUIDGenSer = inst.getUUIDGenSer(ie)


    return nil
}


func (inst*pa2c28bd102_iusers_UserDaoImpl) getDBAgent(ie application.InjectionExt)p89b1fdc13.IDatabaseAgent{
    return ie.GetComponent("#alias-89b1fdc13b09035b5add53e6ef430892-IDatabaseAgent").(p89b1fdc13.IDatabaseAgent)
}


func (inst*pa2c28bd102_iusers_UserDaoImpl) getUUIDGenSer(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}


