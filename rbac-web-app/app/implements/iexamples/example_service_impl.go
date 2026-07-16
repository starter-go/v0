package iexamples

import (
	"context"

	"github.com/starter-go/v0/rbac-web-app/app/classes/examples"
)

type ExampleServiceImpl struct {

	//starter:component

	_as func(examples.Service) //starter:as("#")

	Dao examples.DAO //starter:inject("#")

}

// Find implements examples.Service.
func (inst *ExampleServiceImpl) Find(c context.Context, id examples.ID) (*examples.DTO, error) {
	panic("unimplemented")
}

// Query implements examples.Service.
func (inst *ExampleServiceImpl) Query(c context.Context, q *examples.Query) ([]*examples.DTO, error) {
	panic("unimplemented")
}

func (inst *ExampleServiceImpl) _impl() examples.Service {
	return inst
}
