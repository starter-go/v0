package admin

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
	"github.com/starter-go/rbac"
	"github.com/starter-go/v0/rbac-web-app/app/classes/roles"
	"github.com/starter-go/v0/rbac-web-app/app/data/entity"
)

type AdminRoleController struct {

	//starter:component

	_as func(libgin.Controller) //starter:as(".")

	Responder libgin.Responder //starter:inject("#")

	Service roles.Service //starter:inject("#")

}

func (inst *AdminRoleController) _impl() libgin.Controller {
	return inst
}

func (inst *AdminRoleController) Registration() *libgin.ControllerRegistration {
	r1 := new(libgin.ControllerRegistration)
	r1.Route = inst.route
	return r1
}

func (inst *AdminRoleController) route(rp libgin.RouterProxy) error {

	rp = rp.For("admin/roles")

	rp.GET("", inst.handleGetQuery)
	rp.GET(":id", inst.handleGetOne)
	rp.GET("do/mock", inst.handleGetMock)
	rp.GET("do/query", inst.handleGetQuery)

	rp.POST("", inst.handlePostInsert)

	return nil
}

func (inst *AdminRoleController) handleGetMock(c *gin.Context) {

	task := new(innerAdminRoleTask)
	task.context = c
	task.controller = inst

	task.execute(task.doGetMock)
}

func (inst *AdminRoleController) handleGetOne(c *gin.Context) {

	task := new(innerAdminRoleTask)
	task.context = c
	task.controller = inst

	task.wantRequestID = true

	task.execute(task.doFindOneItem)
}

func (inst *AdminRoleController) handleGetQuery(c *gin.Context) {

	task := new(innerAdminRoleTask)
	task.context = c
	task.controller = inst

	task.wantRequestQuery = true

	task.execute(task.doGetQuery)
}

func (inst *AdminRoleController) handlePostInsert(c *gin.Context) {

	task := new(innerAdminRoleTask)

	task.context = c
	task.controller = inst
	task.wantRequestBody = true

	task.execute(task.doInsert)
}

////////////////////////////////////////////////////////////////////////////////

type innerAdminRoleTask struct {
	context    *gin.Context
	controller *AdminRoleController

	wantRequestID    bool
	wantRequestBody  bool
	wantRequestQuery bool

	id    roles.ID
	query roles.Query

	body1 roles.VO
	body2 roles.VO
}

func (inst *innerAdminRoleTask) open(c *gin.Context) error {

	if inst.wantRequestID {
		idstr := c.Param("id")
		num, err := strconv.ParseInt(idstr, 0, 0)
		if err != nil {
			return err
		}
		inst.id = roles.ID(num)
	}

	if inst.wantRequestQuery {
		q := new(roles.Query)
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

func (inst *innerAdminRoleTask) parseQuery(c *gin.Context, q *roles.Query) error {

	limit := int64(1)
	offset := int64(0)
	want := new(entity.Role)

	strLimit := c.Query("limit")
	strOffset := c.Query("offset")
	strWantID := c.Query("want-id")

	if strWantID != "" {
		num, err := strconv.ParseInt(strWantID, 10, 64)
		if err != nil {
			return err
		}
		want.ID = roles.ID(num)
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

func (inst *innerAdminRoleTask) send(err error) {

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

func (inst *innerAdminRoleTask) execute(fn func() error) {
	ctx := inst.context
	err := inst.open(ctx)
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *innerAdminRoleTask) doGetMock() error {

	it1 := new(roles.DTO)
	it2 := new(roles.DTO)

	it1.ID = 1
	it1.Name = "mock"

	it2.ID = 2
	it2.Name = rbac.RoleFriend

	list := inst.body2.Roles
	list = append(list, it1)
	list = append(list, it2)
	inst.body2.Roles = list

	return nil
}

func (inst *innerAdminRoleTask) doFindOneItem() error {

	ctx := inst.context
	ser := inst.controller.Service
	id := inst.id

	item, err := ser.Find(ctx, id)
	if err != nil {
		return err
	}

	page := new(rbac.Pagination)

	inst.body2.Pagination = page
	inst.body2.Roles = []*roles.DTO{item}
	return nil
}

func (inst *innerAdminRoleTask) doGetQuery() error {

	ctx := inst.context
	ser := inst.controller.Service
	q := &inst.query

	list1, err := ser.Query(ctx, q)
	if err != nil {
		return err
	}

	page := &q.Pagination

	inst.body2.Pagination = page
	inst.body2.Roles = list1
	return nil
}

func (inst *innerAdminRoleTask) doInsert() error {

	ctx := inst.context
	ser := inst.controller.Service
	src := inst.body1.Roles
	dst := inst.body2.Roles

	for _, item1 := range src {
		item2, err := ser.Insert(ctx, item1)
		if err != nil {
			return err
		}
		dst = append(dst, item2)
	}

	inst.body2.Roles = dst
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// EOF
