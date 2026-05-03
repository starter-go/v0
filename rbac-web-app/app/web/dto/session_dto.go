package dto

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/v0/rbac-web-app/app/data/dxo"
)

type Session struct {

	// id
	ID dxo.SessionID `json:"id"`

	Base

	DisplayName string            `json:"nickname"`
	Roles       dxo.RoleNameSlice `json:"roles"`
	Avatar      dxo.URL           `json:"avatar"`

	Alive     bool      `json:"alive"`
	AliveFrom lang.Time `json:"not_before"`
	AliveTo   lang.Time `json:"not_after"`

	Properties map[string]string `json:"properties"`
}
