package dto

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/v0/rbac-web-app/app/data/dxo"
)

// 认证信息
type Authentication struct {

	// 机制 (aka. www-Authorization-type)
	Mechanism dxo.AuthMechanism `json:"mechanism"`

	// 凭据  (aka. www-Authorization-credentials)
	Credentials lang.Base64 `json:"credentials"`

	KeyID     dxo.UserAuthName `json:"auth_key_id"`
	KeySecret lang.Base64      `json:"auth_key_secret"`

	// 附加的参数
	Params map[string]string `json:"params"`
}

// 授权信息
type Authorization struct {
	Action dxo.AuthAction `json:"action"`
	Step   dxo.AuthStep   `json:"step"`

	// 附加的参数
	Params map[string]string `json:"params"`
}
