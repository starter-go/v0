package ipermissions

import (
	perm1 "github.com/starter-go/security/permissions"
	perm2 "github.com/starter-go/v0/rbac-web-app/app/classes/permissions"
	"github.com/starter-go/vlog"
)

type DatabasePermissionLoader struct {

	//starter:component

	_as func(perm1.Registry) //starter:as(".")

	Dao perm2.DAO //starter:inject("#")

}

// ListRegistrations implements [permissions.Registry].
func (inst *DatabasePermissionLoader) ListRegistrations() []*perm1.Registration {
	all, err := inst.innerLoadRegistrations()
	if err != nil {
		panic(err)
	}

	inst.innerLogItems(all)

	return all
}

func (inst *DatabasePermissionLoader) innerConvertItem(src *perm2.Entity, dst *perm1.Registration) error {

	dst.Enabled = src.Enabled
	dst.Method = src.Method
	dst.Path = src.Path
	dst.Priority = src.Priority
	dst.Roles = src.Roles
	dst.URI = src.URI

	return nil
}

func (inst *DatabasePermissionLoader) innerLoadRegistrations() ([]*perm1.Registration, error) {

	src, err := inst.innerLoadEntities()
	if err != nil {
		return nil, err
	}

	dst := make([]*perm1.Registration, 0)
	for _, it1 := range src {

		if !inst.innerIsItemReady(it1) {
			continue
		}

		it2 := new(perm1.Registration)
		err = inst.innerConvertItem(it1, it2)
		if err != nil {
			return nil, err
		}
		dst = append(dst, it2)
	}

	return dst, err
}

func (inst *DatabasePermissionLoader) innerLogItems(list []*perm1.Registration) {

	if !vlog.IsDebugEnabled() {
		return
	}

	const f = "[permission index:%d uri:'%s' method:'%s' path:'%s' enabled:%v priority:%d roles:'%s']"

	for index, item := range list {

		uri := item.URI
		method := item.Method
		path := item.Path
		roles := item.Roles
		priority := item.Priority
		enabled := item.Enabled

		vlog.Debug(f, index, uri, method, path, enabled, priority, roles)
	}
}

func (inst *DatabasePermissionLoader) innerIsItemReady(item *perm2.Entity) bool {

	if item == nil {
		return false
	}

	// if !item.Enabled {
	// 	return false
	// }

	if item.ID == 0 {
		return false
	}

	const empty = ""

	if item.Method == empty || item.Path == empty || item.Roles == empty {
		return false
	}

	return true
}

func (inst *DatabasePermissionLoader) innerLoadEntities() ([]*perm2.Entity, error) {

	q := new(perm2.Query)

	q.All = true
	q.Pagination.Total = 0
	q.Pagination.Size = 999
	q.Pagination.Page = 1

	return inst.Dao.Query(nil, q)
}

// Setup implements permissions.Service.
func (inst *DatabasePermissionLoader) _impl() perm1.Registry {
	return inst
}
