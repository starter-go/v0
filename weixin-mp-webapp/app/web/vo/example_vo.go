package vo

import "github.com/starter-go/v0/weixin-mp-webapp/app/web/dto"

type Examples struct {
	Base

	Items []*dto.Example `json:"examples"`
}
