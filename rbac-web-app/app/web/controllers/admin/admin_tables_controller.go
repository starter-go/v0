package admin

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
	"github.com/starter-go/rbac"
	"github.com/starter-go/v0/rbac-web-app/app/classes/tables"
)

type AdminTableController struct {

	//starter:component

	_as func(libgin.Controller) //starter:as(".")

	Responder libgin.Responder //starter:inject("#")

	Service tables.Service //starter:inject("#")

}

func (inst *AdminTableController) _impl() libgin.Controller {
	return inst
}

func (inst *AdminTableController) Registration() *libgin.ControllerRegistration {
	r1 := new(libgin.ControllerRegistration)
	r1.Route = inst.route
	r1.Groups = nil //  []string{"disabled"}
	return r1
}

func (inst *AdminTableController) route(rp libgin.RouterProxy) error {

	rp = rp.For("admin/tables")

	rp.GET(":id", inst.handleGetOne)
	rp.GET("do/mock", inst.handleGetMock)
	rp.GET("do/query", inst.handleGetQuery)

	rp.POST("do/scan", inst.handlePostScan)

	return nil
}

func (inst *AdminTableController) handleGetMock(c *gin.Context) {

	task := new(innerAdminTableTask)
	task.context = c
	task.controller = inst

	task.execute(task.doGetMock)
}

func (inst *AdminTableController) handleGetOne(c *gin.Context) {

	task := new(innerAdminTableTask)
	task.context = c
	task.controller = inst

	task.wantRequestID = true

	task.execute(task.doFindOneItem)
}

func (inst *AdminTableController) handleGetQuery(c *gin.Context) {

	task := new(innerAdminTableTask)
	task.context = c
	task.controller = inst

	task.wantRequestQuery = true

	task.execute(task.doGetQuery)
}

func (inst *AdminTableController) handlePostScan(c *gin.Context) {

	task := new(innerAdminTableTask)
	task.context = c
	task.controller = inst
	task.wantRequestID = false

	task.execute(task.doScan)
}

////////////////////////////////////////////////////////////////////////////////

type innerAdminTableTask struct {
	context    *gin.Context
	controller *AdminTableController

	wantRequestID    bool
	wantRequestBody  bool
	wantRequestQuery bool

	id    tables.ID
	query tables.Query
	body1 tables.VO
	body2 tables.VO
}

func (inst *innerAdminTableTask) open(c *gin.Context) error {

	if inst.wantRequestID {
		idstr := c.Param("id")
		num, err := strconv.ParseInt(idstr, 0, 0)
		if err != nil {
			return err
		}
		inst.id = tables.ID(num)
	}

	if inst.wantRequestQuery {
		q := new(tables.Query)
		err := inst.parseQuery(c, q)
		if err != nil {
			return err
		}
		inst.query = *q
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

func (inst *innerAdminTableTask) parseQuery(c *gin.Context, q *tables.Query) error {

	// todo ... no impl
	return nil
}

func (inst *innerAdminTableTask) send(err error) {

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

func (inst *innerAdminTableTask) execute(fn func() error) {
	ctx := inst.context
	err := inst.open(ctx)
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *innerAdminTableTask) doGetMock() error {

	ctx := inst.context
	ser := inst.controller.Service
	id := inst.id

	item, err := ser.Find(ctx, id)
	if err != nil {
		return err
	}

	page := new(rbac.Pagination)

	inst.body2.Pagination = page
	inst.body2.Items = []*tables.DTO{item}
	return nil
}

func (inst *innerAdminTableTask) doFindOneItem() error {

	ctx := inst.context
	ser := inst.controller.Service
	id := inst.id

	item, err := ser.Find(ctx, id)
	if err != nil {
		return err
	}

	page := new(rbac.Pagination)

	inst.body2.Pagination = page
	inst.body2.Items = []*tables.DTO{item}
	return nil
}

func (inst *innerAdminTableTask) doGetQuery() error {

	ctx := inst.context
	ser := inst.controller.Service
	q := &inst.query

	list1, err := ser.Query(ctx, q)
	if err != nil {
		return err
	}

	page := &q.Pagination

	inst.body2.Pagination = page
	inst.body2.Items = list1
	return nil
}

func (inst *innerAdminTableTask) doScan() error {

	ctx := inst.context
	ser := inst.controller.Service

	list, err := ser.ScanCurrentRuntime(ctx)
	if err != nil {
		return err
	}

	inst.body2.Items = list
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// EOF
