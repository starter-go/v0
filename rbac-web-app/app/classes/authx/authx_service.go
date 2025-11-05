package authx

import (
	"context"

	"github.com/starter-go/v0/rbac-web-app/app/web/vo"
)

// 综合服务
type Service interface {

	// 执行（认证+授权）的完整过程
	Auth(ctx context.Context, view *vo.Authx) (*vo.Authx, error)
}
