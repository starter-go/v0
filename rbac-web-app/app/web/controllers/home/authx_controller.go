package home

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/libgin"
	"github.com/starter-go/rbac"

	"github.com/starter-go/v0/rbac-web-app/app/classes/authx"
	"github.com/starter-go/v0/rbac-web-app/app/data/dxo"
	"github.com/starter-go/v0/rbac-web-app/app/web/dto"
	"github.com/starter-go/v0/rbac-web-app/app/web/vo"
)

type AuthxController struct {

	//starter:component

	_as func(libgin.Controller) //starter:as(".")

	Responder libgin.Responder //starter:inject("#")

	Service authx.Service //starter:inject("#")

}

func (inst *AuthxController) _impl() libgin.Controller {
	return inst
}

func (inst *AuthxController) Registration() *libgin.ControllerRegistration {
	r1 := new(libgin.ControllerRegistration)
	r1.Route = inst.route
	return r1
}

func (inst *AuthxController) route(rp libgin.RouterProxy) error {

	rp = rp.For("auth")

	rp.GET("", inst.handleGetMock)
	rp.POST("", inst.handlePostLogin)

	return nil
}

func (inst *AuthxController) handleGetMock(c *gin.Context) {

	task := new(innerAuthxTask)
	task.context = c
	task.controller = inst

	task.execute(task.doGetMock)
}

func (inst *AuthxController) handlePostLogin(c *gin.Context) {

	task := new(innerAuthxTask)
	task.context = c
	task.controller = inst
	task.wantRequestAuthInHeader = true
	task.wantRequestBody = true

	task.execute(task.doLogin)
}

////////////////////////////////////////////////////////////////////////////////

type innerAuthxTask struct {
	context    *gin.Context
	controller *AuthxController

	wantRequestID           bool
	wantRequestBody         bool
	wantRequestAuthInHeader bool

	id    dxo.ExampleID
	body1 vo.Authx
	body2 vo.Authx
}

func (inst *innerAuthxTask) open(c *gin.Context) error {

	if inst.wantRequestID {
		idstr := c.Param("id")
		num, err := strconv.ParseInt(idstr, 0, 0)
		if err != nil {
			return err
		}
		inst.id = dxo.ExampleID(num)
	}

	if inst.wantRequestBody {
		body := &inst.body1
		err := c.BindJSON(body)
		if err != nil {
			return err
		}
	}

	if inst.wantRequestAuthInHeader {
		info := c.Request.Header.Get("Authorization")
		err := inst.parseHttpAuthorizationHeader(info)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *innerAuthxTask) parseHttpAuthorizationHeader(hStr string) error {

	body := &inst.body1
	list := body.Authentications
	item := new(dto.Authentication)

	hStr = strings.TrimSpace(hStr)
	idx := strings.IndexRune(hStr, ' ')

	if idx > 0 {

		s1 := hStr[0:idx]
		s2 := hStr[idx+1:]

		s1 = strings.TrimSpace(s1)
		s2 = strings.TrimSpace(s2)

		item.Mechanism = dxo.AuthMechanism(s1)
		item.Credentials = lang.Base64(s2)

		// 以下两个字段由具体的认证机制组件解析
		item.KeyID = ""
		item.KeySecret = ""

	} else {
		return fmt.Errorf("bad HTTP Authorization info")
	}

	list = append(list, item)
	body.Authentications = list
	return nil
}

func (inst *innerAuthxTask) send(err error) {

	ctx := inst.context
	body := &inst.body2
	code := body.Status
	resp := new(libgin.Response)
	sender := inst.controller.Responder

	resp.Error = err
	resp.Context = ctx
	resp.Status = code
	resp.Data = body

	sender.Send(resp)
}

func (inst *innerAuthxTask) execute(fn func() error) {
	ctx := inst.context
	err := inst.open(ctx)
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *innerAuthxTask) doGetMock() error {

	// ctx := inst.context
	// ser := inst.controller.Service
	// id := inst.id

	item1 := new(dto.Authentication)
	item2 := new(dto.Authorization)
	page := new(rbac.Pagination)

	inst.body2.Pagination = page
	inst.body2.Authentications = []*dto.Authentication{item1}
	inst.body2.Authorizations = []*dto.Authorization{item2}

	return nil
}

func (inst *innerAuthxTask) doLogin() error {

	ctx := inst.context
	ser := inst.controller.Service
	v1 := &inst.body1

	v2, err := ser.Auth(ctx, v1)
	if err == nil {
		return err
	}

	inst.body2 = *v2
	return nil
}

func (inst *innerAuthxTask) doSignUp() error {

	return nil
}

////////////////////////////////////////////////////////////////////////////////
// EOF
