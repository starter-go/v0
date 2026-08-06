package idatagroup

import (
	"github.com/starter-go/libgorm"
	"github.com/starter-go/rbac"
)

////////////////////////////////////////////////////////////////////////////////

type StdRbacDataGroup struct {

	//starter:component

	_as func(libgorm.GroupRegistry) //starter:as(".")

	ConfigAlias   string //starter:inject("${datagroup.std-rbac-dg.alias}")
	ConfigEnabled bool   //starter:inject("${datagroup.std-rbac-dg.enabled}")
	ConfigPrefix  string //starter:inject("${datagroup.std-rbac-dg.table-name-prefix}")
	ConfigSource  string //starter:inject("${datagroup.std-rbac-dg.datasource}")
	ConfigURI     string //starter:inject("${datagroup.std-rbac-dg.uri}")

}

// Prototypes implements [libgorm.Group].
func (inst *StdRbacDataGroup) Prototypes() []any {

	prefix := inst.ConfigPrefix

	return rbac.ListEntities(prefix)
}

// Groups implements [libgorm.GroupRegistry].
func (inst *StdRbacDataGroup) Groups() []*libgorm.GroupRegistration {

	r1 := &libgorm.GroupRegistration{

		Alias:   inst.ConfigAlias,
		Enabled: inst.ConfigEnabled,
		Prefix:  inst.ConfigPrefix,
		Source:  inst.ConfigSource,
		URI:     inst.ConfigURI,

		Group: inst,
	}

	return []*libgorm.GroupRegistration{r1}
}

func (inst *StdRbacDataGroup) _impl() (libgorm.GroupRegistry, libgorm.Group) {
	return inst, inst
}

////////////////////////////////////////////////////////////////////////////////
