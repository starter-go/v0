package iauthx

import (
	"context"
	"fmt"

	"github.com/starter-go/v0/rbac-web-app/app/classes/authx"
	"github.com/starter-go/v0/rbac-web-app/app/classes/users"
	"github.com/starter-go/v0/rbac-web-app/app/data/dxo"
	"github.com/starter-go/v0/rbac-web-app/app/data/entity"
)

type UserAuthNameServiceImpl struct {

	//starter:component

	_as func(authx.UserAuthNameService) //starter:as("#")

	Dao users.DAO //starter:inject("#")

}

// FindUserEntity implements authx.UserAuthNameService.
func (inst *UserAuthNameServiceImpl) FindUserEntity(cc context.Context, name dxo.UserAuthName) (*entity.User, error) {

	if name.IsEmailAddress() {
		return inst.innerFindUserEntityByEmailAddr(cc, name)

	} else if name.IsPhoneNumber() {
		return inst.innerFindUserEntityByPhoneNum(cc, name)

	} else if name.IsUserName() {
		return inst.innerFindUserEntityByUserName(cc, name)
	}

	return nil, fmt.Errorf("no account")
}

func (inst *UserAuthNameServiceImpl) innerFindUserEntityByUserName(cc context.Context, name dxo.UserAuthName) (*entity.User, error) {
	un := name.String()
	return inst.Dao.FindByName(nil, dxo.UserName(un))
}

func (inst *UserAuthNameServiceImpl) innerFindUserEntityByPhoneNum(cc context.Context, name dxo.UserAuthName) (*entity.User, error) {
	num := name.String()
	return inst.Dao.FindByMobile(nil, dxo.PhoneNumber(num))
}

func (inst *UserAuthNameServiceImpl) innerFindUserEntityByEmailAddr(cc context.Context, name dxo.UserAuthName) (*entity.User, error) {
	addr := name.String()
	return inst.Dao.FindByEmail(nil, dxo.EmailAddress(addr))
}

func (inst *UserAuthNameServiceImpl) _impl() authx.UserAuthNameService {
	return inst
}
