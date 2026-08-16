package test4libdao
import (
    p8d3a42c23 "github.com/starter-go/v0/libdao/src/test/golang/testcom"
     "github.com/starter-go/application"
)

// type p8d3a42c23.UnitForDaoProxy in package:github.com/starter-go/v0/libdao/src/test/golang/testcom
//
// id:com-8d3a42c23bfb76cd-testcom-UnitForDaoProxy
// class:
// alias:
// scope:singleton
//
type p8d3a42c23b_testcom_UnitForDaoProxy struct {
}

func (inst* p8d3a42c23b_testcom_UnitForDaoProxy) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-8d3a42c23bfb76cd-testcom-UnitForDaoProxy"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p8d3a42c23b_testcom_UnitForDaoProxy) new() any {
    return &p8d3a42c23.UnitForDaoProxy{}
}

func (inst* p8d3a42c23b_testcom_UnitForDaoProxy) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p8d3a42c23.UnitForDaoProxy)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*p8d3a42c23b_testcom_UnitForDaoProxy) getDao(ie application.InjectionExt)p8d3a42c23.MockDao{
    return ie.GetComponent("#alias-8d3a42c23bfb76cd04476b1d889a7f1c-MockDao").(p8d3a42c23.MockDao)
}



// type p8d3a42c23.MockDaoMain in package:github.com/starter-go/v0/libdao/src/test/golang/testcom
//
// id:com-8d3a42c23bfb76cd-testcom-MockDaoMain
// class:
// alias:alias-8d3a42c23bfb76cd04476b1d889a7f1c-MockDao
// scope:singleton
//
type p8d3a42c23b_testcom_MockDaoMain struct {
}

func (inst* p8d3a42c23b_testcom_MockDaoMain) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-8d3a42c23bfb76cd-testcom-MockDaoMain"
	r.Classes = ""
	r.Aliases = "alias-8d3a42c23bfb76cd04476b1d889a7f1c-MockDao"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p8d3a42c23b_testcom_MockDaoMain) new() any {
    return &p8d3a42c23.MockDaoMain{}
}

func (inst* p8d3a42c23b_testcom_MockDaoMain) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p8d3a42c23.MockDaoMain)
	nop(ie, com)

	
    com.Selector = inst.getSelector(ie)
    com.DaoList = inst.getDaoList(ie)


    return nil
}


func (inst*p8d3a42c23b_testcom_MockDaoMain) getSelector(ie application.InjectionExt)string{
    return ie.GetString("${unit.mock-dao.selector}")
}


func (inst*p8d3a42c23b_testcom_MockDaoMain) getDaoList(ie application.InjectionExt)[]p8d3a42c23.MockDao{
    dst := make([]p8d3a42c23.MockDao, 0)
    src := ie.ListComponents(".class-8d3a42c23bfb76cd04476b1d889a7f1c-MockDao")
    for _, item1 := range src {
        item2 := item1.(p8d3a42c23.MockDao)
        dst = append(dst, item2)
    }
    return dst
}



// type p8d3a42c23.MockDao1 in package:github.com/starter-go/v0/libdao/src/test/golang/testcom
//
// id:com-8d3a42c23bfb76cd-testcom-MockDao1
// class:class-8d3a42c23bfb76cd04476b1d889a7f1c-MockDao
// alias:
// scope:singleton
//
type p8d3a42c23b_testcom_MockDao1 struct {
}

func (inst* p8d3a42c23b_testcom_MockDao1) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-8d3a42c23bfb76cd-testcom-MockDao1"
	r.Classes = "class-8d3a42c23bfb76cd04476b1d889a7f1c-MockDao"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p8d3a42c23b_testcom_MockDao1) new() any {
    return &p8d3a42c23.MockDao1{}
}

func (inst* p8d3a42c23b_testcom_MockDao1) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p8d3a42c23.MockDao1)
	nop(ie, com)

	


    return nil
}



// type p8d3a42c23.MockDao2 in package:github.com/starter-go/v0/libdao/src/test/golang/testcom
//
// id:com-8d3a42c23bfb76cd-testcom-MockDao2
// class:class-8d3a42c23bfb76cd04476b1d889a7f1c-MockDao
// alias:
// scope:singleton
//
type p8d3a42c23b_testcom_MockDao2 struct {
}

func (inst* p8d3a42c23b_testcom_MockDao2) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-8d3a42c23bfb76cd-testcom-MockDao2"
	r.Classes = "class-8d3a42c23bfb76cd04476b1d889a7f1c-MockDao"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p8d3a42c23b_testcom_MockDao2) new() any {
    return &p8d3a42c23.MockDao2{}
}

func (inst* p8d3a42c23b_testcom_MockDao2) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p8d3a42c23.MockDao2)
	nop(ie, com)

	


    return nil
}



// type p8d3a42c23.MockDao3 in package:github.com/starter-go/v0/libdao/src/test/golang/testcom
//
// id:com-8d3a42c23bfb76cd-testcom-MockDao3
// class:class-8d3a42c23bfb76cd04476b1d889a7f1c-MockDao
// alias:
// scope:singleton
//
type p8d3a42c23b_testcom_MockDao3 struct {
}

func (inst* p8d3a42c23b_testcom_MockDao3) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-8d3a42c23bfb76cd-testcom-MockDao3"
	r.Classes = "class-8d3a42c23bfb76cd04476b1d889a7f1c-MockDao"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p8d3a42c23b_testcom_MockDao3) new() any {
    return &p8d3a42c23.MockDao3{}
}

func (inst* p8d3a42c23b_testcom_MockDao3) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p8d3a42c23.MockDao3)
	nop(ie, com)

	


    return nil
}



// type p8d3a42c23.Example4t in package:github.com/starter-go/v0/libdao/src/test/golang/testcom
//
// id:com-8d3a42c23bfb76cd-testcom-Example4t
// class:
// alias:
// scope:singleton
//
type p8d3a42c23b_testcom_Example4t struct {
}

func (inst* p8d3a42c23b_testcom_Example4t) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-8d3a42c23bfb76cd-testcom-Example4t"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p8d3a42c23b_testcom_Example4t) new() any {
    return &p8d3a42c23.Example4t{}
}

func (inst* p8d3a42c23b_testcom_Example4t) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p8d3a42c23.Example4t)
	nop(ie, com)

	


    return nil
}


