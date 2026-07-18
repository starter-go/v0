package admin

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
	"github.com/starter-go/rbac"
	"github.com/starter-go/v0/rbac-web-app/app/classes/permissions"
	"github.com/starter-go/v0/rbac-web-app/app/data/dxo"
	"github.com/starter-go/v0/rbac-web-app/app/web/dto"
)

type AdminPermissionController struct {

	//starter:component

	_as func(libgin.Controller) //starter:as(".")

	Responder libgin.Responder //starter:inject("#")

	Service permissions.Service //starter:inject("#")

}

func (inst *AdminPermissionController) _impl() libgin.Controller {
	return inst
}

func (inst *AdminPermissionController) Registration() *libgin.ControllerRegistration {
	r1 := new(libgin.ControllerRegistration)
	r1.Route = inst.route
	return r1
}

func (inst *AdminPermissionController) route(rp libgin.RouterProxy) error {

	rp = rp.For("admin/permissions")

	rp.GET("", inst.handleGetQuery)
	rp.GET(":id", inst.handleGetOne)
	rp.GET("do/query", inst.handleGetQuery)
	rp.GET("do/mock", inst.handleGetMock)

	rp.POST("", inst.handlePostInsert)
	rp.POST("do/insert", inst.handlePostInsert)
	rp.POST("do/reload", inst.handlePostReload)
	rp.POST("do/setup", inst.handlePostSetup)

	return nil
}

func (inst *AdminPermissionController) handleGetMock(c *gin.Context) {

	task := new(innerAdminPermissionTask)
	task.context = c
	task.controller = inst

	task.execute(task.doGetMock)
}

func (inst *AdminPermissionController) handleGetOne(c *gin.Context) {

	task := new(innerAdminPermissionTask)
	task.context = c
	task.controller = inst

	task.wantRequestID = true

	task.execute(task.doFindOneItem)
}

func (inst *AdminPermissionController) handleGetQuery(c *gin.Context) {

	task := new(innerAdminPermissionTask)
	task.context = c
	task.controller = inst

	task.wantRequestQuery = true

	task.execute(task.doGetQuery)
}

func (inst *AdminPermissionController) handlePostSetup(c *gin.Context) {

	task := new(innerAdminPermissionTask)
	task.context = c
	task.controller = inst

	// task.wantRequestQuery = true

	task.execute(task.doSetup)
}

func (inst *AdminPermissionController) handlePostInsert(c *gin.Context) {

	task := new(innerAdminPermissionTask)

	task.context = c
	task.controller = inst
	task.wantRequestBody = true

	task.execute(task.doInsert)
}

func (inst *AdminPermissionController) handlePostReload(c *gin.Context) {

	task := new(innerAdminPermissionTask)

	task.context = c
	task.controller = inst
	task.wantRequestBody = false

	task.execute(task.doReload)
}

////////////////////////////////////////////////////////////////////////////////

type innerAdminPermissionTask struct {
	context    *gin.Context
	controller *AdminPermissionController

	wantRequestID    bool
	wantRequestBody  bool
	wantRequestQuery bool

	id    permissions.ID
	query permissions.Query

	body1 permissions.VO
	body2 permissions.VO
}

func (inst *innerAdminPermissionTask) open(c *gin.Context) error {

	if inst.wantRequestID {
		idstr := c.Param("id")
		num, err := strconv.ParseInt(idstr, 0, 0)
		if err != nil {
			return err
		}
		inst.id = dxo.PermissionID(num)
	}

	if inst.wantRequestQuery {
		q := new(permissions.Query)
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

func (inst *innerAdminPermissionTask) parseQuery(c *gin.Context, q *permissions.Query) error {

	limit := int64(1)
	offset := int64(0)
	want := new(permissions.Entity)

	strLimit := c.Query("limit")
	strOffset := c.Query("offset")
	strWantID := c.Query("want-id")

	if strWantID != "" {
		num, err := strconv.ParseInt(strWantID, 10, 64)
		if err != nil {
			return err
		}
		want.ID = dxo.PermissionID(num)
	}

	if strLimit != "" {
		num, err := strconv.ParseInt(strLimit, 10, 64)
		if err != nil {
			return err
		}
		limit = num
	}

	if strOffset != "" {
		num, err := strconv.ParseInt(strOffset, 10, 64)
		if err != nil {
			return err
		}
		offset = num
	}

	if offset < 0 {
		offset = 0
	}
	if limit < 1 {
		limit = 1
	}

	q.Pagination.Size = int(limit)
	q.Pagination.Page = (offset / limit) + 1
	q.Want = want
	return nil
}

func (inst *innerAdminPermissionTask) send(err error) {

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

func (inst *innerAdminPermissionTask) execute(fn func() error) {
	ctx := inst.context
	err := inst.open(ctx)
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *innerAdminPermissionTask) doGetMock() error {

	it1 := new(permissions.DTO)
	it2 := new(permissions.DTO)

	it1.ID = 1
	it2.ID = 2

	list := inst.body2.Permissions
	list = append(list, it1)
	list = append(list, it2)
	inst.body2.Permissions = list

	return nil
}

func (inst *innerAdminPermissionTask) doFindOneItem() error {

	ctx := inst.context
	ser := inst.controller.Service
	id := inst.id

	item, err := ser.Find(ctx, id)
	if err != nil {
		return err
	}

	page := new(rbac.Pagination)

	inst.body2.Pagination = page
	inst.body2.Permissions = []*dto.Permission{item}
	return nil
}

func (inst *innerAdminPermissionTask) doGetQuery() error {

	ctx := inst.context
	ser := inst.controller.Service
	q := &inst.query

	list1, err := ser.Query(ctx, q)
	if err != nil {
		return err
	}

	page := &q.Pagination

	inst.body2.Pagination = page
	inst.body2.Permissions = list1
	return nil
}

func (inst *innerAdminPermissionTask) doSetup() error {

	ctx := inst.context
	ser := inst.controller.Service

	list, err := ser.Setup(ctx)
	if err != nil {
		return err
	}

	page := new(rbac.Pagination)
	count := len(list)
	page.Total = int64(count)

	inst.body2.Pagination = page
	inst.body2.Permissions = list

	return nil
}

func (inst *innerAdminPermissionTask) doInsert() error {

	ctx := inst.context
	ser := inst.controller.Service
	src := inst.body1.Permissions
	dst := inst.body2.Permissions

	for _, item1 := range src {
		item2, err := ser.Insert(ctx, item1)
		if err != nil {
			return err
		}
		dst = append(dst, item2)
	}

	inst.body2.Permissions = dst
	return nil
}

func (inst *innerAdminPermissionTask) doReload() error {
	ctx := inst.context
	ser := inst.controller.Service
	return ser.Reload(ctx)
}

////////////////////////////////////////////////////////////////////////////////
// EOF
