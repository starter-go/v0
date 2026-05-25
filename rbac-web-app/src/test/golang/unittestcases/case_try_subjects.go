package unittestcases

import (
	"github.com/starter-go/application"
)

type CaseTrySubject struct {

	//starter:component

	_as func(application.Lifecycle) //starter:as(".")

	// Subjects subjects.Adapter //x-starter:inject("#")

}

func (inst *CaseTrySubject) Life() *application.Life {
	l := new(application.Life)
	// l.OnLoop = inst.run
	return l
}

func (inst *CaseTrySubject) run() error {

	// ctx := new(gin.Context)
	// ctx.Request = new(http.Request)
	// // ctx.Writer = new ( gin. )

	// webcontexts.SetupAccessForGinContext(ctx)

	// sub, err := subjects.Current(ctx)
	// if err != nil {
	// 	return err
	// }

	// err = inst.Subjects.InitSubject(ctx, sub)
	// if err != nil {
	// 	return err
	// }

	// sub.Create()
	// sub.Save()
	// sub.GetSession()
	// sub.Delete()

	// return nil

	panic("no impl")
}
