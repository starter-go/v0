package vo

import "github.com/starter-go/v0/web-app-template/app/web/dto"

type Examples struct {
	Base

	Items []*dto.Example `json:"examples"`
}
