package controllers

import (
	"net/http"

	"github.com/starter-go/rbac"
	"github.com/starter-go/security/permissions"
)

type ExamplePermissionsProvider struct {

	//starter:component

	_as func(permissions.Registry) //starter:as(".")

}

// ListRegistrations implements permissions.Registry.
func (inst *ExamplePermissionsProvider) ListRegistrations() []*permissions.Registration {

	list := make([]*permissions.Registration, 0)

	list = inst.accept(list, http.MethodGet, "/api/v1/admin/users/:id", rbac.RoleAdmin)
	list = inst.accept(list, http.MethodGet, "/api/v1/auth", rbac.RoleAnonym)

	return list
}

func (inst *ExamplePermissionsProvider) accept(list []*permissions.Registration, method string, path string, roles ...rbac.RoleName) []*permissions.Registration {

	reg := new(permissions.Registration)

	reg.Method = method
	reg.Path = path
	reg.Roles = rbac.NewRoleNameList(roles...)
	reg.Enabled = true

	list = append(list, reg)
	return list
}

func (inst *ExamplePermissionsProvider) _impl() permissions.Registry {
	return inst
}
