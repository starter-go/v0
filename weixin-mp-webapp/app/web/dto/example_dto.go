package dto

import "github.com/starter-go/v0/weixin-mp-webapp/app/data/dxo"

type Example struct {

	// id
	ID dxo.ExampleID `json:"id"`

	Base

	Foo string `json:"foo"`
	Bar int    `json:"bar"`
}
