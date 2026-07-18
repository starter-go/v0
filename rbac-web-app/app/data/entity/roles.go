package entity

import (
	"github.com/starter-go/rbac"
	"github.com/starter-go/security-gorm/rbacdb"
)

type Role struct {
	rbac.RoleEntity
}

type RoleTB struct {
	rbacdb.BaseEntity
}
