package test4rbacdg
import (
    p24287f458 "github.com/starter-go/rbac"
    p982dc5d72 "github.com/starter-go/v0/rbac-data-group/src/test/golang/testcom"
     "github.com/starter-go/application"
)

// type p982dc5d72.Example4t in package:github.com/starter-go/v0/rbac-data-group/src/test/golang/testcom
//
// id:com-982dc5d72ece47b7-testcom-Example4t
// class:
// alias:
// scope:singleton
//
type p982dc5d72e_testcom_Example4t struct {
}

func (inst* p982dc5d72e_testcom_Example4t) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-982dc5d72ece47b7-testcom-Example4t"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p982dc5d72e_testcom_Example4t) new() any {
    return &p982dc5d72.Example4t{}
}

func (inst* p982dc5d72e_testcom_Example4t) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p982dc5d72.Example4t)
	nop(ie, com)

	


    return nil
}



// type p982dc5d72.TestRbacDgCrudUnit in package:github.com/starter-go/v0/rbac-data-group/src/test/golang/testcom
//
// id:com-982dc5d72ece47b7-testcom-TestRbacDgCrudUnit
// class:
// alias:
// scope:singleton
//
type p982dc5d72e_testcom_TestRbacDgCrudUnit struct {
}

func (inst* p982dc5d72e_testcom_TestRbacDgCrudUnit) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-982dc5d72ece47b7-testcom-TestRbacDgCrudUnit"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p982dc5d72e_testcom_TestRbacDgCrudUnit) new() any {
    return &p982dc5d72.TestRbacDgCrudUnit{}
}

func (inst* p982dc5d72e_testcom_TestRbacDgCrudUnit) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p982dc5d72.TestRbacDgCrudUnit)
	nop(ie, com)

	
    com.PermDao = inst.getPermDao(ie)
    com.RoleDao = inst.getRoleDao(ie)
    com.SessionDao = inst.getSessionDao(ie)
    com.UserDao = inst.getUserDao(ie)


    return nil
}


func (inst*p982dc5d72e_testcom_TestRbacDgCrudUnit) getPermDao(ie application.InjectionExt)p24287f458.PermissionDAO{
    return ie.GetComponent("#alias-24287f4589fe5add27fb48a88d706565-PermissionDAO").(p24287f458.PermissionDAO)
}


func (inst*p982dc5d72e_testcom_TestRbacDgCrudUnit) getRoleDao(ie application.InjectionExt)p24287f458.RoleDAO{
    return ie.GetComponent("#alias-24287f4589fe5add27fb48a88d706565-RoleDAO").(p24287f458.RoleDAO)
}


func (inst*p982dc5d72e_testcom_TestRbacDgCrudUnit) getSessionDao(ie application.InjectionExt)p24287f458.SessionDAO{
    return ie.GetComponent("#alias-24287f4589fe5add27fb48a88d706565-SessionDAO").(p24287f458.SessionDAO)
}


func (inst*p982dc5d72e_testcom_TestRbacDgCrudUnit) getUserDao(ie application.InjectionExt)p24287f458.UserDAO{
    return ie.GetComponent("#alias-24287f4589fe5add27fb48a88d706565-UserDAO").(p24287f458.UserDAO)
}


