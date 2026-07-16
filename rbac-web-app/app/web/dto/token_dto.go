package dto

import (
	"github.com/starter-go/rbac"
)

// type Token struct {

// 	// id
// 	ID dxo.TokenID `json:"id"`

// 	Base

// 	SessionID   dxo.SessionID `json:"ses_id"`
// 	SessionUUID lang.UUID     `json:"ses_uuid"`

// 	AliveFrom lang.Time `json:"not_before"`
// 	AliveTo   lang.Time `json:"not_after"`
// }

type Token = rbac.TokenDTO
