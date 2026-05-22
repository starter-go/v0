package iadapter

import (
	"github.com/starter-go/v0/libjwt/api/jwt"
)

////////////////////////////////////////////////////////////////////////////////

type MainAdapter struct {

	//starter:component

	_as func(jwt.Adapter) //starter:as("#")

}

// GetToken implements jwt.Adapter.
func (inst *MainAdapter) GetToken(acc *jwt.Access) error {
	panic("unimplemented : GetToken() ")
}

// SetToken implements jwt.Adapter.
func (inst *MainAdapter) SetToken(acc *jwt.Access) error {
	panic("unimplemented : SetToken() ")
}

func (inst *MainAdapter) _impl() jwt.Adapter {
	return inst
}

////////////////////////////////////////////////////////////////////////////////
