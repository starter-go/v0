package weixinmessages

import (
	"context"

	"github.com/starter-go/v0/weixin-mp-webapp/app/data/dxo"
)

type Handler interface {
	Handle(cc context.Context, id dxo.ExampleID) error
}
