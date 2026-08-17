package testcom

import (
	"context"
	"strconv"

	"github.com/starter-go/application/properties"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/api/classes/permissions"
	"github.com/starter-go/rbac/api/classes/roles"
	"github.com/starter-go/rbac/api/classes/sessions"
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

	const bar = "--------------------------------------------------------------"
	const f = "[user_entity index:%d id:%d uuid:'%s' username:'%s' ]"

	for index, it := range all {
		if it == nil {
			vlog.Info(bar)
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

	dao := inst.PermDao
	db := dao.GetDB(nil)
	now := lang.Now()

	// insert

	time_str := "t" + strconv.FormatInt(now.Int(), 10)
	item1 := new(rbac.PermissionEntity)

	item1.Method = "GET"
	item1.Path = "/demo/a/b/c"
	item1.Roles = "demo,user"
	item1.URI = lang.URI("uri://roles/" + time_str)

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

	item4, err := dao.Update(db, id, func(old *permissions.Entity) error {
		old.Priority = 123
		return nil
	})

	if err != nil {
		return err
	}

	// list

	q := new(permissions.Query)
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

	list2 := []*permissions.Entity{nil, item1, item2, item3, item4}
	all := list1

	all = append(all, nil)
	all = append(all, list2...)

	const bar = "--------------------------------------------------------------"
	const f = "[perm_entity index:%d id:%d uuid:'%s' method:'%s' path:'%s' roles:'%s' uri:'%s' ]"

	for index, it := range all {
		if it == nil {
			vlog.Info(bar)
			continue
		}
		id := it.ID
		uuid := it.UUID
		method := it.Method
		path := it.Path
		roles := it.Roles
		uri := it.URI
		vlog.Info(f, index, id, uuid, method, path, roles, uri)
	}

	return nil
}

func (inst *TestRbacDgCrudUnit) runTryRoleDao(cc context.Context) error {

	dao := inst.RoleDao
	db := dao.GetDB(nil)
	now := lang.Now()

	// insert

	time_str := "t" + strconv.FormatInt(now.Int(), 10)
	item1 := new(rbac.RoleEntity)
	item1.Name = roles.Name("demo" + time_str)

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

	item4, err := dao.Update(db, id, func(old *roles.Entity) error {
		cl := lang.ClassOf(old)
		old.Description = "" + cl.FullName()
		return nil
	})

	if err != nil {
		return err
	}

	// list

	q := new(roles.Query)
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

	list2 := []*roles.Entity{nil, item1, item2, item3, item4}
	all := list1

	all = append(all, nil)
	all = append(all, list2...)

	const bar = "--------------------------------------------------------------"
	const f = "[role_entity index:%d id:%d uuid:'%s' name:'%s' desc:'%s' ]"

	for index, it := range all {
		if it == nil {
			vlog.Info(bar)
			continue
		}
		id := it.ID
		uuid := it.UUID
		name := it.Name
		desc := it.Description
		vlog.Info(f, index, id, uuid, name, desc)
	}

	return nil
}

func (inst *TestRbacDgCrudUnit) runTrySessionDao(cc context.Context) error {

	dao := inst.SessionDao
	db := dao.GetDB(nil)
	now := lang.Now()

	// insert

	pt := properties.NewTable(nil)
	pt.SetProperty("time", now.String())
	ptm := properties.Map(pt.Export(nil))

	item1 := new(rbac.SessionEntity)
	item1.Authenticated = false
	item1.Properties = ptm.Text()

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

	item4, err := dao.Update(db, id, func(old *sessions.Entity) error {
		old.Authenticated = true
		return nil
	})

	if err != nil {
		return err
	}

	// list

	q := new(sessions.Query)
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

	list2 := []*sessions.Entity{nil, item1, item2, item3, item4}
	all := list1

	all = append(all, nil)
	all = append(all, list2...)

	const bar = "--------------------------------------------------------------"
	const f = "[session_entity index:%d id:%d uuid:'%s' auth:%v props:\n %s \n]"

	for index, it := range all {
		if it == nil {
			vlog.Info(bar)
			continue
		}
		id := it.ID
		uuid := it.UUID
		auth := it.Authenticated
		ptxt := it.Properties
		vlog.Info(f, index, id, uuid, auth, ptxt)
	}

	return nil
}

func (inst *TestRbacDgCrudUnit) _impl() units.Unit {
	return inst
}
