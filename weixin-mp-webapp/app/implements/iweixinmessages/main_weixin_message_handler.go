package iweixinmessages

import (
	"context"

	"github.com/starter-go/v0/weixin-mp-webapp/app/classes/weixinmessages"
	"github.com/starter-go/v0/weixin-mp-webapp/app/data/dxo"
)

type MainWeixinMessageHandler struct {

	//starter:component

	_as func(weixinmessages.Handler) //starter:as("#")
}

// Handle implements weixinmessages.Handler.
func (inst *MainWeixinMessageHandler) Handle(cc context.Context, id dxo.ExampleID) error {
	panic("unimplemented")
}

func (inst *MainWeixinMessageHandler) _impl() weixinmessages.Handler {
	return inst
}
