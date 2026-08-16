package main4keys
import (
    p75e208792 "github.com/starter-go/v0/keys"
    p15a2a0e8c "github.com/starter-go/v0/keys/lib/libkeys"
     "github.com/starter-go/application"
)

// type p15a2a0e8c.LibKeysService in package:github.com/starter-go/v0/keys/lib/libkeys
//
// id:com-15a2a0e8c2c8eb67-libkeys-LibKeysService
// class:
// alias:alias-75e2087923338087ed744e7a6494638b-Service
// scope:singleton
//
type p15a2a0e8c2_libkeys_LibKeysService struct {
}

func (inst* p15a2a0e8c2_libkeys_LibKeysService) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-15a2a0e8c2c8eb67-libkeys-LibKeysService"
	r.Classes = ""
	r.Aliases = "alias-75e2087923338087ed744e7a6494638b-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p15a2a0e8c2_libkeys_LibKeysService) new() any {
    return &p15a2a0e8c.LibKeysService{}
}

func (inst* p15a2a0e8c2_libkeys_LibKeysService) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p15a2a0e8c.LibKeysService)
	nop(ie, com)

	


    return nil
}



// type p15a2a0e8c.Loader in package:github.com/starter-go/v0/keys/lib/libkeys
//
// id:com-15a2a0e8c2c8eb67-libkeys-Loader
// class:
// alias:alias-75e2087923338087ed744e7a6494638b-ServiceLoader
// scope:singleton
//
type p15a2a0e8c2_libkeys_Loader struct {
}

func (inst* p15a2a0e8c2_libkeys_Loader) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-15a2a0e8c2c8eb67-libkeys-Loader"
	r.Classes = ""
	r.Aliases = "alias-75e2087923338087ed744e7a6494638b-ServiceLoader"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p15a2a0e8c2_libkeys_Loader) new() any {
    return &p15a2a0e8c.Loader{}
}

func (inst* p15a2a0e8c2_libkeys_Loader) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p15a2a0e8c.Loader)
	nop(ie, com)

	
    com.Ser = inst.getSer(ie)


    return nil
}


func (inst*p15a2a0e8c2_libkeys_Loader) getSer(ie application.InjectionExt)p75e208792.Service{
    return ie.GetComponent("#alias-75e2087923338087ed744e7a6494638b-Service").(p75e208792.Service)
}


