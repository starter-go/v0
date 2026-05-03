package unittestcases

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/application"
	"github.com/starter-go/v0/rbac-web-app/app/classes/subjects"
	"github.com/starter-go/v0/rbac-web-app/app/classes/webcontexts"
)

type CaseTrySubject struct {

	//starter:component

	_as func(application.Lifecycle) //starter:as(".")

	Subjects subjects.Service //starter:inject("#")

}

func (inst *CaseTrySubject) Life() *application.Life {
	l := new(application.Life)
	// l.OnLoop = inst.run
	return l
}

func (inst *CaseTrySubject) run() error {

	ctx := new(gin.Context)
	ctx.Request = new(http.Request)
	// ctx.Writer = new ( gin. )

	webcontexts.SetupAccessForGinContext(ctx)

	sub, err := subjects.Current(ctx)
	if err != nil {
		return err
	}

	err = inst.Subjects.InitSubject(ctx, sub)
	if err != nil {
		return err
	}

	sub.Create()
	sub.Save()
	sub.GetSession()
	sub.Delete()

	return nil
}
