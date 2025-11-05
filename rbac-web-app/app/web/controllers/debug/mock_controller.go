package debug

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/libgin"
	"github.com/starter-go/vlog"

	"github.com/starter-go/v0/rbac-web-app/app/classes/sessions"
	"github.com/starter-go/v0/rbac-web-app/app/classes/subjects"
	"github.com/starter-go/v0/rbac-web-app/app/classes/tokens"
	"github.com/starter-go/v0/rbac-web-app/app/classes/users"
	"github.com/starter-go/v0/rbac-web-app/app/data/dxo"
	"github.com/starter-go/v0/rbac-web-app/app/web/dto"
	"github.com/starter-go/v0/rbac-web-app/app/web/vo"
)

type MockController struct {

	//starter:component

	_as func(libgin.Controller) //starter:as(".")

	Responder libgin.Responder //starter:inject("#")

	TokenService   tokens.Service   //starter:inject("#")
	SessionService sessions.Service //starter:inject("#")
	UserService    users.Service    //starter:inject("#")

}

func (inst *MockController) _impl() libgin.Controller {
	return inst
}

func (inst *MockController) Registration() *libgin.ControllerRegistration {
	r1 := new(libgin.ControllerRegistration)
	r1.Route = inst.route
	return r1
}

func (inst *MockController) route(rp libgin.RouterProxy) error {

	rp = rp.For("mocks")

	// rp.GET("", inst.handleGetMock)

	rp.GET("subject", inst.handleGetSubject)

	rp.POST("login", inst.handlePostLogin)
	rp.POST("init-session", inst.handlePostInitSession)

	return nil
}

func (inst *MockController) handleGetMock(c *gin.Context) {

	task := new(innerMockTask)
	task.context = c
	task.controller = inst

	task.execute(task.doGetMock)
}

func (inst *MockController) handleGetSubject(c *gin.Context) {

	task := new(innerMockTask)
	task.context = c
	task.controller = inst

	task.execute(task.doGetSubject)
}

func (inst *MockController) handlePostLogin(c *gin.Context) {

	task := new(innerMockTask)
	task.context = c
	task.controller = inst

	task.execute(task.doLogin)
}

func (inst *MockController) handlePostInitSession(c *gin.Context) {

	task := new(innerMockTask)
	task.context = c
	task.controller = inst

	task.execute(task.doInitSession)
}

////////////////////////////////////////////////////////////////////////////////

type innerMockTask struct {
	context    *gin.Context
	controller *MockController

	wantRequestID   bool
	wantRequestBody bool

	id    dxo.ExampleID
	body1 vo.Example
	body2 vo.Example
}

func (inst *innerMockTask) open(c *gin.Context) error {

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

func (inst *innerMockTask) send(err error) {

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

func (inst *innerMockTask) execute(fn func() error) {
	ctx := inst.context
	err := inst.open(ctx)
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *innerMockTask) doGetMock() error {

	return nil
}

func (inst *innerMockTask) doGetSubject() error {

	cc := inst.context
	sub, err := subjects.Current(cc)
	if err != nil {
		return err
	}

	user, err := sub.GetUser()
	if err != nil {
		return err
	}

	uid := user.ID
	vlog.Info("subjects.current.uid=%d", uid)

	return nil
}

func (inst *innerMockTask) doLogin() error {

	return nil
}

func (inst *innerMockTask) doInitSession() error {

	ctx := inst.context
	ser1 := inst.controller.TokenService
	ser2 := inst.controller.SessionService
	ser3 := inst.controller.UserService

	now := lang.Now()
	nowStr := strconv.FormatInt(now.Int(), 10)

	session := new(dto.Session)
	token := new(dto.Token)
	user := new(dto.User)

	// user

	user.Name = "user_" + dxo.UserName(nowStr)
	user.Email = "mail_" + dxo.EmailAddress(nowStr)
	user.Mobile = "phone_" + dxo.PhoneNumber(nowStr)

	user2, err := ser3.Insert(ctx, user)
	if err != nil {
		return err
	}

	// session

	session.DisplayName = user2.DisplayName
	session.Avatar = user2.Avatar
	session.Owner = user2.ID
	session.AliveFrom = now
	session.AliveTo = now + 3600*1000

	sess2, err := ser2.Insert(ctx, session)
	if err != nil {
		return err
	}

	// token

	token.SessionID = sess2.ID
	token.SessionUUID = sess2.UUID
	token2, err := ser1.SetCurrentToken(ctx, token)
	if err != nil {
		return err
	}
	vlog.Info("token2 = %v", token2)

	return nil
}

////////////////////////////////////////////////////////////////////////////////
// EOF
