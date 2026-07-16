package iroles

import (
	"context"

	"github.com/starter-go/v0/rbac-web-app/app/classes/roles"
)

type RoleServiceImpl struct {

	//starter:component

	_as func(roles.Service) //starter:as("#")

	Dao roles.DAO //starter:inject("#")

}

// Insert implements roles.Service.
func (inst *RoleServiceImpl) Insert(c context.Context, item *roles.DTO) (*roles.DTO, error) {

	it2 := new(roles.Entity)
	it4 := new(roles.DTO)

	err := roles.ConvertD2E(item, it2)
	if err != nil {
		return nil, err
	}

	it3, err := inst.Dao.Insert(nil, it2)
	if err != nil {
		return nil, err
	}

	err = roles.ConvertE2D(it3, it4)
	return it4, err
}

// Find implements roles.Service.
func (inst *RoleServiceImpl) Find(c context.Context, id roles.ID) (*roles.DTO, error) {

	item1, err := inst.Dao.Find(nil, id)
	if err != nil {
		return nil, err
	}

	item2 := new(roles.DTO)
	err = roles.ConvertE2D(item1, item2)
	if err != nil {
		return nil, err
	}

	return item2, nil
}

// Query implements roles.Service.
func (inst *RoleServiceImpl) Query(c context.Context, q *roles.Query) ([]*roles.DTO, error) {

	list1, err := inst.Dao.Query(nil, q)
	if err != nil {
		return nil, err
	}

	list2 := make([]*roles.DTO, 0)
	return roles.ConvertListE2D(list1, list2)
}

func (inst *RoleServiceImpl) _impl() roles.Service {
	return inst
}
