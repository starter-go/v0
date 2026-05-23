package main4libjwt

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

    
    inst.register(&p6f45f395ea_iservice_JwtServiceImpl{})
    inst.register(&p80413be4da_ikeyloader_MainKeyLoader{})
    inst.register(&p80413be4da_ikeyloader_PropertyKeyLoader{})
    inst.register(&pd2ae4a4ea8_icodec_DefaultCodec{})
    inst.register(&pd2ae4a4ea8_icodec_MainCodec{})
    inst.register(&pdd265c3560_iadapter_MainAdapter{})
    inst.register(&pdd265c3560_iadapter_MockAdapter{})


    return nil
}
