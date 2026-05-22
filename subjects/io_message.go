package subjects

import (
	"github.com/starter-go/application/properties"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac"
)

type Message struct {
	Properties properties.Table

	// Session     *rbac.SessionDTO

	SessionID   rbac.SessionID
	SessionUUID lang.UUID
}

type Want struct {

	// base of Want
	Message

	Method Method
}

type Have struct {

	// base of Have
	Message

	Status int
}
