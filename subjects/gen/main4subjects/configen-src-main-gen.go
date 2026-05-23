package main4subjects
import (
    p512a30914 "github.com/starter-go/libgorm"
    p9621e8b71 "github.com/starter-go/security/random"
    p40146ff25 "github.com/starter-go/v0/libjwt/api/jwt"
    pfd2c28477 "github.com/starter-go/v0/subjects"
    pdb4050fdb "github.com/starter-go/v0/subjects/core/autoloaders"
    p85899ef78 "github.com/starter-go/v0/subjects/core/classes/sessions"
    pe36083c14 "github.com/starter-go/v0/subjects/core/consoles"
    pba9b83e76 "github.com/starter-go/v0/subjects/core/data/subjectdb"
    p45b883c73 "github.com/starter-go/v0/subjects/core/filters/filter4buffer"
    pd3bd430b1 "github.com/starter-go/v0/subjects/core/filters/filter4cache"
    pe5401edcd "github.com/starter-go/v0/subjects/core/filters/filter4codec"
    p1e478b5d4 "github.com/starter-go/v0/subjects/core/filters/filter4db"
    p4270aadd3 "github.com/starter-go/v0/subjects/core/filters/filter4jwt"
    pd00479541 "github.com/starter-go/v0/subjects/core/filters/filter4log"
    p0d8344d2d "github.com/starter-go/v0/subjects/core/implements/isessions"
    p29a9720d8 "github.com/starter-go/v0/subjects/core/implements/isubjectdb"
     "github.com/starter-go/application"
)

// type pdb4050fdb.DefaultFilterChainHolder in package:github.com/starter-go/v0/subjects/core/autoloaders
//
// id:com-db4050fdbddde455-autoloaders-DefaultFilterChainHolder
// class:
// alias:alias-fd2c28477d8555ea1fa4190037afa453-FilterChainHolder
// scope:singleton
//
type pdb4050fdbd_autoloaders_DefaultFilterChainHolder struct {
}

func (inst* pdb4050fdbd_autoloaders_DefaultFilterChainHolder) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-db4050fdbddde455-autoloaders-DefaultFilterChainHolder"
	r.Classes = ""
	r.Aliases = "alias-fd2c28477d8555ea1fa4190037afa453-FilterChainHolder"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pdb4050fdbd_autoloaders_DefaultFilterChainHolder) new() any {
    return &pdb4050fdb.DefaultFilterChainHolder{}
}

func (inst* pdb4050fdbd_autoloaders_DefaultFilterChainHolder) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pdb4050fdb.DefaultFilterChainHolder)
	nop(ie, com)

	
    com.Loader = inst.getLoader(ie)


    return nil
}


func (inst*pdb4050fdbd_autoloaders_DefaultFilterChainHolder) getLoader(ie application.InjectionExt)pfd2c28477.FilterChainLoader{
    return ie.GetComponent("#alias-fd2c28477d8555ea1fa4190037afa453-FilterChainLoader").(pfd2c28477.FilterChainLoader)
}



// type pdb4050fdb.DefaultFilterChainLoader in package:github.com/starter-go/v0/subjects/core/autoloaders
//
// id:com-db4050fdbddde455-autoloaders-DefaultFilterChainLoader
// class:
// alias:alias-fd2c28477d8555ea1fa4190037afa453-FilterChainLoader
// scope:singleton
//
type pdb4050fdbd_autoloaders_DefaultFilterChainLoader struct {
}

func (inst* pdb4050fdbd_autoloaders_DefaultFilterChainLoader) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-db4050fdbddde455-autoloaders-DefaultFilterChainLoader"
	r.Classes = ""
	r.Aliases = "alias-fd2c28477d8555ea1fa4190037afa453-FilterChainLoader"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pdb4050fdbd_autoloaders_DefaultFilterChainLoader) new() any {
    return &pdb4050fdb.DefaultFilterChainLoader{}
}

func (inst* pdb4050fdbd_autoloaders_DefaultFilterChainLoader) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pdb4050fdb.DefaultFilterChainLoader)
	nop(ie, com)

	
    com.RegList = inst.getRegList(ie)


    return nil
}


func (inst*pdb4050fdbd_autoloaders_DefaultFilterChainLoader) getRegList(ie application.InjectionExt)[]pfd2c28477.FilterRegistry{
    dst := make([]pfd2c28477.FilterRegistry, 0)
    src := ie.ListComponents(".class-fd2c28477d8555ea1fa4190037afa453-FilterRegistry")
    for _, item1 := range src {
        item2 := item1.(pfd2c28477.FilterRegistry)
        dst = append(dst, item2)
    }
    return dst
}



// type pe36083c14.ConsoleJwtAdapter in package:github.com/starter-go/v0/subjects/core/consoles
//
// id:com-e36083c14cb95b87-consoles-ConsoleJwtAdapter
// class:class-40146ff2506e686679e1be8b766901e1-AdapterProvider
// alias:
// scope:singleton
//
type pe36083c14c_consoles_ConsoleJwtAdapter struct {
}

func (inst* pe36083c14c_consoles_ConsoleJwtAdapter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-e36083c14cb95b87-consoles-ConsoleJwtAdapter"
	r.Classes = "class-40146ff2506e686679e1be8b766901e1-AdapterProvider"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pe36083c14c_consoles_ConsoleJwtAdapter) new() any {
    return &pe36083c14.ConsoleJwtAdapter{}
}

func (inst* pe36083c14c_consoles_ConsoleJwtAdapter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pe36083c14.ConsoleJwtAdapter)
	nop(ie, com)

	
    com.MyName = inst.getMyName(ie)
    com.MyEnabled = inst.getMyEnabled(ie)
    com.MyPriority = inst.getMyPriority(ie)


    return nil
}


func (inst*pe36083c14c_consoles_ConsoleJwtAdapter) getMyName(ie application.InjectionExt)string{
    return ie.GetString("${jwt-adapter.console.name}")
}


func (inst*pe36083c14c_consoles_ConsoleJwtAdapter) getMyEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${jwt-adapter.console.enabled}")
}


func (inst*pe36083c14c_consoles_ConsoleJwtAdapter) getMyPriority(ie application.InjectionExt)int{
    return ie.GetInt("${jwt-adapter.console.priority}")
}



// type p45b883c73.Filter4Buffer in package:github.com/starter-go/v0/subjects/core/filters/filter4buffer
//
// id:com-45b883c7373e9f00-filter4buffer-Filter4Buffer
// class:class-fd2c28477d8555ea1fa4190037afa453-FilterRegistry
// alias:
// scope:singleton
//
type p45b883c737_filter4buffer_Filter4Buffer struct {
}

func (inst* p45b883c737_filter4buffer_Filter4Buffer) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-45b883c7373e9f00-filter4buffer-Filter4Buffer"
	r.Classes = "class-fd2c28477d8555ea1fa4190037afa453-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p45b883c737_filter4buffer_Filter4Buffer) new() any {
    return &p45b883c73.Filter4Buffer{}
}

func (inst* p45b883c737_filter4buffer_Filter4Buffer) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p45b883c73.Filter4Buffer)
	nop(ie, com)

	


    return nil
}



// type pd3bd430b1.Filter4Cache1 in package:github.com/starter-go/v0/subjects/core/filters/filter4cache
//
// id:com-d3bd430b1e9a539a-filter4cache-Filter4Cache1
// class:class-fd2c28477d8555ea1fa4190037afa453-FilterRegistry
// alias:
// scope:singleton
//
type pd3bd430b1e_filter4cache_Filter4Cache1 struct {
}

func (inst* pd3bd430b1e_filter4cache_Filter4Cache1) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-d3bd430b1e9a539a-filter4cache-Filter4Cache1"
	r.Classes = "class-fd2c28477d8555ea1fa4190037afa453-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pd3bd430b1e_filter4cache_Filter4Cache1) new() any {
    return &pd3bd430b1.Filter4Cache1{}
}

func (inst* pd3bd430b1e_filter4cache_Filter4Cache1) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pd3bd430b1.Filter4Cache1)
	nop(ie, com)

	


    return nil
}



// type pd3bd430b1.Filter4Cache2 in package:github.com/starter-go/v0/subjects/core/filters/filter4cache
//
// id:com-d3bd430b1e9a539a-filter4cache-Filter4Cache2
// class:class-fd2c28477d8555ea1fa4190037afa453-FilterRegistry
// alias:
// scope:singleton
//
type pd3bd430b1e_filter4cache_Filter4Cache2 struct {
}

func (inst* pd3bd430b1e_filter4cache_Filter4Cache2) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-d3bd430b1e9a539a-filter4cache-Filter4Cache2"
	r.Classes = "class-fd2c28477d8555ea1fa4190037afa453-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pd3bd430b1e_filter4cache_Filter4Cache2) new() any {
    return &pd3bd430b1.Filter4Cache2{}
}

func (inst* pd3bd430b1e_filter4cache_Filter4Cache2) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pd3bd430b1.Filter4Cache2)
	nop(ie, com)

	


    return nil
}



// type pd3bd430b1.Filter4GoodResult in package:github.com/starter-go/v0/subjects/core/filters/filter4cache
//
// id:com-d3bd430b1e9a539a-filter4cache-Filter4GoodResult
// class:class-fd2c28477d8555ea1fa4190037afa453-FilterRegistry
// alias:
// scope:singleton
//
type pd3bd430b1e_filter4cache_Filter4GoodResult struct {
}

func (inst* pd3bd430b1e_filter4cache_Filter4GoodResult) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-d3bd430b1e9a539a-filter4cache-Filter4GoodResult"
	r.Classes = "class-fd2c28477d8555ea1fa4190037afa453-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pd3bd430b1e_filter4cache_Filter4GoodResult) new() any {
    return &pd3bd430b1.Filter4GoodResult{}
}

func (inst* pd3bd430b1e_filter4cache_Filter4GoodResult) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pd3bd430b1.Filter4GoodResult)
	nop(ie, com)

	


    return nil
}



// type pe5401edcd.Filter4Codec in package:github.com/starter-go/v0/subjects/core/filters/filter4codec
//
// id:com-e5401edcdeb1c54f-filter4codec-Filter4Codec
// class:class-fd2c28477d8555ea1fa4190037afa453-FilterRegistry
// alias:
// scope:singleton
//
type pe5401edcde_filter4codec_Filter4Codec struct {
}

func (inst* pe5401edcde_filter4codec_Filter4Codec) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-e5401edcdeb1c54f-filter4codec-Filter4Codec"
	r.Classes = "class-fd2c28477d8555ea1fa4190037afa453-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pe5401edcde_filter4codec_Filter4Codec) new() any {
    return &pe5401edcd.Filter4Codec{}
}

func (inst* pe5401edcde_filter4codec_Filter4Codec) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pe5401edcd.Filter4Codec)
	nop(ie, com)

	


    return nil
}



// type p1e478b5d4.Filter4db in package:github.com/starter-go/v0/subjects/core/filters/filter4db
//
// id:com-1e478b5d42300309-filter4db-Filter4db
// class:class-fd2c28477d8555ea1fa4190037afa453-FilterRegistry
// alias:
// scope:singleton
//
type p1e478b5d42_filter4db_Filter4db struct {
}

func (inst* p1e478b5d42_filter4db_Filter4db) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1e478b5d42300309-filter4db-Filter4db"
	r.Classes = "class-fd2c28477d8555ea1fa4190037afa453-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1e478b5d42_filter4db_Filter4db) new() any {
    return &p1e478b5d4.Filter4db{}
}

func (inst* p1e478b5d42_filter4db_Filter4db) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1e478b5d4.Filter4db)
	nop(ie, com)

	
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p1e478b5d42_filter4db_Filter4db) getService(ie application.InjectionExt)p85899ef78.Service{
    return ie.GetComponent("#alias-85899ef785c3f033c4d8293618ff61de-Service").(p85899ef78.Service)
}



// type p4270aadd3.Filter4jwt in package:github.com/starter-go/v0/subjects/core/filters/filter4jwt
//
// id:com-4270aadd3a4a14c5-filter4jwt-Filter4jwt
// class:class-fd2c28477d8555ea1fa4190037afa453-FilterRegistry
// alias:
// scope:singleton
//
type p4270aadd3a_filter4jwt_Filter4jwt struct {
}

func (inst* p4270aadd3a_filter4jwt_Filter4jwt) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-4270aadd3a4a14c5-filter4jwt-Filter4jwt"
	r.Classes = "class-fd2c28477d8555ea1fa4190037afa453-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p4270aadd3a_filter4jwt_Filter4jwt) new() any {
    return &p4270aadd3.Filter4jwt{}
}

func (inst* p4270aadd3a_filter4jwt_Filter4jwt) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p4270aadd3.Filter4jwt)
	nop(ie, com)

	
    com.JWTService = inst.getJWTService(ie)


    return nil
}


func (inst*p4270aadd3a_filter4jwt_Filter4jwt) getJWTService(ie application.InjectionExt)p40146ff25.Service{
    return ie.GetComponent("#alias-40146ff2506e686679e1be8b766901e1-Service").(p40146ff25.Service)
}



// type pd00479541.Filter4Log in package:github.com/starter-go/v0/subjects/core/filters/filter4log
//
// id:com-d00479541ffec028-filter4log-Filter4Log
// class:class-fd2c28477d8555ea1fa4190037afa453-FilterRegistry
// alias:
// scope:singleton
//
type pd00479541f_filter4log_Filter4Log struct {
}

func (inst* pd00479541f_filter4log_Filter4Log) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-d00479541ffec028-filter4log-Filter4Log"
	r.Classes = "class-fd2c28477d8555ea1fa4190037afa453-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pd00479541f_filter4log_Filter4Log) new() any {
    return &pd00479541.Filter4Log{}
}

func (inst* pd00479541f_filter4log_Filter4Log) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pd00479541.Filter4Log)
	nop(ie, com)

	


    return nil
}



// type p0d8344d2d.SessionDaoImpl in package:github.com/starter-go/v0/subjects/core/implements/isessions
//
// id:com-0d8344d2dcb25264-isessions-SessionDaoImpl
// class:
// alias:alias-85899ef785c3f033c4d8293618ff61de-DAO
// scope:singleton
//
type p0d8344d2dc_isessions_SessionDaoImpl struct {
}

func (inst* p0d8344d2dc_isessions_SessionDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-0d8344d2dcb25264-isessions-SessionDaoImpl"
	r.Classes = ""
	r.Aliases = "alias-85899ef785c3f033c4d8293618ff61de-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p0d8344d2dc_isessions_SessionDaoImpl) new() any {
    return &p0d8344d2d.SessionDaoImpl{}
}

func (inst* p0d8344d2dc_isessions_SessionDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p0d8344d2d.SessionDaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDSer = inst.getUUIDSer(ie)


    return nil
}


func (inst*p0d8344d2dc_isessions_SessionDaoImpl) getAgent(ie application.InjectionExt)pba9b83e76.Agent{
    return ie.GetComponent("#alias-ba9b83e768291a6dc6c6a9a0b0051833-Agent").(pba9b83e76.Agent)
}


func (inst*p0d8344d2dc_isessions_SessionDaoImpl) getUUIDSer(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type p0d8344d2d.SessionServiceImpl in package:github.com/starter-go/v0/subjects/core/implements/isessions
//
// id:com-0d8344d2dcb25264-isessions-SessionServiceImpl
// class:
// alias:alias-85899ef785c3f033c4d8293618ff61de-Service
// scope:singleton
//
type p0d8344d2dc_isessions_SessionServiceImpl struct {
}

func (inst* p0d8344d2dc_isessions_SessionServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-0d8344d2dcb25264-isessions-SessionServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-85899ef785c3f033c4d8293618ff61de-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p0d8344d2dc_isessions_SessionServiceImpl) new() any {
    return &p0d8344d2d.SessionServiceImpl{}
}

func (inst* p0d8344d2dc_isessions_SessionServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p0d8344d2d.SessionServiceImpl)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*p0d8344d2dc_isessions_SessionServiceImpl) getDao(ie application.InjectionExt)p85899ef78.DAO{
    return ie.GetComponent("#alias-85899ef785c3f033c4d8293618ff61de-DAO").(p85899ef78.DAO)
}



// type p29a9720d8.SubjectDataGroup in package:github.com/starter-go/v0/subjects/core/implements/isubjectdb
//
// id:com-29a9720d88726dd7-isubjectdb-SubjectDataGroup
// class:class-512a309140d0ad99eb1c95c8dc0d02f9-GroupRegistry
// alias:
// scope:singleton
//
type p29a9720d88_isubjectdb_SubjectDataGroup struct {
}

func (inst* p29a9720d88_isubjectdb_SubjectDataGroup) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-29a9720d88726dd7-isubjectdb-SubjectDataGroup"
	r.Classes = "class-512a309140d0ad99eb1c95c8dc0d02f9-GroupRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p29a9720d88_isubjectdb_SubjectDataGroup) new() any {
    return &p29a9720d8.SubjectDataGroup{}
}

func (inst* p29a9720d88_isubjectdb_SubjectDataGroup) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p29a9720d8.SubjectDataGroup)
	nop(ie, com)

	
    com.Alias = inst.getAlias(ie)
    com.Enabled = inst.getEnabled(ie)
    com.Prefix = inst.getPrefix(ie)
    com.Source = inst.getSource(ie)
    com.URI = inst.getURI(ie)


    return nil
}


func (inst*p29a9720d88_isubjectdb_SubjectDataGroup) getAlias(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.subjects.alias}")
}


func (inst*p29a9720d88_isubjectdb_SubjectDataGroup) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${datagroup.subjects.enabled}")
}


func (inst*p29a9720d88_isubjectdb_SubjectDataGroup) getPrefix(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.subjects.table-name-prefix}")
}


func (inst*p29a9720d88_isubjectdb_SubjectDataGroup) getSource(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.subjects.datasource}")
}


func (inst*p29a9720d88_isubjectdb_SubjectDataGroup) getURI(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.subjects.uri}")
}



// type p29a9720d8.SubjectDataBaseAgent in package:github.com/starter-go/v0/subjects/core/implements/isubjectdb
//
// id:com-29a9720d88726dd7-isubjectdb-SubjectDataBaseAgent
// class:
// alias:alias-ba9b83e768291a6dc6c6a9a0b0051833-Agent
// scope:singleton
//
type p29a9720d88_isubjectdb_SubjectDataBaseAgent struct {
}

func (inst* p29a9720d88_isubjectdb_SubjectDataBaseAgent) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-29a9720d88726dd7-isubjectdb-SubjectDataBaseAgent"
	r.Classes = ""
	r.Aliases = "alias-ba9b83e768291a6dc6c6a9a0b0051833-Agent"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p29a9720d88_isubjectdb_SubjectDataBaseAgent) new() any {
    return &p29a9720d8.SubjectDataBaseAgent{}
}

func (inst* p29a9720d88_isubjectdb_SubjectDataBaseAgent) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p29a9720d8.SubjectDataBaseAgent)
	nop(ie, com)

	
    com.DSM = inst.getDSM(ie)
    com.SourceName = inst.getSourceName(ie)


    return nil
}


func (inst*p29a9720d88_isubjectdb_SubjectDataBaseAgent) getDSM(ie application.InjectionExt)p512a30914.DataSourceManager{
    return ie.GetComponent("#alias-512a309140d0ad99eb1c95c8dc0d02f9-DataSourceManager").(p512a30914.DataSourceManager)
}


func (inst*p29a9720d88_isubjectdb_SubjectDataBaseAgent) getSourceName(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.subjects.datasource}")
}


