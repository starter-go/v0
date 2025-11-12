package main4wxmp
import (
    pd1a916a20 "github.com/starter-go/libgin"
    p512a30914 "github.com/starter-go/libgorm"
    p0d56fa662 "github.com/starter-go/v0/weixin-mp-webapp/app/classes/weixinmessages"
    p21cbfb264 "github.com/starter-go/v0/weixin-mp-webapp/app/implements/idatabases"
    p0051ce0c6 "github.com/starter-go/v0/weixin-mp-webapp/app/implements/iweixinmessages"
    p870199766 "github.com/starter-go/v0/weixin-mp-webapp/app/web/controllers"
     "github.com/starter-go/application"
)

// type p21cbfb264.MyDataGroup in package:github.com/starter-go/v0/weixin-mp-webapp/app/implements/idatabases
//
// id:com-21cbfb264c2f19cb-idatabases-MyDataGroup
// class:class-512a309140d0ad99eb1c95c8dc0d02f9-GroupRegistry
// alias:
// scope:singleton
//
type p21cbfb264c_idatabases_MyDataGroup struct {
}

func (inst* p21cbfb264c_idatabases_MyDataGroup) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-21cbfb264c2f19cb-idatabases-MyDataGroup"
	r.Classes = "class-512a309140d0ad99eb1c95c8dc0d02f9-GroupRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p21cbfb264c_idatabases_MyDataGroup) new() any {
    return &p21cbfb264.MyDataGroup{}
}

func (inst* p21cbfb264c_idatabases_MyDataGroup) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p21cbfb264.MyDataGroup)
	nop(ie, com)

	
    com.Alias = inst.getAlias(ie)
    com.URI = inst.getURI(ie)
    com.Prefix = inst.getPrefix(ie)
    com.Source = inst.getSource(ie)
    com.Enabled = inst.getEnabled(ie)


    return nil
}


func (inst*p21cbfb264c_idatabases_MyDataGroup) getAlias(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.weixin-mp.alias}")
}


func (inst*p21cbfb264c_idatabases_MyDataGroup) getURI(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.weixin-mp.uri}")
}


func (inst*p21cbfb264c_idatabases_MyDataGroup) getPrefix(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.weixin-mp.table-name-prefix}")
}


func (inst*p21cbfb264c_idatabases_MyDataGroup) getSource(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.weixin-mp.datasource}")
}


func (inst*p21cbfb264c_idatabases_MyDataGroup) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${datagroup.weixin-mp.enabled}")
}



// type p21cbfb264.MyDatabaseAgentImpl in package:github.com/starter-go/v0/weixin-mp-webapp/app/implements/idatabases
//
// id:com-21cbfb264c2f19cb-idatabases-MyDatabaseAgentImpl
// class:
// alias:alias-2b57da14cb7fe76d91e26740c215ea5f-Agent
// scope:singleton
//
type p21cbfb264c_idatabases_MyDatabaseAgentImpl struct {
}

func (inst* p21cbfb264c_idatabases_MyDatabaseAgentImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-21cbfb264c2f19cb-idatabases-MyDatabaseAgentImpl"
	r.Classes = ""
	r.Aliases = "alias-2b57da14cb7fe76d91e26740c215ea5f-Agent"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p21cbfb264c_idatabases_MyDatabaseAgentImpl) new() any {
    return &p21cbfb264.MyDatabaseAgentImpl{}
}

func (inst* p21cbfb264c_idatabases_MyDatabaseAgentImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p21cbfb264.MyDatabaseAgentImpl)
	nop(ie, com)

	
    com.DSM = inst.getDSM(ie)
    com.DSName = inst.getDSName(ie)


    return nil
}


func (inst*p21cbfb264c_idatabases_MyDatabaseAgentImpl) getDSM(ie application.InjectionExt)p512a30914.DataSourceManager{
    return ie.GetComponent("#alias-512a309140d0ad99eb1c95c8dc0d02f9-DataSourceManager").(p512a30914.DataSourceManager)
}


func (inst*p21cbfb264c_idatabases_MyDatabaseAgentImpl) getDSName(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.weixin-mp.datasource}")
}



// type p0051ce0c6.MainWeixinMessageHandler in package:github.com/starter-go/v0/weixin-mp-webapp/app/implements/iweixinmessages
//
// id:com-0051ce0c69b193de-iweixinmessages-MainWeixinMessageHandler
// class:
// alias:alias-0d56fa662bcf8ab7f739ee025f94262a-Handler
// scope:singleton
//
type p0051ce0c69_iweixinmessages_MainWeixinMessageHandler struct {
}

func (inst* p0051ce0c69_iweixinmessages_MainWeixinMessageHandler) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-0051ce0c69b193de-iweixinmessages-MainWeixinMessageHandler"
	r.Classes = ""
	r.Aliases = "alias-0d56fa662bcf8ab7f739ee025f94262a-Handler"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p0051ce0c69_iweixinmessages_MainWeixinMessageHandler) new() any {
    return &p0051ce0c6.MainWeixinMessageHandler{}
}

func (inst* p0051ce0c69_iweixinmessages_MainWeixinMessageHandler) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p0051ce0c6.MainWeixinMessageHandler)
	nop(ie, com)

	


    return nil
}



// type p870199766.DebugMonitorController in package:github.com/starter-go/v0/weixin-mp-webapp/app/web/controllers
//
// id:com-8701997662bfc211-controllers-DebugMonitorController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p8701997662_controllers_DebugMonitorController struct {
}

func (inst* p8701997662_controllers_DebugMonitorController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-8701997662bfc211-controllers-DebugMonitorController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p8701997662_controllers_DebugMonitorController) new() any {
    return &p870199766.DebugMonitorController{}
}

func (inst* p8701997662_controllers_DebugMonitorController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p870199766.DebugMonitorController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Enabled = inst.getEnabled(ie)
    com.IntervalInSec = inst.getIntervalInSec(ie)


    return nil
}


func (inst*p8701997662_controllers_DebugMonitorController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p8701997662_controllers_DebugMonitorController) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${debug.weixin-mp.monitor.enabled}")
}


func (inst*p8701997662_controllers_DebugMonitorController) getIntervalInSec(ie application.InjectionExt)int{
    return ie.GetInt("${debug.weixin-mp.monitor.interval-in-sec}")
}



// type p870199766.ExampleController in package:github.com/starter-go/v0/weixin-mp-webapp/app/web/controllers
//
// id:com-8701997662bfc211-controllers-ExampleController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p8701997662_controllers_ExampleController struct {
}

func (inst* p8701997662_controllers_ExampleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-8701997662bfc211-controllers-ExampleController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p8701997662_controllers_ExampleController) new() any {
    return &p870199766.ExampleController{}
}

func (inst* p8701997662_controllers_ExampleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p870199766.ExampleController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)


    return nil
}


func (inst*p8701997662_controllers_ExampleController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}



// type p870199766.WeixinPlatformMessageController in package:github.com/starter-go/v0/weixin-mp-webapp/app/web/controllers
//
// id:com-8701997662bfc211-controllers-WeixinPlatformMessageController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p8701997662_controllers_WeixinPlatformMessageController struct {
}

func (inst* p8701997662_controllers_WeixinPlatformMessageController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-8701997662bfc211-controllers-WeixinPlatformMessageController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p8701997662_controllers_WeixinPlatformMessageController) new() any {
    return &p870199766.WeixinPlatformMessageController{}
}

func (inst* p8701997662_controllers_WeixinPlatformMessageController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p870199766.WeixinPlatformMessageController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Handler = inst.getHandler(ie)
    com.PlatformMessageToken = inst.getPlatformMessageToken(ie)
    com.GatewayPath = inst.getGatewayPath(ie)
    com.IgnoreRequestSignature = inst.getIgnoreRequestSignature(ie)


    return nil
}


func (inst*p8701997662_controllers_WeixinPlatformMessageController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p8701997662_controllers_WeixinPlatformMessageController) getHandler(ie application.InjectionExt)p0d56fa662.Handler{
    return ie.GetComponent("#alias-0d56fa662bcf8ab7f739ee025f94262a-Handler").(p0d56fa662.Handler)
}


func (inst*p8701997662_controllers_WeixinPlatformMessageController) getPlatformMessageToken(ie application.InjectionExt)string{
    return ie.GetString("${weixin.mp-gateway.token}")
}


func (inst*p8701997662_controllers_WeixinPlatformMessageController) getGatewayPath(ie application.InjectionExt)string{
    return ie.GetString("${weixin.mp-gateway.path}")
}


func (inst*p8701997662_controllers_WeixinPlatformMessageController) getIgnoreRequestSignature(ie application.InjectionExt)bool{
    return ie.GetBool("${weixin.mp-gateway.ignore-request-signature}")
}


