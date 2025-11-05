package dto

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/v0/rbac-web-app/app/data/dxo"
)

type Session struct {

	// id
	ID dxo.SessionID `json:"id"`

	Base

	Roles []dxo.RoleName `json:"roles"`

	DisplayName string `json:"nickname"`

	Avatar dxo.URL `json:"avatar"`

	AliveFrom lang.Time `json:"not_before"`
	AliveTo   lang.Time `json:"not_after"`
}
