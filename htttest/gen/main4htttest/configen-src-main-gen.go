package main4htttest
import (
    p581c1dbb5 "github.com/starter-go/v0/htttest"
    p8753df477 "github.com/starter-go/v0/htttest/core"
    pbe5934bc9 "github.com/starter-go/v0/htttest/core/codecs"
     "github.com/starter-go/application"
)

// type p8753df477.CodecManagerImpl in package:github.com/starter-go/v0/htttest/core
//
// id:com-8753df477433f937-core-CodecManagerImpl
// class:
// alias:alias-581c1dbb530beaf64ae39cc43b21dea9-CodecManager
// scope:singleton
//
type p8753df4774_core_CodecManagerImpl struct {
}

func (inst* p8753df4774_core_CodecManagerImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-8753df477433f937-core-CodecManagerImpl"
	r.Classes = ""
	r.Aliases = "alias-581c1dbb530beaf64ae39cc43b21dea9-CodecManager"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p8753df4774_core_CodecManagerImpl) new() any {
    return &p8753df477.CodecManagerImpl{}
}

func (inst* p8753df4774_core_CodecManagerImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p8753df477.CodecManagerImpl)
	nop(ie, com)

	
    com.RawProviderList = inst.getRawProviderList(ie)


    return nil
}


func (inst*p8753df4774_core_CodecManagerImpl) getRawProviderList(ie application.InjectionExt)[]p581c1dbb5.CodecProvider{
    dst := make([]p581c1dbb5.CodecProvider, 0)
    src := ie.ListComponents(".class-581c1dbb530beaf64ae39cc43b21dea9-CodecProvider")
    for _, item1 := range src {
        item2 := item1.(p581c1dbb5.CodecProvider)
        dst = append(dst, item2)
    }
    return dst
}



// type p8753df477.ServiceImpl in package:github.com/starter-go/v0/htttest/core
//
// id:com-8753df477433f937-core-ServiceImpl
// class:
// alias:alias-581c1dbb530beaf64ae39cc43b21dea9-Service
// scope:singleton
//
type p8753df4774_core_ServiceImpl struct {
}

func (inst* p8753df4774_core_ServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-8753df477433f937-core-ServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-581c1dbb530beaf64ae39cc43b21dea9-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p8753df4774_core_ServiceImpl) new() any {
    return &p8753df477.ServiceImpl{}
}

func (inst* p8753df4774_core_ServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p8753df477.ServiceImpl)
	nop(ie, com)

	
    com.BaseURL = inst.getBaseURL(ie)
    com.MaxContentLength = inst.getMaxContentLength(ie)
    com.CodecManager = inst.getCodecManager(ie)


    return nil
}


func (inst*p8753df4774_core_ServiceImpl) getBaseURL(ie application.InjectionExt)string{
    return ie.GetString("${starter.htttest.base-url}")
}


func (inst*p8753df4774_core_ServiceImpl) getMaxContentLength(ie application.InjectionExt)int{
    return ie.GetInt("${starter.htttest.max-content-length}")
}


func (inst*p8753df4774_core_ServiceImpl) getCodecManager(ie application.InjectionExt)p581c1dbb5.CodecManager{
    return ie.GetComponent("#alias-581c1dbb530beaf64ae39cc43b21dea9-CodecManager").(p581c1dbb5.CodecManager)
}



// type pbe5934bc9.BinaryCodec in package:github.com/starter-go/v0/htttest/core/codecs
//
// id:com-be5934bc9df3e530-codecs-BinaryCodec
// class:class-581c1dbb530beaf64ae39cc43b21dea9-CodecProvider
// alias:
// scope:singleton
//
type pbe5934bc9d_codecs_BinaryCodec struct {
}

func (inst* pbe5934bc9d_codecs_BinaryCodec) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-be5934bc9df3e530-codecs-BinaryCodec"
	r.Classes = "class-581c1dbb530beaf64ae39cc43b21dea9-CodecProvider"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pbe5934bc9d_codecs_BinaryCodec) new() any {
    return &pbe5934bc9.BinaryCodec{}
}

func (inst* pbe5934bc9d_codecs_BinaryCodec) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pbe5934bc9.BinaryCodec)
	nop(ie, com)

	


    return nil
}



// type pbe5934bc9.ExampleCodec in package:github.com/starter-go/v0/htttest/core/codecs
//
// id:com-be5934bc9df3e530-codecs-ExampleCodec
// class:class-581c1dbb530beaf64ae39cc43b21dea9-CodecProvider
// alias:
// scope:singleton
//
type pbe5934bc9d_codecs_ExampleCodec struct {
}

func (inst* pbe5934bc9d_codecs_ExampleCodec) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-be5934bc9df3e530-codecs-ExampleCodec"
	r.Classes = "class-581c1dbb530beaf64ae39cc43b21dea9-CodecProvider"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pbe5934bc9d_codecs_ExampleCodec) new() any {
    return &pbe5934bc9.ExampleCodec{}
}

func (inst* pbe5934bc9d_codecs_ExampleCodec) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pbe5934bc9.ExampleCodec)
	nop(ie, com)

	


    return nil
}



// type pbe5934bc9.HtmlCodec in package:github.com/starter-go/v0/htttest/core/codecs
//
// id:com-be5934bc9df3e530-codecs-HtmlCodec
// class:class-581c1dbb530beaf64ae39cc43b21dea9-CodecProvider
// alias:
// scope:singleton
//
type pbe5934bc9d_codecs_HtmlCodec struct {
}

func (inst* pbe5934bc9d_codecs_HtmlCodec) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-be5934bc9df3e530-codecs-HtmlCodec"
	r.Classes = "class-581c1dbb530beaf64ae39cc43b21dea9-CodecProvider"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pbe5934bc9d_codecs_HtmlCodec) new() any {
    return &pbe5934bc9.HtmlCodec{}
}

func (inst* pbe5934bc9d_codecs_HtmlCodec) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pbe5934bc9.HtmlCodec)
	nop(ie, com)

	


    return nil
}



// type pbe5934bc9.JsonCodec in package:github.com/starter-go/v0/htttest/core/codecs
//
// id:com-be5934bc9df3e530-codecs-JsonCodec
// class:class-581c1dbb530beaf64ae39cc43b21dea9-CodecProvider
// alias:
// scope:singleton
//
type pbe5934bc9d_codecs_JsonCodec struct {
}

func (inst* pbe5934bc9d_codecs_JsonCodec) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-be5934bc9df3e530-codecs-JsonCodec"
	r.Classes = "class-581c1dbb530beaf64ae39cc43b21dea9-CodecProvider"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pbe5934bc9d_codecs_JsonCodec) new() any {
    return &pbe5934bc9.JsonCodec{}
}

func (inst* pbe5934bc9d_codecs_JsonCodec) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pbe5934bc9.JsonCodec)
	nop(ie, com)

	


    return nil
}



// type pbe5934bc9.PlainTextCodec in package:github.com/starter-go/v0/htttest/core/codecs
//
// id:com-be5934bc9df3e530-codecs-PlainTextCodec
// class:class-581c1dbb530beaf64ae39cc43b21dea9-CodecProvider
// alias:
// scope:singleton
//
type pbe5934bc9d_codecs_PlainTextCodec struct {
}

func (inst* pbe5934bc9d_codecs_PlainTextCodec) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-be5934bc9df3e530-codecs-PlainTextCodec"
	r.Classes = "class-581c1dbb530beaf64ae39cc43b21dea9-CodecProvider"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pbe5934bc9d_codecs_PlainTextCodec) new() any {
    return &pbe5934bc9.PlainTextCodec{}
}

func (inst* pbe5934bc9d_codecs_PlainTextCodec) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pbe5934bc9.PlainTextCodec)
	nop(ie, com)

	


    return nil
}


