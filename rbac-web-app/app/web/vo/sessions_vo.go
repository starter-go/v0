package vo

import "github.com/starter-go/v0/rbac-web-app/app/web/dto"

type Sessions struct {
	Base

	Items []*dto.Session `json:"sessions"`
}
