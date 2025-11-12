package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
	"github.com/starter-go/v0/weixin-mp-webapp/app/classes/weixinmessages"
	"github.com/starter-go/v0/weixin-mp-webapp/app/data/dxo"
	"github.com/starter-go/v0/weixin-mp-webapp/app/web/dto"
	"github.com/starter-go/v0/weixin-mp-webapp/app/web/vo"
	"github.com/starter-go/v0/weixin-mp-webapp/app/web/webutils"
	"github.com/starter-go/vlog"
)

// WeixinPlatformMessageController ...
type WeixinPlatformMessageController struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Sender  libgin.Responder       //starter:inject("#")
	Handler weixinmessages.Handler //starter:inject("#")

	PlatformMessageToken string //starter:inject("${weixin.mp-gateway.token}")
	GatewayPath          string //starter:inject("${weixin.mp-gateway.path}")

}

func (inst *WeixinPlatformMessageController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *WeixinPlatformMessageController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{Route: inst.route}
}

func (inst *WeixinPlatformMessageController) route(rp libgin.RouterProxy) error {

	path2gw := inst.GatewayPath
	rp = rp.For("weixin-mp")

	rp.GET(path2gw, inst.handleGetGateway)
	rp.POST(path2gw, inst.handlePostGateway)

	return nil
}

func (inst *WeixinPlatformMessageController) handle(c *gin.Context) {
	req := &myWeixinPlatformMessageRequest{
		context:       c,
		controller:    inst,
		wantRequestID: false,
	}
	req.execute(req.doNOP)
}

func (inst *WeixinPlatformMessageController) handleGetGateway(c *gin.Context) {
	req := &myWeixinPlatformMessageRequest{
		context:          c,
		controller:       inst,
		wantRequestQuery: true,
	}
	req.execute(req.doGetGateway)
}

func (inst *WeixinPlatformMessageController) handlePostGateway(c *gin.Context) {
	req := &myWeixinPlatformMessageRequest{
		context:            c,
		controller:         inst,
		wantRequestQuery:   true,
		wantRequestBody:    true,
		wantRequestBodyXML: true,
	}
	req.execute(req.doPostGateway)
}

////////////////////////////////////////////////////////////////////////////////

type myWeixinPlatformMessageRequest struct {
	context    *gin.Context
	controller *WeixinPlatformMessageController

	wantRequestID      bool
	wantRequestBody    bool
	wantRequestBodyXML bool
	wantRequestQuery   bool

	id    dxo.ExampleID
	body1 vo.WeixinPlatformMessages
	body2 vo.WeixinPlatformMessages
}

func (inst *myWeixinPlatformMessageRequest) open() error {

	c := inst.context

	if inst.wantRequestID {
		idstr := c.Param("id")
		idnum, err := strconv.ParseInt(idstr, 10, 64)
		if err != nil {
			return err
		}
		inst.id = dxo.ExampleID(idnum)
	}

	if inst.wantRequestBody {
		if inst.wantRequestBodyXML {
			err := inst.parseRequestBodyXML()
			if err != nil {
				return err
			}
		}
	}

	if inst.wantRequestQuery {
		q := &inst.body1.Query
		q.Signature = c.Query("signature")
		q.TimeStamp = c.Query("timestamp")
		q.Nonce = c.Query("nonce")
		q.EchoStr = c.Query("echostr")
	}

	return nil
}

func (inst *myWeixinPlatformMessageRequest) parseRequestBodyXML() error {

	ctx := inst.context
	obj := new(dto.WeixinPlatformMessageBase)
	holder := new(dto.WeixinPlatformMessageHolder)
	binder := new(webutils.RequestBodyBinder)

	err := binder.Init(ctx)
	if err != nil {
		return err
	}

	err = binder.BindXML(obj)
	if err != nil {
		return err
	}

	mtype := obj.MsgType.Normalize()
	switch mtype {

	case dto.WxMsgTypeText:
		sub := new(dto.WeixinPlatformMessageText)
		err = binder.BindXML(sub)
		holder.Text = sub.Normalize()

	case dto.WxMsgTypeImage:
		sub := new(dto.WeixinPlatformMessageImage)
		err = binder.BindXML(sub)
		holder.Image = sub.Normalize()

	case dto.WxMsgTypeVoice:
		sub := new(dto.WeixinPlatformMessageVoice)
		err = binder.BindXML(sub)
		holder.Voice = sub.Normalize()

	case dto.WxMsgTypeVideo:
		sub := new(dto.WeixinPlatformMessageVideo)
		err = binder.BindXML(sub)
		holder.Video = sub.Normalize()

	case dto.WxMsgTypeShortVideo:
		sub := new(dto.WeixinPlatformMessageShortVideo)
		err = binder.BindXML(sub)
		holder.ShortVideo = sub.Normalize()

	case dto.WxMsgTypeLocation:
		sub := new(dto.WeixinPlatformMessageLocation)
		err = binder.BindXML(sub)
		holder.Location = sub.Normalize()

	case dto.WxMsgTypeLink:
		sub := new(dto.WeixinPlatformMessageLink)
		err = binder.BindXML(sub)
		holder.Link = sub.Normalize()

	default:
	}
	if err != nil {
		return err
	}

	jstr, err := json.MarshalIndent(holder, "\t", "\t")
	if err == nil {
		vlog.Info("req.BindXML = %v", string(jstr))
	}

	inst.body1.Items = []*dto.WeixinPlatformMessageHolder{holder}
	return nil
}

func (inst *myWeixinPlatformMessageRequest) send(err error) {
	ctx := inst.context
	method := ctx.Request.Method
	switch method {
	case http.MethodPost:
		inst.sendForPost(err)
	case http.MethodGet:
		inst.sendForGet(err)
	}
}

func (inst *myWeixinPlatformMessageRequest) sendForGet(err error) {
	ctx := inst.context
	code := http.StatusOK
	ctype := "text/plain"
	data := inst.body1.Query.EchoStr
	if err != nil {
		code = http.StatusInternalServerError
		data = "error: " + err.Error()
	}
	ctx.Data(code, ctype, []byte(data))
}

func (inst *myWeixinPlatformMessageRequest) sendForPost(err error) {
	data := &inst.body2
	code := inst.body2.Status
	resp := new(libgin.Response)
	resp.Context = inst.context
	resp.Error = err
	resp.Data = data
	resp.Status = code
	inst.controller.Sender.Send(resp)
}

func (inst *myWeixinPlatformMessageRequest) execute(fn func() error) {
	err := inst.open()
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *myWeixinPlatformMessageRequest) doNOP() error {
	return nil
}

func (inst *myWeixinPlatformMessageRequest) checkQueryInfo() error {

	q := &inst.body1.Query

	signature := q.Signature
	timestamp := q.TimeStamp
	nonce := q.Nonce
	echostr := q.EchoStr
	token := inst.controller.PlatformMessageToken

	vlog.Info("myWeixinPlatformMessageRequest.doGetGateway()")
	vlog.Info("   q.signature = %v", signature)
	vlog.Info("   q.timestamp = %v", timestamp)
	vlog.Info("   q.nonce     = %v", nonce)
	vlog.Info("   q.echostr   = %v", echostr)

	return q.VerifyWithToken(token)
}

func (inst *myWeixinPlatformMessageRequest) doPostGateway() error {

	err := inst.checkQueryInfo()
	if err != nil {
		return err
	}

	return nil
}

func (inst *myWeixinPlatformMessageRequest) doGetGateway() error {

	err := inst.checkQueryInfo()
	if err != nil {
		return err
	}

	return nil
}
