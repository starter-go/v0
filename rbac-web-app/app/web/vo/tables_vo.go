package vo

import "github.com/starter-go/v0/rbac-web-app/app/web/dto"

type Tables struct {
	Base

	Items []*dto.Table `json:"tables"`
}
