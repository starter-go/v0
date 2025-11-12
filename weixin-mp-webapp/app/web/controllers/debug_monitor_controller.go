package controllers

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/application"
	"github.com/starter-go/libgin"
	"github.com/starter-go/v0/weixin-mp-webapp/app/data/dxo"
	"github.com/starter-go/v0/weixin-mp-webapp/app/web/dto"
	"github.com/starter-go/v0/weixin-mp-webapp/app/web/vo"
	"github.com/starter-go/vlog"
)

// DebugMonitorController ...
type DebugMonitorController struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Sender libgin.Responder //starter:inject("#")

	Enabled       bool //starter:inject("${debug.weixin-mp.monitor.enabled}")
	IntervalInSec int  //starter:inject("${debug.weixin-mp.monitor.interval-in-sec}")

	isTimerActive bool
}

func (inst *DebugMonitorController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *DebugMonitorController) Registration() *libgin.ControllerRegistration {

	if inst.Enabled {
		return &libgin.ControllerRegistration{Route: inst.route}
	}

	return &libgin.ControllerRegistration{}
}

func (inst *DebugMonitorController) Life() *application.Life {
	l := new(application.Life)
	if inst.Enabled {
		l.OnStart = inst.startTimer
		l.OnStop = inst.stopTimer
	}
	return l
}

func (inst *DebugMonitorController) route(rp libgin.RouterProxy) error {

	rp = rp.For("debug/monitor")

	rp.GET("", inst.handle)

	// rp.POST("", inst.handle)
	// rp.PUT(":id", inst.handle)
	// rp.DELETE(":id", inst.handle)
	// rp.GET(":id", inst.handleGetOne)

	return nil
}

func (inst *DebugMonitorController) handle(c *gin.Context) {
	req := &myDebugMonitorRequest{
		context:       c,
		controller:    inst,
		wantRequestID: false,
	}
	req.execute(req.doNOP)
}

func (inst *DebugMonitorController) handleGetOne(c *gin.Context) {
	req := &myDebugMonitorRequest{
		context:       c,
		controller:    inst,
		wantRequestID: true,
	}
	req.execute(req.doGetOne)
}

func (inst *DebugMonitorController) startTimer() error {
	inst.isTimerActive = true
	go inst.runTimerLoop()
	return nil
}

func (inst *DebugMonitorController) stopTimer() error {
	inst.isTimerActive = false
	return nil
}

func (inst *DebugMonitorController) runTimerLoop() {

	interval := inst.IntervalInSec
	d := time.Second
	d = d * time.Duration(interval)

	vlog.Info("DebugMonitorController.onTimer(): started")

	for inst.isTimerActive {
		inst.onTimer()
		time.Sleep(d)
	}

	vlog.Info("DebugMonitorController.onTimer(): stopped")
}

func (inst *DebugMonitorController) onTimer() {
	vlog.Info("DebugMonitorController.onTimer()")
}

////////////////////////////////////////////////////////////////////////////////

type myDebugMonitorRequest struct {
	context    *gin.Context
	controller *DebugMonitorController

	wantRequestID   bool
	wantRequestBody bool

	id    dxo.ExampleID
	body1 vo.Examples
	body2 vo.Examples
}

func (inst *myDebugMonitorRequest) open() error {

	c := inst.context

	if inst.wantRequestID {
		idstr := c.Param("id")
		idnum, err := strconv.ParseInt(idstr, 10, 64)
		if err != nil {
			return err
		}
		inst.id = dxo.ExampleID(idnum)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myDebugMonitorRequest) send(err error) {
	data := &inst.body2
	code := inst.body2.Status
	resp := new(libgin.Response)
	resp.Context = inst.context
	resp.Error = err
	resp.Data = data
	resp.Status = code
	inst.controller.Sender.Send(resp)
}

func (inst *myDebugMonitorRequest) execute(fn func() error) {
	err := inst.open()
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *myDebugMonitorRequest) doNOP() error {
	return nil
}

func (inst *myDebugMonitorRequest) doGetOne() error {

	// id := inst.id
	// o1, err := inst.controller.Dao.Find(nil, id)
	// if err != nil {
	// 	return err
	// }

	o2 := &dto.Example{
		// ID:  o1.ID,
		// Foo: o1.Foo,
		// Bar: o1.Bar,
	}
	inst.body2.Items = []*dto.Example{o2}
	return nil
}
