package entities

import (
	"time"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac"
	"github.com/starter-go/v0/subjects/core/data/dxo"
)

type SessionEntity struct {

	// id

	ID dxo.SessionID

	// base

	Base

	// fields

	NotBefore time.Time

	NotAfter time.Time

	Authenticated bool

	Roles rbac.RoleNameList

	// Username rbac.UserName
	// Email rbac.EmailAddress
	// Language localization.Locale
	// Avatar string
	// Nickname string

	MaxTokenAge lang.Milliseconds

	// ext

	Properties dxo.ProText // include: { avatar , nickname, roles ,  ... }
}
