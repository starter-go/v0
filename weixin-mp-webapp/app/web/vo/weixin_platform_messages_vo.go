package vo

import "github.com/starter-go/v0/weixin-mp-webapp/app/web/dto"

type WeixinPlatformMessages struct {
	Base

	Query dto.WeixinPlatformQuery

	Items []*dto.WeixinPlatformMessageHolder `json:"messages"`
}
