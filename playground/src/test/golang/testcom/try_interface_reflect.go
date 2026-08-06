package testcom

import (
	"context"
	"fmt"
	"reflect"

	"github.com/starter-go/units"
	"github.com/starter-go/vlog"
)

type TryInterfaceReflect struct {

	//starter:component

	Ser1 units.Unit
}

// ListRegistrations implements units.Unit.
func (inst *TryInterfaceReflect) ListRegistrations(list []*units.Registration) []*units.Registration {

	u1 := &units.Registration{
		Name:    "try-iface-reflect",
		Enabled: true,
		Do:      inst.run,
	}

	list = append(list, u1)
	return list
}

func (inst *TryInterfaceReflect) run(cc context.Context) error {

	str1 := inst.reflect(&inst.Ser1)

	vlog.Info("typeof(units.Unit) = %s", str1)

	return nil
}

func (inst *TryInterfaceReflect) reflect(api any) string {

	tt := reflect.TypeOf(api)

	name := tt.Name()
	str := tt.String()
	ppath := tt.PkgPath()

	return fmt.Sprintf("[reflect.TypeOf name:'%s' string:'%s' pkg_path:'%s' ]", name, str, ppath)
}

func (inst *TryInterfaceReflect) _impl() units.Unit {
	return inst
}
