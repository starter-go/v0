package tokens

import (
	"github.com/starter-go/rbac"
	"github.com/starter-go/v0/rbac-web-app/app/data/entity"
)

type Query struct {
	All bool

	Pagination rbac.Pagination

	Want *entity.Example
}
