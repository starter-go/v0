package testcom

import (
	"context"
	"fmt"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/v0/libdao"

	"github.com/starter-go/units"
	"github.com/starter-go/vlog"
)

type UnitForDaoProxy struct {

	//starter:component

	Dao MockDao //starter:inject("#")

}

// ListRegistrations implements units.Unit.
func (inst *UnitForDaoProxy) ListRegistrations(list []*units.Registration) []*units.Registration {

	u1 := &units.Registration{
		Name:    "unit-for-dao-proxy",
		Enabled: true,
		Do:      inst.run,
	}

	list = append(list, u1)

	return list
}

func (inst *UnitForDaoProxy) run(cc context.Context) error {

	dao := inst.Dao
	ids := []int{1, 3, 5, 7, 11, 13}

	for index, id := range ids {
		txt, err := dao.Find(id)
		if err != nil {
			return err
		}
		const f = "[test_mock_dao_find index:%d id:%d text:'%s']"
		vlog.Warn(f, index, id, txt)
	}

	return nil
}

func (inst *UnitForDaoProxy) _impl() units.Unit {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

type MockDao interface {
	libdao.DAO

	Find(id int) (string, error)
}

////////////////////////////////////////////////////////////////////////////////

type MockDaoMain struct {

	//starter:component

	_as func(MockDao) //starter:as("#")

	Selector string //starter:inject("${unit.mock-dao.selector}")

	DaoList []MockDao //starter:inject(".")

	holder libdao.DaoHolder[MockDao]
}

// Find implements [MockDao].
func (inst *MockDaoMain) Find(id int) (string, error) {

	sel := inst.Selector
	all := inst.DaoList
	sub := inst.holder.Select(sel, all)

	return sub.Find(id)
}

// GetRegistration implements [MockDao].
func (inst *MockDaoMain) GetRegistration() *libdao.DaoRegistration {
	return new(libdao.DaoRegistration)
}

func (inst *MockDaoMain) _impl() MockDao {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

type MockDao1 struct {

	//starter:component

	_as func(MockDao) //starter:as(".")

}

// Find implements [MockDao].
func (inst *MockDao1) Find(id int) (string, error) {

	// vlog.Warn("no impl: MockDao1.Find() ")

	const f = "%s?id=%d"
	cl := lang.ClassOf(inst)
	txt := fmt.Sprintf(f, cl.SimpleName(), id)
	return txt, nil

}

// GetRegistration implements [MockDao].
func (inst *MockDao1) GetRegistration() *libdao.DaoRegistration {

	dr1 := &libdao.DaoRegistration{

		ID:    "mock-dao-1",
		Name:  "MD1",
		Class: "mock-dao",

		Enabled:  true,
		Priority: 32,

		DAO: inst,
	}

	return dr1
}

func (inst *MockDao1) _impl() MockDao {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

type MockDao2 struct {

	//starter:component

	_as func(MockDao) //starter:as(".")

}

// Find implements [MockDao].
func (inst *MockDao2) Find(id int) (string, error) {

	// vlog.Warn("no impl: MockDao2.Find() ")

	const f = "%s?id=%d"
	cl := lang.ClassOf(inst)
	txt := fmt.Sprintf(f, cl.SimpleName(), id)
	return txt, nil

}

// GetRegistration implements [MockDao].
func (inst *MockDao2) GetRegistration() *libdao.DaoRegistration {

	dr1 := &libdao.DaoRegistration{

		ID:    "mock-dao-2",
		Name:  "MD2",
		Class: "mock-dao",

		Enabled:  true,
		Priority: 33,

		DAO: inst,
	}

	return dr1
}

func (inst *MockDao2) _impl() MockDao {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

type MockDao3 struct {

	//starter:component

	_as func(MockDao) //starter:as(".")

}

// Find implements [MockDao].
func (inst *MockDao3) Find(id int) (string, error) {

	// vlog.Warn("no impl: MockDao3.Find() ")

	const f = "%s?id=%d"
	cl := lang.ClassOf(inst)
	txt := fmt.Sprintf(f, cl.SimpleName(), id)
	return txt, nil

}

// GetRegistration implements [MockDao].
func (inst *MockDao3) GetRegistration() *libdao.DaoRegistration {

	dr1 := &libdao.DaoRegistration{

		ID:    "mock-dao-3",
		Name:  "MD3",
		Class: "mock-dao",

		Enabled:  true,
		Priority: 31,

		DAO: inst,
	}

	return dr1
}

func (inst *MockDao3) _impl() MockDao {
	return inst
}

////////////////////////////////////////////////////////////////////////////////
// EOF
