package vo

import "github.com/starter-go/v0/weixin-mp-webapp/app/web/dto"

type WeixinPlatformMessages struct {
	Base

	Query dto.WeixinPlatformQuery

	// request
	Items []*dto.WeixinPlatformMessageHolder `json:"messages"`

	// response
	Items2 []*dto.WeixinPlatformMsgResponseHolder `json:"msg2s"`
}
