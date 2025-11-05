package controllers

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

type ExampleController struct {

	//starter:component

	_as func(libgin.Controller) //starter:as(".")

	Responder libgin.Responder //starter:inject("#")

	Service users.Service //starter:inject("#")

}

func (inst *ExampleController) _impl() libgin.Controller {
	return inst
}

func (inst *ExampleController) Registration() *libgin.ControllerRegistration {
	r1 := new(libgin.ControllerRegistration)
	r1.Route = inst.route
	return r1
}

func (inst *ExampleController) route(rp libgin.RouterProxy) error {

	rp = rp.For("examples")

	rp.GET("", inst.handleGetMock)

	return nil
}

func (inst *ExampleController) handleGetMock(c *gin.Context) {

	task := new(innerExampleTask)
	task.context = c
	task.controller = inst

	task.execute(task.doGetMock)
}

////////////////////////////////////////////////////////////////////////////////

type innerExampleTask struct {
	context    *gin.Context
	controller *ExampleController

	wantRequestID   bool
	wantRequestBody bool

	id    dxo.ExampleID
	body1 vo.Example
	body2 vo.Example
}

func (inst *innerExampleTask) open(c *gin.Context) error {

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

func (inst *innerExampleTask) send(err error) {

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

func (inst *innerExampleTask) execute(fn func() error) {
	ctx := inst.context
	err := inst.open(ctx)
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *innerExampleTask) doGetMock() error {

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
