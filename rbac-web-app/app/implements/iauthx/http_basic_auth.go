package iauthx

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/v0/rbac-web-app/app/classes/authx"
	"github.com/starter-go/v0/rbac-web-app/app/classes/users"
	"github.com/starter-go/v0/rbac-web-app/app/data/dxo"
	"github.com/starter-go/v0/rbac-web-app/app/data/entity"
	"github.com/starter-go/v0/rbac-web-app/app/web/dto"
	"github.com/starter-go/vlog"
)

// 参考:
// title:  rfc7617  (The 'Basic' HTTP Authentication Scheme)
// url:    https://datatracker.ietf.org/doc/html/rfc7617

type HTTPBasicAuthenticator struct {

	//starter:component

	_as func(authx.Authenticator) //starter:as(".")

	UserAuthNameService authx.UserAuthNameService //starter:inject("#")

	DebugEnabled bool //starter:inject("${debug.enabled}")

}

// Accept implements authx.Authenticator.
func (inst *HTTPBasicAuthenticator) Accept(ctx *authx.AuthenticationContext) bool {
	const want = authx.AuthMechanismBasicHTTP
	have := ctx.Auth.Mechanism
	return have == want
}

// Authenticate implements authx.Authenticator.
func (inst *HTTPBasicAuthenticator) Authenticate(ctx *authx.AuthenticationContext) error {
	err := inst.innerAuthenticate(ctx)
	if err != nil {
		time.Sleep(time.Second)
		return err
	}
	return nil
}

func (inst *HTTPBasicAuthenticator) innerAuthenticate(ctx *authx.AuthenticationContext) error {

	cred := ctx.Auth.Credentials.Bytes()
	task := new(innerHTTPBasicAuthenticatorTask)

	task.context = ctx.Context
	task.parent = inst

	err := task.initWithCredentials(cred)
	if err != nil {
		return err
	}

	err = task.loadUserInfo()
	if err != nil {
		return err
	}

	err = task.verify()
	if err != nil {
		return err
	}

	ctx.User = task.userDto
	return nil
}

// GetRegistration implements authx.Authenticator.
func (inst *HTTPBasicAuthenticator) GetRegistration() *authx.AuthenticatorRegistration {

	const mech = authx.AuthMechanismBasicHTTP

	r1 := new(authx.AuthenticatorRegistration)
	r1.Mechanism = mech
	r1.Authenticator = inst
	r1.Enabled = true
	r1.Priority = 0
	r1.Name = "HTTPBasicAuthenticator"

	return r1
}

func (inst *HTTPBasicAuthenticator) _impl() authx.Authenticator {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

type innerHTTPBasicAuthenticatorTask struct {
	context context.Context
	parent  *HTTPBasicAuthenticator

	username dxo.UserAuthName
	password string

	userDto *dto.User
	userEnt *entity.User
}

func (inst *innerHTTPBasicAuthenticatorTask) initWithCredentials(cred []byte) error {

	const sep = ":"
	str := string(cred)
	strlist := strings.SplitN(str, sep, 2)

	if len(strlist) != 2 {
		return fmt.Errorf("bad auth")
	}

	username := strlist[0]
	inst.username = dxo.UserAuthName(username)
	inst.password = strlist[1]

	return nil
}

func (inst *innerHTTPBasicAuthenticatorTask) loadUserInfo() error {

	cc := inst.context
	ser := inst.parent.UserAuthNameService
	name := inst.username

	userEnt, err := ser.FindUserEntity(cc, name)
	if err != nil {
		return err
	}

	userDto := new(dto.User)
	err = users.ConvertE2D(userEnt, userDto)
	if err != nil {
		return err
	}

	inst.userDto = userDto
	inst.userEnt = userEnt
	return nil
}

func (inst *innerHTTPBasicAuthenticatorTask) verify() error {

	upsc := new(UserPasswordSumComputer)
	user := inst.userEnt

	// 注意: 这里使用 user.uuid 作为 "用户名", 以保证密码的 sum 不会因为用户改名而失效
	name := user.UUID.String()

	upsc.Username = name
	upsc.Password = inst.password
	upsc.Salt = user.Salt.Bytes()

	sumWeb := upsc.Sum()
	sumDB := user.Password.Bytes()

	if !bytes.Equal(sumWeb, sumDB) {
		inst.tryLogBadSum(sumWeb, sumDB)
		return fmt.Errorf("bad auth")
	}

	return nil
}

func (inst *innerHTTPBasicAuthenticatorTask) tryLogBadSum(sumWeb, sumDB []byte) {

	debug := inst.parent.DebugEnabled
	if !debug {
		return
	}

	hexW := lang.HexFromBytes(sumWeb)
	hexD := lang.HexFromBytes(sumDB)

	vlog.Debug("bad password, sum(web)=[%s], sum(db)=[%s]", hexW.String(), hexD.String())

}
