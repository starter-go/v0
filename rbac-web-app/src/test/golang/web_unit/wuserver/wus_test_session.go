package wuserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
	"github.com/starter-go/v0/rbac-web-app/app/classes/subjects"
	"github.com/starter-go/v0/rbac-web-app/app/classes/users"
	"github.com/starter-go/v0/rbac-web-app/app/classes/webcontexts"
)

type WebUnitTestSessionController struct {

	//starter:component

	_as func(libgin.Controller) //starter:as(".")

	Responder libgin.Responder //starter:inject("#")

	Service users.Service //starter:inject("#")

}

func (inst *WebUnitTestSessionController) _impl() libgin.Controller {
	return inst
}

func (inst *WebUnitTestSessionController) Registration() *libgin.ControllerRegistration {
	r1 := new(libgin.ControllerRegistration)
	r1.Route = inst.route
	return r1
}

func (inst *WebUnitTestSessionController) route(rp libgin.RouterProxy) error {

	rp = rp.For("/test/web-unit/sessions")

	rp.GET("fetch", inst.handleFetchSession)
	rp.GET("insert", inst.handleInsertSession)
	rp.GET("update", inst.handleUpdateSession)
	rp.GET("remove", inst.handleRemoveSession)

	return nil
}

func (inst *WebUnitTestSessionController) setupWebContext(c *gin.Context) {
	webcontexts.SetupAccessForGinContext(c)
}

func (inst *WebUnitTestSessionController) handleFetchSession(c *gin.Context) {

	inst.setupWebContext(c)

	code := http.StatusOK
	ctype := "text/plain"
	str := "hello, web-unit"

	c.Data(code, ctype, []byte(str))
}

func (inst *WebUnitTestSessionController) handleInsertSession(c *gin.Context) {

	inst.setupWebContext(c)

	sub, err := subjects.Current(c)
	if err != nil {
		c.AbortWithError(0, err)
		return
	}

	sub.SetProperty("a", "11")
	sub.SetProperty("b", "22")

	err = sub.Create()
	if err != nil {
		c.AbortWithError(0, err)
		return
	}

	code := http.StatusOK
	ctype := "text/plain"
	str := "hello, web-unit"

	c.Data(code, ctype, []byte(str))
}

func (inst *WebUnitTestSessionController) handleUpdateSession(c *gin.Context) {

	inst.setupWebContext(c)

	code := http.StatusOK
	ctype := "text/plain"
	str := "hello, web-unit"

	c.Data(code, ctype, []byte(str))
}

func (inst *WebUnitTestSessionController) handleRemoveSession(c *gin.Context) {

	inst.setupWebContext(c)

	code := http.StatusOK
	ctype := "text/plain"
	str := "hello, web-unit"

	c.Data(code, ctype, []byte(str))
}
