package authx

import (
	"context"

	"github.com/starter-go/v0/rbac-web-app/app/web/dto"
	"github.com/starter-go/v0/rbac-web-app/app/web/vo"
)

// 授权上下文
type AuthorizationContext struct {
	Context context.Context

	View *vo.Authx

	Auth *dto.Authorization
}

// 授权服务
type AuthorizationService interface {

	// 执行授权操作
	Authorize(ctx *AuthorizationContext) error
}

type Authorizer interface {

	// 取 Authorizer 的注册信息
	GetRegistration() *AuthorizerRegistration

	Accept(ctx *AuthorizationContext) bool

	Authorize(ctx *AuthorizationContext) error
}

type AuthorizerRegistration struct {
	Name string

	Enabled bool

	Priority int

	Authorizer Authorizer
}
