package entity

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/v0/rbac-web-app/app/data/dxo"
)

type User struct {

	// id
	ID dxo.UserID

	Base

	// 用户名 （系统唯一）
	Name dxo.UserName `gorm:"unique"`

	// 用户的昵称 （显示名称 ）
	DisplayName string

	//  用户头像的图片地址
	Avatar dxo.URL

	// 用户的邮箱地址
	Email dxo.EmailAddress `gorm:"unique"`

	// 用户的手机号码
	Mobile dxo.PhoneNumber `gorm:"unique"`

	Roles dxo.RoleNameList

	Password lang.Hex // sum of password(with_salt)

	Salt lang.Base64 //  a random number

	Use2FA bool // 是否启用 2FA 机制

	Admin bool // 指出该用户是否为管理员 (admin|root)

	Enabled bool // 表示该账号是否已经激活可用
}
