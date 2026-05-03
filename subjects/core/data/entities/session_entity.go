package entities

import (
	"time"

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

	Properties dxo.ProText // include: { avatar , nickname, roles ,  ... }
}
