package main4libvlog

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

    
    inst.register(&p83b2f29a09_filters_ConsoleWriterFilter{})
    inst.register(&p83b2f29a09_filters_DefaultValueFilter{})
    inst.register(&p83b2f29a09_filters_FormatterFilter{})
    inst.register(&p83b2f29a09_filters_LevelFilter{})
    inst.register(&p83b2f29a09_filters_TimeFilter{})
    inst.register(&pdc2a427af1_boot_VLogBootLoader{})


    return nil
}
