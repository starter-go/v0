package dto

import (
	"github.com/starter-go/rbac"
)

// type Session struct {

// 	// id
// 	ID dxo.SessionID `json:"id"`

// 	Base

// 	DisplayName string            `json:"nickname"`
// 	Roles       dxo.RoleNameSlice `json:"roles"`
// 	Avatar      dxo.URL           `json:"avatar"`

// 	Alive     bool      `json:"alive"`
// 	AliveFrom lang.Time `json:"not_before"`
// 	AliveTo   lang.Time `json:"not_after"`

// 	Properties map[string]string `json:"properties"`
// }

type Session = rbac.SessionDTO
