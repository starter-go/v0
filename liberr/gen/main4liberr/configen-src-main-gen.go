package main4liberr
import (
    p52d293b71 "github.com/starter-go/i18n"
    p430c7b94c "github.com/starter-go/v0/liberr"
    p0059025e6 "github.com/starter-go/v0/liberr/core"
     "github.com/starter-go/application"
)

// type p0059025e6.ErrorManagerImpl in package:github.com/starter-go/v0/liberr/core
//
// id:com-0059025e6cc2f67a-core-ErrorManagerImpl
// class:
// alias:alias-430c7b94c7c73bc48c254d5d9d9f6a6c-ErrorManager
// scope:singleton
//
type p0059025e6c_core_ErrorManagerImpl struct {
}

func (inst* p0059025e6c_core_ErrorManagerImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-0059025e6cc2f67a-core-ErrorManagerImpl"
	r.Classes = ""
	r.Aliases = "alias-430c7b94c7c73bc48c254d5d9d9f6a6c-ErrorManager"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p0059025e6c_core_ErrorManagerImpl) new() any {
    return &p0059025e6.ErrorManagerImpl{}
}

func (inst* p0059025e6c_core_ErrorManagerImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p0059025e6.ErrorManagerImpl)
	nop(ie, com)

	
    com.PropLogOnStart = inst.getPropLogOnStart(ie)
    com.PropReCode = inst.getPropReCode(ie)
    com.PropCodeMask = inst.getPropCodeMask(ie)
    com.ErrSets = inst.getErrSets(ie)


    return nil
}


func (inst*p0059025e6c_core_ErrorManagerImpl) getPropLogOnStart(ie application.InjectionExt)bool{
    return ie.GetBool("${liberr.log-on-start}")
}


func (inst*p0059025e6c_core_ErrorManagerImpl) getPropReCode(ie application.InjectionExt)bool{
    return ie.GetBool("${liberr.re-code}")
}


func (inst*p0059025e6c_core_ErrorManagerImpl) getPropCodeMask(ie application.InjectionExt)int{
    return ie.GetInt("${liberr.code-mask}")
}


func (inst*p0059025e6c_core_ErrorManagerImpl) getErrSets(ie application.InjectionExt)[]p430c7b94c.ErrorSet{
    dst := make([]p430c7b94c.ErrorSet, 0)
    src := ie.ListComponents(".class-430c7b94c7c73bc48c254d5d9d9f6a6c-ErrorSet")
    for _, item1 := range src {
        item2 := item1.(p430c7b94c.ErrorSet)
        dst = append(dst, item2)
    }
    return dst
}



// type p0059025e6.ErrorServiceImpl in package:github.com/starter-go/v0/liberr/core
//
// id:com-0059025e6cc2f67a-core-ErrorServiceImpl
// class:
// alias:alias-430c7b94c7c73bc48c254d5d9d9f6a6c-Service
// scope:singleton
//
type p0059025e6c_core_ErrorServiceImpl struct {
}

func (inst* p0059025e6c_core_ErrorServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-0059025e6cc2f67a-core-ErrorServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-430c7b94c7c73bc48c254d5d9d9f6a6c-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p0059025e6c_core_ErrorServiceImpl) new() any {
    return &p0059025e6.ErrorServiceImpl{}
}

func (inst* p0059025e6c_core_ErrorServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p0059025e6.ErrorServiceImpl)
	nop(ie, com)

	
    com.Man = inst.getMan(ie)
    com.I18nser = inst.getI18nser(ie)


    return nil
}


func (inst*p0059025e6c_core_ErrorServiceImpl) getMan(ie application.InjectionExt)p430c7b94c.ErrorManager{
    return ie.GetComponent("#alias-430c7b94c7c73bc48c254d5d9d9f6a6c-ErrorManager").(p430c7b94c.ErrorManager)
}


func (inst*p0059025e6c_core_ErrorServiceImpl) getI18nser(ie application.InjectionExt)p52d293b71.Service{
    return ie.GetComponent("#alias-52d293b7197ae4adfab092ade2e718a2-Service").(p52d293b71.Service)
}


