package dxo

////////////////////////////////////////////////////////////////////////////////

// UserAuthName 表示用户登录时，使用的 ID 名称，它可以是 [user_name|user_email|user_phone|...]
type UserAuthName string

////////////////////////////////////////////////////////////////////////////////

// 授权步奏

type AuthStep string

const (
	AuthStepOne   AuthStep = "one-step"
	AuthStep2FA1  AuthStep = "2fa-1"
	AuthStep2FA2  AuthStep = "2fa-2"
	AuthStepOther AuthStep = "other"
)

////////////////////////////////////////////////////////////////////////////////

// 授权动作

type AuthAction string

const (
	AuthActionLogin          AuthAction = "login"
	AuthActionSignUp         AuthAction = "sign-up"
	AuthActionChangePassword AuthAction = "change-password"
)

////////////////////////////////////////////////////////////////////////////////

// 认证机制

type AuthMechanism string

const (
	AuthMechanismPassword           AuthMechanism = "password"
	AuthMechanismEmail              AuthMechanism = "email"
	AuthMechanismPhone              AuthMechanism = "phone"
	AuthMechanismPublicKeySignature AuthMechanism = "public-key-signature"
)

////////////////////////////////////////////////////////////////////////////////
// EOF
