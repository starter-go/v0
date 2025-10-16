package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
	"github.com/starter-go/v0/web-app-template/app/data/dxo"
	"github.com/starter-go/v0/web-app-template/app/web/dto"
	"github.com/starter-go/v0/web-app-template/app/web/vo"
)

////////////////////////////////////////////////////////////////////////////////

type ExampleController struct {

	//starter:component

	_as func(libgin.Controller) //starter:as(".")

	Sender libgin.Responder //starter:inject("#")
}

func (inst *ExampleController) _impl() libgin.Controller {
	return inst
}

func (inst *ExampleController) Registration() *libgin.ControllerRegistration {
	cr1 := new(libgin.ControllerRegistration)
	cr1.Route = inst.route
	return cr1
}

func (inst *ExampleController) route(rp libgin.RouterProxy) error {

	rp = rp.For("examples")

	rp.GET("", inst.handleGetList)
	rp.GET(":id", inst.handleGetOne)

	return nil
}

func (inst *ExampleController) handleGetOne(gc *gin.Context) {

	req := new(myExampleRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = false
	req.wantRequestBody = false

	req.execute(req.doGetOne)
}

func (inst *ExampleController) handleGetList(gc *gin.Context) {

	req := new(myExampleRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = false
	req.wantRequestBody = false

	req.execute(req.doGetList)
}

////////////////////////////////////////////////////////////////////////////////

type myExampleRequest struct {
	wantRequestID   bool
	wantRequestBody bool

	context    *gin.Context
	controller *ExampleController

	id    dxo.ExampleID
	body1 vo.Examples
	body2 vo.Examples
}

func (inst *myExampleRequest) open(ctx *gin.Context) error {

	// inst.context = ctx
	// inst.controller = ctr

	if inst.wantRequestID {
	}

	if inst.wantRequestBody {
	}

	return nil
}

func (inst *myExampleRequest) execute(task func() error) {

	ex := new(libgin.Executor)
	ex.Context = inst.context
	ex.Responder = inst.controller.Sender
	ex.Body1 = &inst.body1
	ex.Body2 = &inst.body2

	ex.OnOpen = inst.open
	ex.OnTask = task

	ex.Execute()
}

func (inst *myExampleRequest) doGetList() error {

	it := &dto.Example{}

	inst.body2.Items = []*dto.Example{it, it, it}
	return nil
}

func (inst *myExampleRequest) doGetOne() error {

	it := &dto.Example{}

	inst.body2.Items = []*dto.Example{it}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// EOF
