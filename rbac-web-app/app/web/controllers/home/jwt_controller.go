package home

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
	"github.com/starter-go/rbac"

	"github.com/starter-go/v0/rbac-web-app/app/classes/sessions"
	"github.com/starter-go/v0/rbac-web-app/app/classes/subjects"
	"github.com/starter-go/v0/rbac-web-app/app/classes/tokens"
	"github.com/starter-go/v0/rbac-web-app/app/classes/users"
	"github.com/starter-go/v0/rbac-web-app/app/classes/webcontexts"
	"github.com/starter-go/v0/rbac-web-app/app/data/dxo"
	"github.com/starter-go/v0/rbac-web-app/app/web/dto"
	"github.com/starter-go/v0/rbac-web-app/app/web/vo"
)

type JWTokenController struct {

	//starter:component

	_as func(libgin.Controller) //starter:as(".")

	Responder libgin.Responder //starter:inject("#")

	TokenService   tokens.Service   //starter:inject("#")
	SessionService sessions.Service //starter:inject("#")
	UserService    users.Service    //starter:inject("#")

}

func (inst *JWTokenController) _impl() libgin.Controller {
	return inst
}

func (inst *JWTokenController) Registration() *libgin.ControllerRegistration {
	r1 := new(libgin.ControllerRegistration)
	r1.Route = inst.route
	return r1
}

func (inst *JWTokenController) route(rp libgin.RouterProxy) error {

	rp = rp.For("tokens")

	// rp.GET("", inst.handleGetMock)

	inst.setupMiddleware(rp)

	return nil
}

func (inst *JWTokenController) setupMiddleware(rp libgin.RouterProxy) {

	r1 := new(libgin.Routing)
	fnlist := r1.Handlers

	fnlist = append(fnlist, inst.handlePrepareRbacContext)

	r1.Method = ""
	r1.Path = "/api/*"
	r1.Priority = 1
	r1.Middleware = true
	r1.Handlers = fnlist

	rp.Route(r1)
}

func (inst *JWTokenController) handleGetMock(c *gin.Context) {

	task := new(innerJWTokenTask)
	task.context = c
	task.controller = inst

	task.execute(task.doGetMock)
}

func (inst *JWTokenController) handlePrepareRbacContext(c *gin.Context) {
	err := inst.doPrepareRbacContext(c)
	if err != nil {
		const code = http.StatusInternalServerError
		c.AbortWithError(code, err)
	}
}

func (inst *JWTokenController) doPrepareRbacContext(c *gin.Context) error {

	webcontexts.SetupAccessForGinContext(c)

	sub, err := subjects.Current(c)
	if err != nil {
		return err
	}

	ctx := sub.GetContext()
	ctx.Services.Tokens = inst.TokenService
	ctx.Services.Sessions = inst.SessionService
	ctx.Services.Users = inst.UserService
	ctx.Parent = c

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type innerJWTokenTask struct {
	context    *gin.Context
	controller *JWTokenController

	wantRequestID   bool
	wantRequestBody bool

	id    dxo.ExampleID
	body1 vo.Example
	body2 vo.Example
}

func (inst *innerJWTokenTask) open(c *gin.Context) error {

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

	return nil
}

func (inst *innerJWTokenTask) send(err error) {

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

func (inst *innerJWTokenTask) execute(fn func() error) {
	ctx := inst.context
	err := inst.open(ctx)
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *innerJWTokenTask) doGetMock() error {

	// ctx := inst.context
	// ser := inst.controller.Service
	// id := inst.id

	item := new(dto.Example)
	page := new(rbac.Pagination)

	inst.body2.Pagination = page
	inst.body2.Items = []*dto.Example{item}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// EOF
