package dto

import "github.com/starter-go/v0/rbac-web-app/app/data/dxo"

type User struct {

	// id
	ID dxo.UserID `json:"id"`

	Base

	Name dxo.UserName `json:"username"`

	Email dxo.EmailAddress `json:"email"`

	Mobile dxo.PhoneNumber `json:"mobile"`

	DisplayName string `json:"nickname"`

	Avatar dxo.URL `json:"avatar"`

	Roles dxo.RoleNameSlice `json:"roles"`
}
