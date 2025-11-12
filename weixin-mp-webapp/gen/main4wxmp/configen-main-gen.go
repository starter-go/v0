package main4wxmp

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

    
    inst.register(&p0051ce0c69_iweixinmessages_MainWeixinMessageHandler{})
    inst.register(&p21cbfb264c_idatabases_MyDataGroup{})
    inst.register(&p21cbfb264c_idatabases_MyDatabaseAgentImpl{})
    inst.register(&p8701997662_controllers_DebugMonitorController{})
    inst.register(&p8701997662_controllers_ExampleController{})
    inst.register(&p8701997662_controllers_WeixinPlatformMessageController{})


    return nil
}
