package isubjectdb

import (
	"github.com/starter-go/libgorm"
	"github.com/starter-go/v0/subjects/core/data/entities"
)

type SubjectDataGroup struct {

	//starter:component

	_as func(libgorm.GroupRegistry) //starter:as(".")

	Alias   string //starter:inject("${datagroup.subjects.alias}")
	Enabled bool   //starter:inject("${datagroup.subjects.enabled}")
	Prefix  string //starter:inject("${datagroup.subjects.table-name-prefix}")
	Source  string //starter:inject("${datagroup.subjects.datasource}")
	URI     string //starter:inject("${datagroup.subjects.uri}")

}

// Prototypes implements libgorm.Group.
func (inst *SubjectDataGroup) Prototypes() []any {

	prefix := inst.Prefix
	all := entities.ListAll(prefix)

	return all
}

// Groups implements libgorm.GroupRegistry.
func (inst *SubjectDataGroup) Groups() []*libgorm.GroupRegistration {

	r1 := &libgorm.GroupRegistration{

		Alias:   inst.Alias,
		Enabled: inst.Enabled,
		Prefix:  inst.Prefix,
		Source:  inst.Source,
		URI:     inst.URI,

		Group: inst,
	}

	return []*libgorm.GroupRegistration{r1}
}

func (inst *SubjectDataGroup) _impl() (libgorm.GroupRegistry, libgorm.Group) {
	return inst, inst
}
