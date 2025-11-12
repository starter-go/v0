package weixinmessages

import (
	"github.com/starter-go/rbac"
	"github.com/starter-go/v0/weixin-mp-webapp/app/data/entity"
)

type Query struct {
	All bool

	Pagination rbac.Pagination

	Want *entity.Example

	QueryString string

	QueryArgs []any
}
