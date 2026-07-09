package main4libvlog
import (
    pdc2a427af "github.com/starter-go/v0/libvlog/lib/boot"
    p83b2f29a0 "github.com/starter-go/v0/libvlog/lib/filters"
    p55f0853be "github.com/starter-go/vlog"
     "github.com/starter-go/application"
)

// type pdc2a427af.VLogBootLoader in package:github.com/starter-go/v0/libvlog/lib/boot
//
// id:com-dc2a427af12e181a-boot-VLogBootLoader
// class:class-0ef6f2938681e99da4b0c19ce3d3fb4f-Lifecycle
// alias:
// scope:singleton
//
type pdc2a427af1_boot_VLogBootLoader struct {
}

func (inst* pdc2a427af1_boot_VLogBootLoader) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-dc2a427af12e181a-boot-VLogBootLoader"
	r.Classes = "class-0ef6f2938681e99da4b0c19ce3d3fb4f-Lifecycle"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pdc2a427af1_boot_VLogBootLoader) new() any {
    return &pdc2a427af.VLogBootLoader{}
}

func (inst* pdc2a427af1_boot_VLogBootLoader) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pdc2a427af.VLogBootLoader)
	nop(ie, com)

	
    com.FilterList = inst.getFilterList(ie)


    return nil
}


func (inst*pdc2a427af1_boot_VLogBootLoader) getFilterList(ie application.InjectionExt)[]p55f0853be.FilterRegistry{
    dst := make([]p55f0853be.FilterRegistry, 0)
    src := ie.ListComponents(".class-55f0853bedbc094981acd8da904ae269-FilterRegistry")
    for _, item1 := range src {
        item2 := item1.(p55f0853be.FilterRegistry)
        dst = append(dst, item2)
    }
    return dst
}



// type p83b2f29a0.ConsoleWriterFilter in package:github.com/starter-go/v0/libvlog/lib/filters
//
// id:com-83b2f29a09be0e1c-filters-ConsoleWriterFilter
// class:class-55f0853bedbc094981acd8da904ae269-FilterRegistry
// alias:
// scope:singleton
//
type p83b2f29a09_filters_ConsoleWriterFilter struct {
}

func (inst* p83b2f29a09_filters_ConsoleWriterFilter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-83b2f29a09be0e1c-filters-ConsoleWriterFilter"
	r.Classes = "class-55f0853bedbc094981acd8da904ae269-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p83b2f29a09_filters_ConsoleWriterFilter) new() any {
    return &p83b2f29a0.ConsoleWriterFilter{}
}

func (inst* p83b2f29a09_filters_ConsoleWriterFilter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p83b2f29a0.ConsoleWriterFilter)
	nop(ie, com)

	


    return nil
}



// type p83b2f29a0.DefaultValueFilter in package:github.com/starter-go/v0/libvlog/lib/filters
//
// id:com-83b2f29a09be0e1c-filters-DefaultValueFilter
// class:class-55f0853bedbc094981acd8da904ae269-FilterRegistry
// alias:
// scope:singleton
//
type p83b2f29a09_filters_DefaultValueFilter struct {
}

func (inst* p83b2f29a09_filters_DefaultValueFilter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-83b2f29a09be0e1c-filters-DefaultValueFilter"
	r.Classes = "class-55f0853bedbc094981acd8da904ae269-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p83b2f29a09_filters_DefaultValueFilter) new() any {
    return &p83b2f29a0.DefaultValueFilter{}
}

func (inst* p83b2f29a09_filters_DefaultValueFilter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p83b2f29a0.DefaultValueFilter)
	nop(ie, com)

	


    return nil
}



// type p83b2f29a0.FormatterFilter in package:github.com/starter-go/v0/libvlog/lib/filters
//
// id:com-83b2f29a09be0e1c-filters-FormatterFilter
// class:class-55f0853bedbc094981acd8da904ae269-FilterRegistry
// alias:
// scope:singleton
//
type p83b2f29a09_filters_FormatterFilter struct {
}

func (inst* p83b2f29a09_filters_FormatterFilter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-83b2f29a09be0e1c-filters-FormatterFilter"
	r.Classes = "class-55f0853bedbc094981acd8da904ae269-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p83b2f29a09_filters_FormatterFilter) new() any {
    return &p83b2f29a0.FormatterFilter{}
}

func (inst* p83b2f29a09_filters_FormatterFilter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p83b2f29a0.FormatterFilter)
	nop(ie, com)

	
    com.HeadFormat = inst.getHeadFormat(ie)


    return nil
}


func (inst*p83b2f29a09_filters_FormatterFilter) getHeadFormat(ie application.InjectionExt)string{
    return ie.GetString("${vlog.formatters.default.format}")
}



// type p83b2f29a0.LevelFilter in package:github.com/starter-go/v0/libvlog/lib/filters
//
// id:com-83b2f29a09be0e1c-filters-LevelFilter
// class:class-55f0853bedbc094981acd8da904ae269-FilterRegistry
// alias:
// scope:singleton
//
type p83b2f29a09_filters_LevelFilter struct {
}

func (inst* p83b2f29a09_filters_LevelFilter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-83b2f29a09be0e1c-filters-LevelFilter"
	r.Classes = "class-55f0853bedbc094981acd8da904ae269-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p83b2f29a09_filters_LevelFilter) new() any {
    return &p83b2f29a0.LevelFilter{}
}

func (inst* p83b2f29a09_filters_LevelFilter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p83b2f29a0.LevelFilter)
	nop(ie, com)

	
    com.Level = inst.getLevel(ie)


    return nil
}


func (inst*p83b2f29a09_filters_LevelFilter) getLevel(ie application.InjectionExt)string{
    return ie.GetString("${vlog.level}")
}



// type p83b2f29a0.TimeFilter in package:github.com/starter-go/v0/libvlog/lib/filters
//
// id:com-83b2f29a09be0e1c-filters-TimeFilter
// class:class-55f0853bedbc094981acd8da904ae269-FilterRegistry
// alias:
// scope:singleton
//
type p83b2f29a09_filters_TimeFilter struct {
}

func (inst* p83b2f29a09_filters_TimeFilter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-83b2f29a09be0e1c-filters-TimeFilter"
	r.Classes = "class-55f0853bedbc094981acd8da904ae269-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p83b2f29a09_filters_TimeFilter) new() any {
    return &p83b2f29a0.TimeFilter{}
}

func (inst* p83b2f29a09_filters_TimeFilter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p83b2f29a0.TimeFilter)
	nop(ie, com)

	


    return nil
}


