package vo

import "github.com/starter-go/v0/rbac-web-app/app/web/dto"

type Authx struct {
	Base

	// 认证信息列表
	Authentications []*dto.Authentication `json:"authentications"`

	// 授权信息列表
	Authorizations []*dto.Authorization `json:"authorizations"`
}
