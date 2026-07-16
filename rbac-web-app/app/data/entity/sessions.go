package entity

import "github.com/starter-go/v0/subjects/core/classes/sessions"

// type Session struct {

// 	// id
// 	ID dxo.SessionID

// 	Base

// 	AliveFrom time.Time
// 	AliveTo   time.Time

// 	DisplayName string
// 	Avatar      dxo.URL
// 	Roles       dxo.RoleNameList

// 	Alive bool
// }

type Session = sessions.Entity
