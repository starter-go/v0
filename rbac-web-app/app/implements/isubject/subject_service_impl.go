package isubject

import (
	"context"

	"github.com/starter-go/v0/rbac-web-app/app/classes/sessions"
	"github.com/starter-go/v0/rbac-web-app/app/classes/statestores"
	"github.com/starter-go/v0/rbac-web-app/app/classes/subjects"
	"github.com/starter-go/v0/rbac-web-app/app/classes/tokens"
	"github.com/starter-go/v0/rbac-web-app/app/classes/users"
)

type SubjectServiceImpl struct {

	//starter:component

	_as func(subjects.Service) //starter:as("#")

	Tokens   tokens.Service      //starter:inject("#")
	Users    users.Service       //starter:inject("#")
	Sessions sessions.Service    //starter:inject("#")
	SSS      statestores.Service //starter:inject("#")

}

// InitContext implements subjects.Service.
func (inst *SubjectServiceImpl) InitContext(c1 context.Context, c2 *subjects.Context) error {

	sto, err := inst.SSS.GetStore(c1)
	if err != nil {
		return err
	}

	c2.Parent = c1
	c2.Services.Subjects = inst
	c2.Services.Sessions = inst.Sessions
	c2.Services.Tokens = inst.Tokens
	c2.Services.Users = inst.Users
	c2.Store = sto

	return nil
}

// InitSubject implements subjects.Service.
func (inst *SubjectServiceImpl) InitSubject(c1 context.Context, sub subjects.Subject) error {
	c2 := sub.GetContext()
	return inst.InitContext(c1, c2)
}

func (inst *SubjectServiceImpl) _impl() subjects.Service {
	return inst
}
