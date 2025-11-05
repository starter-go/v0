package vo

import "github.com/starter-go/v0/rbac-web-app/app/web/dto"

type Users struct {
	Base

	Items []*dto.User `json:"users"`
}
