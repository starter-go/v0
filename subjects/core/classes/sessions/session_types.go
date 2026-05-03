package sessions

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac"
	"github.com/starter-go/v0/subjects/core/data/entities"
)

type ID = rbac.SessionID

type UUID = lang.UUID

type DTO = rbac.SessionDTO

type Entity = entities.SessionEntity
