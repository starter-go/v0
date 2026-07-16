package my

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/libgin"
	"github.com/starter-go/rbac"

	"github.com/starter-go/v0/rbac-web-app/app/data/dxo"
	"github.com/starter-go/v0/rbac-web-app/app/web/dto"
	"github.com/starter-go/v0/rbac-web-app/app/web/vo"
	"github.com/starter-go/v0/subjects"
	"github.com/starter-go/v0/subjects/core/classes/sessions"
)

type SessionController struct {

	//starter:component

	_as func(libgin.Controller) //starter:as(".")

	Responder libgin.Responder //starter:inject("#")

	// Service sessions.Service //x-starter:inject("#")

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

	rp = rp.For("/api/v1/sessions")

	rp.GET("", inst.handleGetCurrentSessionInfo) // alias to 'current'
	rp.GET("current", inst.handleGetCurrentSessionInfo)
	rp.GET("example", inst.handleGetExample)

	rp.POST("keep-alive", inst.handlePostKeepAlive)
	rp.POST("exit", inst.handlePostExit)

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

func (inst *SessionController) handlePostKeepAlive(c *gin.Context) {

	task := new(innerSessionTask)
	task.context = c
	task.controller = inst

	task.execute(task.doKeepAlive)
}

func (inst *SessionController) handlePostExit(c *gin.Context) {

	task := new(innerSessionTask)
	task.context = c
	task.controller = inst

	task.execute(task.doExit)
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

	sub, err := subjects.GetCurrent(ctx)
	if err != nil {
		return err
	}

	gett, err := sub.DoGet()
	if err != nil {
		return err
	}

	se := new(sessions.DTO)
	gett.GetSession(se)

	tk := new(rbac.TokenDTO)
	gett.GetToken(tk)

	inst.body2.Items = []*rbac.SessionDTO{se}
	inst.body2.Tokens = []*rbac.TokenDTO{tk}

	return nil
}

func (inst *innerSessionTask) doKeepAlive() error {

	ctx := inst.context
	now := lang.Now()

	sub, err := subjects.GetCurrent(ctx)
	if err != nil {
		return err
	}

	gett, err := sub.DoGet()
	if err != nil {
		return err
	}

	sett, err := sub.DoSet()
	if err != nil {
		return err
	}

	const key = "x-keep-alive-at"
	gett.GetProperty(key)
	sett.SetProperty(key, now.String())

	err = sub.Update()
	if err != nil {
		return err
	}

	err = sub.Flush()
	if err != nil {
		return err
	}

	return nil
}

func (inst *innerSessionTask) doExit() error {
	ctx := inst.context
	sub, err := subjects.GetCurrent(ctx)
	if err != nil {
		return err
	}
	return sub.Exit()
}

////////////////////////////////////////////////////////////////////////////////
// EOF
