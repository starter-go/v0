package testcom

import (
	"context"

	"github.com/starter-go/units"
	"github.com/starter-go/v0/keys"
	"github.com/starter-go/vlog"
)

type DriversUnit struct {

	//starter:component

	Ser keys.Service //starter:inject("#")
}

// ListRegistrations implements [units.Unit].
func (inst *DriversUnit) ListRegistrations(list []*units.Registration) []*units.Registration {

	ur1 := &units.Registration{
		ID:       "test-drivers",
		Enabled:  true,
		Priority: 0,
		Do:       inst.run,
	}

	list = append(list, ur1)
	return list
}

func (inst *DriversUnit) run(c context.Context) error {

	md := inst.Ser.GetDriverManager()
	list := md.ListDrivers()

	vlog.Info("list of keys (algorithm) drviers")

	for idx, drv := range list {
		inst.innerLogDriver(idx, drv)
	}

	return nil
}

func (inst *DriversUnit) innerLogDriver(index int, driver keys.Driver) {

	reg := driver.GetRegistration()

	algorithm := reg.Algorithm
	name := reg.Name
	priority := reg.Priority
	en := reg.Enabled

	const f = "[driver index:%d name:'%s' algorithm:'%s' priority:%d enabled:%v ]"
	vlog.Info(f, index, name, algorithm, priority, en)
}

func (inst *DriversUnit) _impl() units.Unit {
	return inst
}
