package subjects

import (
	"github.com/starter-go/application/properties"
	"github.com/starter-go/rbac"
	"github.com/starter-go/v0/subjects/core/data/dxo"
)

type Cache struct {
	Properties properties.Table

	// Avatar string // a http URL to avatar
	// UserID rbac.UserID
	// UserName string
	// NickName string
	// Roles rbac.RoleNameList
	// SessionDTO *rbac.SessionDTO

	SessionID rbac.SessionID

	SessionUUID dxo.SessionUUID

	Loaded bool
}

func (inst *Cache) Ready() bool {
	loaded := inst.Loaded
	pro := inst.Properties
	return ((pro != nil) && loaded)
}

////////////////////////////////////////////////////////////////////////////////

type Buffer struct {
	Properties properties.Table

	SessionID rbac.SessionID

	SessionUUID dxo.SessionUUID
}

func (inst *Buffer) Init() *Buffer {
	if inst == nil {
		inst = new(Buffer)
	}
	if inst.Properties == nil {
		inst.Properties = properties.NewTable(nil)
	}
	return inst
}

func (inst *Buffer) HasData() bool {

	if inst == nil {
		return false
	}

	pt := inst.Properties
	if pt == nil {
		return false
	}

	keys := pt.Names()
	if keys == nil {
		return false
	}

	if len(keys) == 0 {
		return false
	}

	return true
}

func (inst *Buffer) IsEmpty() bool {
	return !inst.HasData()
}

////////////////////////////////////////////////////////////////////////////////
