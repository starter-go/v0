package test4keys
import (
    p75e208792 "github.com/starter-go/v0/keys"
    pacfbdceae "github.com/starter-go/v0/keys/src/test/golang/testcom"
     "github.com/starter-go/application"
)

// type pacfbdceae.DriversUnit in package:github.com/starter-go/v0/keys/src/test/golang/testcom
//
// id:com-acfbdceae66e0fbe-testcom-DriversUnit
// class:
// alias:
// scope:singleton
//
type pacfbdceae6_testcom_DriversUnit struct {
}

func (inst* pacfbdceae6_testcom_DriversUnit) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-acfbdceae66e0fbe-testcom-DriversUnit"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pacfbdceae6_testcom_DriversUnit) new() any {
    return &pacfbdceae.DriversUnit{}
}

func (inst* pacfbdceae6_testcom_DriversUnit) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pacfbdceae.DriversUnit)
	nop(ie, com)

	
    com.Ser = inst.getSer(ie)


    return nil
}


func (inst*pacfbdceae6_testcom_DriversUnit) getSer(ie application.InjectionExt)p75e208792.Service{
    return ie.GetComponent("#alias-75e2087923338087ed744e7a6494638b-Service").(p75e208792.Service)
}



// type pacfbdceae.ExampleUnit in package:github.com/starter-go/v0/keys/src/test/golang/testcom
//
// id:com-acfbdceae66e0fbe-testcom-ExampleUnit
// class:
// alias:
// scope:singleton
//
type pacfbdceae6_testcom_ExampleUnit struct {
}

func (inst* pacfbdceae6_testcom_ExampleUnit) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-acfbdceae66e0fbe-testcom-ExampleUnit"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pacfbdceae6_testcom_ExampleUnit) new() any {
    return &pacfbdceae.ExampleUnit{}
}

func (inst* pacfbdceae6_testcom_ExampleUnit) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pacfbdceae.ExampleUnit)
	nop(ie, com)

	


    return nil
}


