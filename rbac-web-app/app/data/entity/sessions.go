package entity

import (
	"time"

	"github.com/starter-go/v0/rbac-web-app/app/data/dxo"
)

type Session struct {

	// id
	ID dxo.SessionID

	Base

	AliveFrom time.Time
	AliveTo   time.Time

	Alive bool
}
