package daos

import "github.com/starter-go/rbac"

type IAuthenticationDao = rbac.AuthenticationDAO

type IPermissionDao = rbac.PermissionDAO

type IRoleDao = rbac.RoleDAO

type ISessionDao = rbac.SessionDAO

type ITableDao = rbac.TableDAO

type IUserDao = rbac.UserDAO
