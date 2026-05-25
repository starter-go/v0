package main4htttest

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

    
    inst.register(&p8753df4774_core_CodecManagerImpl{})
    inst.register(&p8753df4774_core_ServiceImpl{})
    inst.register(&pbe5934bc9d_codecs_BinaryCodec{})
    inst.register(&pbe5934bc9d_codecs_ExampleCodec{})
    inst.register(&pbe5934bc9d_codecs_HtmlCodec{})
    inst.register(&pbe5934bc9d_codecs_JsonCodec{})
    inst.register(&pbe5934bc9d_codecs_PlainTextCodec{})


    return nil
}
