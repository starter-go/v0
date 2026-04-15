package authx

import (
	"context"

	"github.com/starter-go/v0/rbac-web-app/app/data/dxo"
	"github.com/starter-go/v0/rbac-web-app/app/web/dto"
	"github.com/starter-go/v0/rbac-web-app/app/web/vo"
)

// 授权操作上下文
type AuthorizationContext struct {

	// input

	AuthUser *dto.User // 已经认证的用户信息

	Context context.Context

	View *vo.Authx

	Auth *dto.Authorization
}

// 授权服务
type AuthorizationService interface {

	// 执行授权操作
	Authorize(ctx *AuthorizationContext) error
}

// 授权服务提供者
type Authorizer interface {

	// 取 Authorizer 的注册信息
	GetRegistration() *AuthorizerRegistration

	Accept(ctx *AuthorizationContext) bool

	Authorize(ctx *AuthorizationContext) error
}

// 授权服务注册信息
type AuthorizerRegistration struct {
	Name string

	Enabled bool

	Priority int

	Action dxo.AuthAction

	Authorizer Authorizer
}
