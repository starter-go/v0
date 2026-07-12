package entity

import (
	"github.com/starter-go/security-gorm/rbacdb"
)

type Permission struct {
	rbacdb.PermissionEntity
}
