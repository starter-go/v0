package ipermissions

import (
	"context"
	"net/http"

	"github.com/starter-go/rbac"
	"github.com/starter-go/v0/rbac-web-app/app/classes/permissions"
	"github.com/starter-go/vlog"
	"gorm.io/gorm"
)

type PermissionServiceImpl struct {

	//starter:component

	_as func(permissions.Service) //starter:as("#")

	Dao permissions.DAO //starter:inject("#")

}

// Setup implements permissions.Service.
func (inst *PermissionServiceImpl) Setup(cc context.Context) ([]*permissions.DTO, error) {

	src := make([]*permissions.DTO, 0)
	dst := make([]*permissions.DTO, 0)

	src = append(src, &rbac.PermissionDTO{
		Method:  http.MethodGet,
		Path:    "/a/b/c1",
		Roles:   rbac.NewRoleNameList(rbac.RoleAdmin, rbac.RoleUser),
		Enabled: true,
	})

	src = append(src, &rbac.PermissionDTO{
		Method:  http.MethodGet,
		Path:    "/a/b/c2",
		Roles:   rbac.NewRoleNameList(rbac.RoleAdmin, rbac.RoleUser),
		Enabled: true,
	})

	src = append(src, &rbac.PermissionDTO{
		Method:  http.MethodPost,
		Path:    "/a/b/c3",
		Roles:   rbac.NewRoleNameList(rbac.RoleAdmin, rbac.RoleUser),
		Enabled: true,
	})

	for _, it1 := range src {
		it2, err := inst.Insert(cc, it1)
		if err != nil {
			vlog.Warn("%s", err.Error())
			continue
		}
		dst = append(dst, it2)
	}

	return dst, nil
}

// Find implements permissions.Service.
func (inst *PermissionServiceImpl) Find(cc context.Context, id permissions.ID) (*permissions.DTO, error) {

	item1, err := inst.Dao.Find(nil, id)
	if err != nil {
		return nil, err
	}

	item2 := new(permissions.DTO)
	err = permissions.ConvertE2D(item1, item2)
	return item2, err
}

// Insert implements permissions.Service.
func (inst *PermissionServiceImpl) Insert(cc context.Context, item *permissions.DTO) (*permissions.DTO, error) {

	var db *gorm.DB = nil
	dao := inst.Dao
	item2 := new(permissions.Entity)
	item4 := new(permissions.DTO)

	err := permissions.ConvertD2E(item, item2)
	if err != nil {
		return nil, err
	}

	item3, err := dao.Insert(db, item2)
	if err != nil {
		return nil, err
	}

	err = permissions.ConvertE2D(item3, item4)
	return item4, err
}

// Query implements permissions.Service.
func (inst *PermissionServiceImpl) Query(cc context.Context, q *permissions.Query) ([]*permissions.DTO, error) {
	list1, err := inst.Dao.Query(nil, q)
	if err != nil {
		return nil, err
	}
	return permissions.ConvertListE2D(list1, nil)
}

func (inst *PermissionServiceImpl) _impl() permissions.Service {
	return inst
}
