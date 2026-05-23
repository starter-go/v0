package main4libjwt
import (
    p40146ff25 "github.com/starter-go/v0/libjwt/api/jwt"
    pdd265c356 "github.com/starter-go/v0/libjwt/core/iadapter"
    pd2ae4a4ea "github.com/starter-go/v0/libjwt/core/icodec"
    p80413be4d "github.com/starter-go/v0/libjwt/core/ikeyloader"
    p6f45f395e "github.com/starter-go/v0/libjwt/core/iservice"
     "github.com/starter-go/application"
)

// type pdd265c356.MainAdapter in package:github.com/starter-go/v0/libjwt/core/iadapter
//
// id:com-dd265c35606f28b4-iadapter-MainAdapter
// class:
// alias:alias-40146ff2506e686679e1be8b766901e1-Adapter
// scope:singleton
//
type pdd265c3560_iadapter_MainAdapter struct {
}

func (inst* pdd265c3560_iadapter_MainAdapter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-dd265c35606f28b4-iadapter-MainAdapter"
	r.Classes = ""
	r.Aliases = "alias-40146ff2506e686679e1be8b766901e1-Adapter"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pdd265c3560_iadapter_MainAdapter) new() any {
    return &pdd265c356.MainAdapter{}
}

func (inst* pdd265c3560_iadapter_MainAdapter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pdd265c356.MainAdapter)
	nop(ie, com)

	
    com.ProviderList = inst.getProviderList(ie)
    com.AutoTermEnabled = inst.getAutoTermEnabled(ie)
    com.AutoTermAge = inst.getAutoTermAge(ie)


    return nil
}


func (inst*pdd265c3560_iadapter_MainAdapter) getProviderList(ie application.InjectionExt)[]p40146ff25.AdapterProvider{
    dst := make([]p40146ff25.AdapterProvider, 0)
    src := ie.ListComponents(".class-40146ff2506e686679e1be8b766901e1-AdapterProvider")
    for _, item1 := range src {
        item2 := item1.(p40146ff25.AdapterProvider)
        dst = append(dst, item2)
    }
    return dst
}


func (inst*pdd265c3560_iadapter_MainAdapter) getAutoTermEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${jwt.auto-term.enabled}")
}


func (inst*pdd265c3560_iadapter_MainAdapter) getAutoTermAge(ie application.InjectionExt)int64{
    return ie.GetInt64("${jwt.auto-term.age-in-ms}")
}



// type pdd265c356.MockAdapter in package:github.com/starter-go/v0/libjwt/core/iadapter
//
// id:com-dd265c35606f28b4-iadapter-MockAdapter
// class:class-40146ff2506e686679e1be8b766901e1-AdapterProvider
// alias:
// scope:singleton
//
type pdd265c3560_iadapter_MockAdapter struct {
}

func (inst* pdd265c3560_iadapter_MockAdapter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-dd265c35606f28b4-iadapter-MockAdapter"
	r.Classes = "class-40146ff2506e686679e1be8b766901e1-AdapterProvider"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pdd265c3560_iadapter_MockAdapter) new() any {
    return &pdd265c356.MockAdapter{}
}

func (inst* pdd265c3560_iadapter_MockAdapter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pdd265c356.MockAdapter)
	nop(ie, com)

	
    com.MyName = inst.getMyName(ie)
    com.MyPriority = inst.getMyPriority(ie)
    com.MyEnabled = inst.getMyEnabled(ie)


    return nil
}


func (inst*pdd265c3560_iadapter_MockAdapter) getMyName(ie application.InjectionExt)string{
    return ie.GetString("${jwt-adapter.mock.name}")
}


func (inst*pdd265c3560_iadapter_MockAdapter) getMyPriority(ie application.InjectionExt)int{
    return ie.GetInt("${jwt-adapter.mock.priority}")
}


func (inst*pdd265c3560_iadapter_MockAdapter) getMyEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${jwt-adapter.mock.enabled}")
}



// type pd2ae4a4ea.DefaultCodec in package:github.com/starter-go/v0/libjwt/core/icodec
//
// id:com-d2ae4a4ea8269add-icodec-DefaultCodec
// class:class-40146ff2506e686679e1be8b766901e1-CodecProvider
// alias:
// scope:singleton
//
type pd2ae4a4ea8_icodec_DefaultCodec struct {
}

func (inst* pd2ae4a4ea8_icodec_DefaultCodec) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-d2ae4a4ea8269add-icodec-DefaultCodec"
	r.Classes = "class-40146ff2506e686679e1be8b766901e1-CodecProvider"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pd2ae4a4ea8_icodec_DefaultCodec) new() any {
    return &pd2ae4a4ea.DefaultCodec{}
}

func (inst* pd2ae4a4ea8_icodec_DefaultCodec) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pd2ae4a4ea.DefaultCodec)
	nop(ie, com)

	
    com.Name = inst.getName(ie)
    com.Enabled = inst.getEnabled(ie)
    com.Priority = inst.getPriority(ie)
    com.KeyLoader = inst.getKeyLoader(ie)


    return nil
}


func (inst*pd2ae4a4ea8_icodec_DefaultCodec) getName(ie application.InjectionExt)string{
    return ie.GetString("${jwt-codec.default.name}")
}


func (inst*pd2ae4a4ea8_icodec_DefaultCodec) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${jwt-codec.default.enabled}")
}


func (inst*pd2ae4a4ea8_icodec_DefaultCodec) getPriority(ie application.InjectionExt)int{
    return ie.GetInt("${jwt-codec.default.priority}")
}


func (inst*pd2ae4a4ea8_icodec_DefaultCodec) getKeyLoader(ie application.InjectionExt)p40146ff25.KeyLoader{
    return ie.GetComponent("#alias-40146ff2506e686679e1be8b766901e1-KeyLoader").(p40146ff25.KeyLoader)
}



// type pd2ae4a4ea.MainCodec in package:github.com/starter-go/v0/libjwt/core/icodec
//
// id:com-d2ae4a4ea8269add-icodec-MainCodec
// class:
// alias:alias-40146ff2506e686679e1be8b766901e1-CODEC
// scope:singleton
//
type pd2ae4a4ea8_icodec_MainCodec struct {
}

func (inst* pd2ae4a4ea8_icodec_MainCodec) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-d2ae4a4ea8269add-icodec-MainCodec"
	r.Classes = ""
	r.Aliases = "alias-40146ff2506e686679e1be8b766901e1-CODEC"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pd2ae4a4ea8_icodec_MainCodec) new() any {
    return &pd2ae4a4ea.MainCodec{}
}

func (inst* pd2ae4a4ea8_icodec_MainCodec) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pd2ae4a4ea.MainCodec)
	nop(ie, com)

	
    com.ProviderList = inst.getProviderList(ie)


    return nil
}


func (inst*pd2ae4a4ea8_icodec_MainCodec) getProviderList(ie application.InjectionExt)[]p40146ff25.CodecProvider{
    dst := make([]p40146ff25.CodecProvider, 0)
    src := ie.ListComponents(".class-40146ff2506e686679e1be8b766901e1-CodecProvider")
    for _, item1 := range src {
        item2 := item1.(p40146ff25.CodecProvider)
        dst = append(dst, item2)
    }
    return dst
}



// type p80413be4d.MainKeyLoader in package:github.com/starter-go/v0/libjwt/core/ikeyloader
//
// id:com-80413be4da841351-ikeyloader-MainKeyLoader
// class:
// alias:alias-40146ff2506e686679e1be8b766901e1-KeyLoader
// scope:singleton
//
type p80413be4da_ikeyloader_MainKeyLoader struct {
}

func (inst* p80413be4da_ikeyloader_MainKeyLoader) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-80413be4da841351-ikeyloader-MainKeyLoader"
	r.Classes = ""
	r.Aliases = "alias-40146ff2506e686679e1be8b766901e1-KeyLoader"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p80413be4da_ikeyloader_MainKeyLoader) new() any {
    return &p80413be4d.MainKeyLoader{}
}

func (inst* p80413be4da_ikeyloader_MainKeyLoader) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p80413be4d.MainKeyLoader)
	nop(ie, com)

	
    com.ProviderList = inst.getProviderList(ie)
    com.LengthInBytes = inst.getLengthInBytes(ie)


    return nil
}


func (inst*p80413be4da_ikeyloader_MainKeyLoader) getProviderList(ie application.InjectionExt)[]p40146ff25.KeyLoaderProvider{
    dst := make([]p40146ff25.KeyLoaderProvider, 0)
    src := ie.ListComponents(".class-40146ff2506e686679e1be8b766901e1-KeyLoaderProvider")
    for _, item1 := range src {
        item2 := item1.(p40146ff25.KeyLoaderProvider)
        dst = append(dst, item2)
    }
    return dst
}


func (inst*p80413be4da_ikeyloader_MainKeyLoader) getLengthInBytes(ie application.InjectionExt)int{
    return ie.GetInt("${jwt-key-loader.main.length}")
}



// type p80413be4d.PropertyKeyLoader in package:github.com/starter-go/v0/libjwt/core/ikeyloader
//
// id:com-80413be4da841351-ikeyloader-PropertyKeyLoader
// class:class-40146ff2506e686679e1be8b766901e1-KeyLoaderProvider
// alias:
// scope:singleton
//
type p80413be4da_ikeyloader_PropertyKeyLoader struct {
}

func (inst* p80413be4da_ikeyloader_PropertyKeyLoader) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-80413be4da841351-ikeyloader-PropertyKeyLoader"
	r.Classes = "class-40146ff2506e686679e1be8b766901e1-KeyLoaderProvider"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p80413be4da_ikeyloader_PropertyKeyLoader) new() any {
    return &p80413be4d.PropertyKeyLoader{}
}

func (inst* p80413be4da_ikeyloader_PropertyKeyLoader) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p80413be4d.PropertyKeyLoader)
	nop(ie, com)

	
    com.MyEnabled = inst.getMyEnabled(ie)
    com.MyName = inst.getMyName(ie)
    com.MyPriority = inst.getMyPriority(ie)
    com.MySecretData = inst.getMySecretData(ie)


    return nil
}


func (inst*p80413be4da_ikeyloader_PropertyKeyLoader) getMyEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${jwt-key-loader.property.enabled}")
}


func (inst*p80413be4da_ikeyloader_PropertyKeyLoader) getMyName(ie application.InjectionExt)string{
    return ie.GetString("${jwt-key-loader.property.name}")
}


func (inst*p80413be4da_ikeyloader_PropertyKeyLoader) getMyPriority(ie application.InjectionExt)int{
    return ie.GetInt("${jwt-key-loader.property.priority}")
}


func (inst*p80413be4da_ikeyloader_PropertyKeyLoader) getMySecretData(ie application.InjectionExt)string{
    return ie.GetString("${jwt-key-loader.property.secret}")
}



// type p6f45f395e.JwtServiceImpl in package:github.com/starter-go/v0/libjwt/core/iservice
//
// id:com-6f45f395ea260fc1-iservice-JwtServiceImpl
// class:
// alias:alias-40146ff2506e686679e1be8b766901e1-Service
// scope:singleton
//
type p6f45f395ea_iservice_JwtServiceImpl struct {
}

func (inst* p6f45f395ea_iservice_JwtServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-6f45f395ea260fc1-iservice-JwtServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-40146ff2506e686679e1be8b766901e1-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p6f45f395ea_iservice_JwtServiceImpl) new() any {
    return &p6f45f395e.JwtServiceImpl{}
}

func (inst* p6f45f395ea_iservice_JwtServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p6f45f395e.JwtServiceImpl)
	nop(ie, com)

	
    com.Adapter = inst.getAdapter(ie)
    com.Codec = inst.getCodec(ie)


    return nil
}


func (inst*p6f45f395ea_iservice_JwtServiceImpl) getAdapter(ie application.InjectionExt)p40146ff25.Adapter{
    return ie.GetComponent("#alias-40146ff2506e686679e1be8b766901e1-Adapter").(p40146ff25.Adapter)
}


func (inst*p6f45f395ea_iservice_JwtServiceImpl) getCodec(ie application.InjectionExt)p40146ff25.CODEC{
    return ie.GetComponent("#alias-40146ff2506e686679e1be8b766901e1-CODEC").(p40146ff25.CODEC)
}


