package subjects

import (
	"github.com/starter-go/application/properties"
	"github.com/starter-go/rbac"
	"github.com/starter-go/v0/subjects/core/data/dxo"
)

type Cache struct {
	Properties properties.Table

	Avatar string // a http URL to avatar

	UserID rbac.UserID

	UserName string

	NickName string

	Roles rbac.RoleNameList

	SessionID rbac.SessionID

	SessionUUID dxo.SessionUUID

	SessionDTO *rbac.SessionDTO

	Loaded bool
}

func (inst *Cache) Ready() bool {
	return inst.Loaded
}

////////////////////////////////////////////////////////////////////////////////

type Buffer struct {
	Properties properties.Table

	SessionID rbac.SessionID

	SessionUUID dxo.SessionUUID
}
