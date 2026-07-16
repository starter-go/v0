package dxo

import "github.com/starter-go/rbac"

type IntegerID int64

//////////////////////////////////

type ExampleID IntegerID

type UserID = rbac.UserID

type PermissionID = rbac.PermissionID

type RoleID = rbac.RoleID

type SessionID = rbac.SessionIID

type TokenID IntegerID
