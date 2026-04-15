package home

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
	"github.com/starter-go/rbac"

	"github.com/starter-go/v0/rbac-web-app/app/classes/sessions"
	"github.com/starter-go/v0/rbac-web-app/app/data/dxo"
	"github.com/starter-go/v0/rbac-web-app/app/web/dto"
	"github.com/starter-go/v0/rbac-web-app/app/web/vo"
)

type SessionController struct {

	//starter:component

	_as func(libgin.Controller) //starter:as(".")

	Responder libgin.Responder //starter:inject("#")

	Service sessions.Service //starter:inject("#")

}

func (inst *SessionController) _impl() libgin.Controller {
	return inst
}

func (inst *SessionController) Registration() *libgin.ControllerRegistration {
	r1 := new(libgin.ControllerRegistration)
	r1.Route = inst.route
	return r1
}

func (inst *SessionController) route(rp libgin.RouterProxy) error {

	rp = rp.For("sessions")

	rp.GET("", inst.handleGetCurrentSessionInfo) // alias to 'current'

	rp.GET("example", inst.handleGetExample)
	rp.GET("current", inst.handleGetCurrentSessionInfo)

	return nil
}

func (inst *SessionController) handleGetExample(c *gin.Context) {

	task := new(innerSessionTask)
	task.context = c
	task.controller = inst

	task.execute(task.doGetMock)
}

func (inst *SessionController) handleGetCurrentSessionInfo(c *gin.Context) {

	task := new(innerSessionTask)
	task.context = c
	task.controller = inst

	task.execute(task.doGetCurrentSessionInfo)
}

////////////////////////////////////////////////////////////////////////////////

type innerSessionTask struct {
	context    *gin.Context
	controller *SessionController

	wantRequestID   bool
	wantRequestBody bool

	id    dxo.SessionID
	body1 vo.Sessions
	body2 vo.Sessions
}

func (inst *innerSessionTask) open(c *gin.Context) error {

	if inst.wantRequestID {
		idstr := c.Param("id")
		num, err := strconv.ParseInt(idstr, 0, 0)
		if err != nil {
			return err
		}
		inst.id = dxo.SessionID(num)
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

func (inst *innerSessionTask) send(err error) {

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

func (inst *innerSessionTask) execute(fn func() error) {
	ctx := inst.context
	err := inst.open(ctx)
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *innerSessionTask) doGetMock() error {

	// ctx := inst.context
	// ser := inst.controller.Service
	// id := inst.id

	item := new(dto.Session)
	page := new(rbac.Pagination)

	inst.body2.Pagination = page
	inst.body2.Items = []*dto.Session{item}
	return nil
}

func (inst *innerSessionTask) doGetCurrentSessionInfo() error {

	ctx := inst.context
	ser := inst.controller.Service

	info, err := ser.GetCurrent(ctx)
	if err != nil {
		return err
	}

	inst.body2.Items = []*dto.Session{info}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// EOF
