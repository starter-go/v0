package dto

import "github.com/starter-go/v0/web-app-template/app/data/dxo"

type Example struct {
	ID dxo.ExampleID `json:"id"`

	Base

	Foo string `json:"foo"`
	Bar int    `json:"bar"`
}
