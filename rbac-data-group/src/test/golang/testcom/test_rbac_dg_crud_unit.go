package testcom

import (
	"context"
	"strconv"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/api/classes/users"
	"github.com/starter-go/units"
	"github.com/starter-go/vlog"
)

type TestRbacDgCrudUnit struct {

	//starter:component

	PermDao    rbac.PermissionDAO //starter:inject("#")
	RoleDao    rbac.RoleDAO       //starter:inject("#")
	SessionDao rbac.SessionDAO    //starter:inject("#")
	UserDao    rbac.UserDAO       //starter:inject("#")

}

// ListRegistrations implements units.Unit.
func (inst *TestRbacDgCrudUnit) ListRegistrations(list []*units.Registration) []*units.Registration {

	u1 := &units.Registration{
		Name:    "test-rbacdg-permission-dao",
		Enabled: true,
		Do:      inst.runTryPermDao,
	}

	u2 := &units.Registration{
		Name:    "test-rbacdg-role-dao",
		Enabled: true,
		Do:      inst.runTryRoleDao,
	}

	u3 := &units.Registration{
		Name:    "test-rbacdg-session-dao",
		Enabled: true,
		Do:      inst.runTrySessionDao,
	}

	u4 := &units.Registration{
		Name:    "test-rbacdg-user-dao",
		Enabled: true,
		Do:      inst.runTryUserDao,
	}

	list = append(list, u1, u2, u3, u4)
	return list
}

func (inst *TestRbacDgCrudUnit) run(cc context.Context) error {
	return nil
}

func (inst *TestRbacDgCrudUnit) runTryUserDao(cc context.Context) error {

	dao := inst.UserDao
	db := dao.GetDB(nil)
	now := lang.Now()

	// insert

	user_name := "user-" + strconv.FormatInt(now.Int(), 10)

	item1 := new(rbac.UserEntity)
	item1.Avatar = "http://example.com/avatar.jpg"
	item1.Name = users.Name(user_name)
	item1.Phone = users.PhoneNumber("phone:" + user_name)
	item1.Email = users.EmailAddress("mailto:" + user_name + "@mock.com")

	item2, err := dao.Insert(db, item1)
	if err != nil {
		return err
	}

	// find

	id := item2.ID

	item3, err := dao.Find(db, id)
	if err != nil {
		return err
	}

	// update

	item4, err := dao.Update(db, id, func(old *users.Entity) error {

		old.Domain = "foo.bar.com"
		return nil
	})

	if err != nil {
		return err
	}

	// list

	q := new(users.Query)
	q.Pagination.Limit = 5
	q.Pagination.Offset = 0

	list1, err := dao.Query(db, q)
	if err != nil {
		return err
	}

	// delete
	err = dao.Delete(db, id)
	if err != nil {
		return err
	}

	// log all items:

	list2 := []*users.Entity{nil, item1, item2, item3, item4}
	all := list1

	all = append(all, nil)
	all = append(all, list2...)

	const f = "[user_entity index:%d id:%d uuid:'%s' username:'%s' ]"

	for index, it := range all {
		if it == nil {
			vlog.Info("-------------------")
			continue
		}
		id := it.ID
		uuid := it.UUID
		username := it.Name
		vlog.Info(f, index, id, uuid, username)
	}

	return nil
}

func (inst *TestRbacDgCrudUnit) runTryPermDao(cc context.Context) error {
	return nil
}

func (inst *TestRbacDgCrudUnit) runTryRoleDao(cc context.Context) error {
	return nil
}

func (inst *TestRbacDgCrudUnit) runTrySessionDao(cc context.Context) error {
	return nil
}

func (inst *TestRbacDgCrudUnit) _impl() units.Unit {
	return inst
}
