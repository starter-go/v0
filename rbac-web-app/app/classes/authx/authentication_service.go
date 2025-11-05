package authx

import (
	"context"

	"github.com/starter-go/v0/rbac-web-app/app/data/dxo"
	"github.com/starter-go/v0/rbac-web-app/app/web/dto"
	"github.com/starter-go/v0/rbac-web-app/app/web/vo"
)

// 认证上下文
type AuthenticationContext struct {

	// input

	Context context.Context

	Auth *dto.Authentication

	View *vo.Authx

	// output

	User *dto.User
}

// 认证服务
type AuthenticationService interface {

	// 执行认证操作
	Authenticate(ctx *AuthenticationContext) error
}

// 认证服务-提供者
type Authenticator interface {

	// 取 Authenticator 的注册信息
	GetRegistration() *AuthenticatorRegistration

	Accept(ctx *AuthenticationContext) bool

	// 执行认证操作
	Authenticate(ctx *AuthenticationContext) error
}

type AuthenticatorRegistration struct {
	Name string

	Enabled bool

	Priority int

	Mechanism dxo.AuthMechanism

	Authenticator Authenticator
}
