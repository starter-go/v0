package dto

import "github.com/starter-go/v0/rbac-web-app/app/data/dxo"

type Example struct {

	// id
	ID dxo.UserID `json:"id"`

	Base

	Foo int    `json:"foo"`
	Bar string `json:"bar"`
}
