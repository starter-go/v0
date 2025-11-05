package admin

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
	"github.com/starter-go/rbac"
	"github.com/starter-go/v0/rbac-web-app/app/classes/users"
	"github.com/starter-go/v0/rbac-web-app/app/data/dxo"
	"github.com/starter-go/v0/rbac-web-app/app/web/dto"
	"github.com/starter-go/v0/rbac-web-app/app/web/vo"
)

type UsersController struct {

	//starter:component

	_as func(libgin.Controller) //starter:as(".")

	Responder libgin.Responder //starter:inject("#")

	Service users.Service //starter:inject("#")

}

func (inst *UsersController) _impl() libgin.Controller {
	return inst
}

func (inst *UsersController) Registration() *libgin.ControllerRegistration {
	r1 := new(libgin.ControllerRegistration)
	r1.Route = inst.route
	return r1
}

func (inst *UsersController) route(rp libgin.RouterProxy) error {

	rp = rp.For("admin/users")

	rp.GET("", inst.handleGetMock)

	return nil
}

func (inst *UsersController) handleGetMock(c *gin.Context) {

	task := new(innerUsersTask)
	task.context = c
	task.controller = inst

	task.execute(task.doGetMock)
}

////////////////////////////////////////////////////////////////////////////////

type innerUsersTask struct {
	context    *gin.Context
	controller *UsersController

	wantRequestID   bool
	wantRequestBody bool

	id    dxo.UserID
	body1 vo.Users
	body2 vo.Users
}

func (inst *innerUsersTask) open(c *gin.Context) error {

	if inst.wantRequestID {
		idstr := c.Param("id")
		num, err := strconv.ParseInt(idstr, 0, 0)
		if err != nil {
			return err
		}
		inst.id = dxo.UserID(num)
	}

	if inst.wantRequestBody {
		body := &inst.body1
		err := c.BindJSON(body)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *innerUsersTask) send(err error) {

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

func (inst *innerUsersTask) execute(fn func() error) {
	ctx := inst.context
	err := inst.open(ctx)
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *innerUsersTask) doGetMock() error {

	ctx := inst.context
	ser := inst.controller.Service
	id := inst.id

	item, err := ser.Find(ctx, id)
	if err != nil {
		return err
	}

	page := new(rbac.Pagination)

	inst.body2.Pagination = page
	inst.body2.Items = []*dto.User{item}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// EOF
