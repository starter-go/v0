package idatagroup

import (
	"github.com/starter-go/rbac"
	"github.com/starter-go/v0/rbac-data-group/src/main/golang/api/daos"
)

////////////////////////////////////////////////////////////////////////////////

type StdRbacDaoSet struct {

	//starter:component

	_as func(rbac.DaoSet) //starter:as(".")

	// config :

	ConfigEnabled  bool //starter:inject("${rbac-dao-set.standard.enabled}")
	ConfigPriority int  //starter:inject("${rbac-dao-set.standard.priority}")

	// daos :

	Authents    daos.IAuthenticationDao //starter:inject("#")
	Permissions daos.IPermissionDao     //starter:inject("#")
	Roles       daos.IRoleDao           //starter:inject("#")
	Sessions    daos.ISessionDao        //starter:inject("#")
	Tables      daos.ITableDao          //starter:inject("#")
	Users       daos.IUserDao           //starter:inject("#")
}

// Provide implements [rbac.DaoSetProvider].
func (inst *StdRbacDaoSet) Provide(dst *rbac.DaoSet) *rbac.DaoSet {

	if dst == nil {
		dst = new(rbac.DaoSet)
	}

	dst.Authentications = inst.Authents

	dst.Tables = inst.Tables

	dst.Users = inst.Users

	return dst
}

// Registration implements [rbac.DaoSetRegistry].
func (inst *StdRbacDaoSet) Registration() *rbac.DaoSetRegistration {
	r1 := &rbac.DaoSetRegistration{
		Provider: inst,
		Label:    "Standard_RBAC_DAO_Set",
		Enabled:  inst.ConfigEnabled,
		Priority: inst.ConfigPriority,
	}
	return r1
}

func (inst *StdRbacDaoSet) _impl() (rbac.DaoSetRegistry, rbac.DaoSetProvider) {
	return inst, inst
}

////////////////////////////////////////////////////////////////////////////////
