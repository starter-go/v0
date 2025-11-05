package entity

import "github.com/starter-go/v0/rbac-web-app/app/data/dxo"

type Example struct {

	// id
	ID dxo.ExampleID

	Base

	Foo int
	Bar string
}
