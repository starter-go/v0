package admin

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
	"github.com/starter-go/rbac"
	"github.com/starter-go/v0/rbac-web-app/app/classes/users"
	"github.com/starter-go/v0/rbac-web-app/app/data/dxo"
	"github.com/starter-go/v0/rbac-web-app/app/data/entity"
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

	rp.GET(":id", inst.handleGetOne)
	rp.GET("mock", inst.handleGetMock)
	rp.GET("query", inst.handleGetQuery)

	return nil
}

func (inst *UsersController) handleGetMock(c *gin.Context) {

	task := new(innerUsersTask)
	task.context = c
	task.controller = inst

	task.execute(task.doGetMock)
}

func (inst *UsersController) handleGetOne(c *gin.Context) {

	task := new(innerUsersTask)
	task.context = c
	task.controller = inst

	task.wantRequestID = true

	task.execute(task.doFindOneItem)
}

func (inst *UsersController) handleGetQuery(c *gin.Context) {

	task := new(innerUsersTask)
	task.context = c
	task.controller = inst

	task.wantRequestQuery = true

	task.execute(task.doGetQuery)
}

////////////////////////////////////////////////////////////////////////////////

type innerUsersTask struct {
	context    *gin.Context
	controller *UsersController

	wantRequestID    bool
	wantRequestBody  bool
	wantRequestQuery bool

	id    dxo.UserID
	query *users.Query

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

	if inst.wantRequestQuery {
		q := new(users.Query)
		err := inst.parseQuery(c, q)
		if err != nil {
			return err
		}
		inst.query = q
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

func (inst *innerUsersTask) parseQuery(c *gin.Context, q *users.Query) error {

	limit := int64(1)
	offset := int64(0)
	want := new(entity.User)

	strLimit := c.Query("limit")
	strOffset := c.Query("offset")
	strWantID := c.Query("want-id")
	strWantEmail := c.Query("want-email")
	strWantPhone := c.Query("want-phone")
	strWantName := c.Query("want-name")
	strWantNickName := c.Query("want-nickname")

	if strWantID != "" {
		num, err := strconv.ParseInt(strWantID, 10, 64)
		if err != nil {
			return err
		}
		want.ID = dxo.UserID(num)
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

	if strWantPhone != "" {
		want.Mobile = dxo.PhoneNumber(strWantPhone)
	}

	if strWantEmail != "" {
		want.Email = dxo.EmailAddress(strWantEmail)
	}

	if strWantNickName != "" {
		want.DisplayName = strWantNickName
	}

	if strWantName != "" {
		want.Name = dxo.UserName(strWantName)
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

func (inst *innerUsersTask) doFindOneItem() error {

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

func (inst *innerUsersTask) doGetQuery() error {

	ctx := inst.context
	ser := inst.controller.Service
	q := inst.query

	list1, err := ser.Query(ctx, q)
	if err != nil {
		return err
	}

	page := &q.Pagination

	inst.body2.Pagination = page
	inst.body2.Items = list1
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// EOF
