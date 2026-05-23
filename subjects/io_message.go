package subjects

import (
	"github.com/starter-go/application/properties"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac"
)

type Message struct {
	SessionID rbac.SessionID

	SessionUUID lang.UUID

	Properties properties.Table // 引用 buffer | cache 中的属性
}

type Want struct {

	// base of Want
	Message

	Buffer *Buffer

	Method Method
}

type Have struct {

	// base of Have
	Message

	Cache *Cache

	Status int
}
