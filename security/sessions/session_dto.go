package sessions

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac"
	"github.com/starter-go/v0/security/share"
)

type SessionID = share.SessionID

type SessionUUID = share.SessionUUID

type SessionDTO struct {
	ID SessionID

	rbac.BaseDTO

	Phone       string
	Email       string
	DisplayName string // aka. nick-name
	Avatar      string // a URL to avatar image

	Authenticated bool // 表示这个 session 的用户已经通过验证

	StartedAt lang.Time // 这个 session 的开始生效时间
	StoppedAt lang.Time // 这个 session 的失效时间
}
