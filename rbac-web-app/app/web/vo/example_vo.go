package vo

import "github.com/starter-go/v0/rbac-web-app/app/web/dto"

type Example struct {
	Base

	Items []*dto.Example `json:"examples"`
}
