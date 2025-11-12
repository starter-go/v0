package webutils

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/v0/weixin-mp-webapp/app/data/dxo"
	"github.com/starter-go/v0/weixin-mp-webapp/app/web/dto"
	"github.com/starter-go/v0/weixin-mp-webapp/app/web/vo"
	"github.com/starter-go/vlog"
)

//  参考:
// <能力接入 /基础消息与订阅通知 /被动回复用户消息>
// https://developers.weixin.qq.com/doc/service/guide/product/message/Passive_user_reply_message.html

////////////////////////////////////////////////////////////////////////////////

type ResponseBodyBuilder struct{}

////////////////////////////////////////////////////////////////////////////////

type ResponseBodySender struct {
	Context *gin.Context

	ClientSideUserName dxo.WxMsgUserName
	ServerSideUserName dxo.WxMsgUserName
	TimeStamp          dxo.WxMsgTime
}

func (inst *ResponseBodySender) Send(view *vo.WeixinPlatformMessages) error {

	item, err := inst.innerGetFirstItem(view)
	if err != nil {
		return err
	}

	ok := inst.trySendText(item)
	if ok {
		return nil
	}

	ok = inst.trySendImage(item)
	if ok {
		return nil
	}

	ok = inst.trySendVoice(item)
	if ok {
		return nil
	}

	ok = inst.trySendVideo(item)
	if ok {
		return nil
	}

	return nil
}

func (inst *ResponseBodySender) innerGetFirstItem(view *vo.WeixinPlatformMessages) (*dto.WeixinPlatformMsgResponseHolder, error) {
	all := view.Items2
	for _, it := range all {
		if it != nil {
			return it, nil
		}
	}
	return nil, fmt.Errorf("no response message")
}

func (inst *ResponseBodySender) trySendText(h *dto.WeixinPlatformMsgResponseHolder) bool {
	obj := h.Text
	if obj != nil {
		inst.innerPrepareBodyBase(&obj.WeixinPlatformMsgResponseBase, dxo.WxMsgTypeText)
		inst.innerSendXML(obj)
		return true
	}
	return false
}
func (inst *ResponseBodySender) trySendImage(h *dto.WeixinPlatformMsgResponseHolder) bool {
	obj := h.Image
	if obj != nil {
		inst.innerPrepareBodyBase(&obj.WeixinPlatformMsgResponseBase, dxo.WxMsgTypeImage)
		inst.innerSendXML(obj)
		return true
	}
	return false
}
func (inst *ResponseBodySender) trySendVoice(h *dto.WeixinPlatformMsgResponseHolder) bool {
	obj := h.Voice
	if obj != nil {
		inst.innerPrepareBodyBase(&obj.WeixinPlatformMsgResponseBase, dxo.WxMsgTypeVoice)
		inst.innerSendXML(obj)
		return true
	}
	return false
}

func (inst *ResponseBodySender) trySendVideo(h *dto.WeixinPlatformMsgResponseHolder) bool {
	obj := h.Video
	if obj != nil {
		inst.innerPrepareBodyBase(&obj.WeixinPlatformMsgResponseBase, dxo.WxMsgTypeVideo)
		inst.innerSendXML(obj)
		return true
	}
	return false
}

func (inst *ResponseBodySender) innerPrepareBodyBase(bb *dto.WeixinPlatformMsgResponseBase, mt dxo.WxMsgType) {
	bb.CreateTime = inst.TimeStamp
	bb.FromUserName = inst.ServerSideUserName
	bb.ToUserName = inst.ClientSideUserName
	bb.MsgType = mt
}

func (inst *ResponseBodySender) innerSendXML(o any) error {

	data, err := xml.Marshal(o)
	if err != nil {
		return err
	}

	if vlog.IsDebugEnabled() {
		vlog.Debug("xml.content = [%s]", string(data))
	}

	code := http.StatusOK
	ctype := "application/xml"
	ctx := inst.Context

	ctx.Data(code, ctype, data)
	return nil
}

////////////////////////////////////////////////////////////////////////////////
