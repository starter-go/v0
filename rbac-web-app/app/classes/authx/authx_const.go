package authx

import "github.com/starter-go/v0/rbac-web-app/app/data/dxo"

const (
	AuthMechanismPassword           = dxo.AuthMechanismPassword
	AuthMechanismEmail              = dxo.AuthMechanismEmail
	AuthMechanismPhone              = dxo.AuthMechanismPhone
	AuthMechanismPublicKeySignature = dxo.AuthMechanismPublicKeySignature

	AuthMechanismBasicHTTP = dxo.AuthMechanismBasicHTTP
)
