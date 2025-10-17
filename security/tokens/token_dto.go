package tokens

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/v0/security/share"
)

type TokenDTO struct {
	SessionID   share.SessionID
	SessionUUID share.SessionUUID
	StartedAt   lang.Time
	StoppedAt   lang.Time
}
