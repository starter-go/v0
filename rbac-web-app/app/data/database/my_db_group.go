package database

import (
	"github.com/starter-go/libgorm"
	"github.com/starter-go/v0/rbac-web-app/app/data/entity"
)

type MyDataGroup struct {

	//starter:component

	_as func(libgorm.GroupRegistry) //starter:as(".")

	Enabled bool   //starter:inject("${datagroup.default.enabled}")
	Alias   string //starter:inject("${datagroup.default.alias}")
	Prefix  string //starter:inject("${datagroup.default.table-name-prefix}")
	Source  string //starter:inject("${datagroup.default.datasource}")
	URI     string //starter:inject("${datagroup.default.uri}")

}

func (inst *MyDataGroup) _impl() (libgorm.GroupRegistry, libgorm.Group) {
	return inst, inst
}

func (inst *MyDataGroup) Groups() []*libgorm.GroupRegistration {

	r1 := new(libgorm.GroupRegistration)
	r1.Alias = inst.Alias
	r1.Enabled = inst.Enabled
	r1.Prefix = inst.Prefix
	r1.Source = inst.Source
	r1.URI = inst.URI

	r1.Group = inst

	return []*libgorm.GroupRegistration{r1}
}

func (inst *MyDataGroup) Prototypes() []any {
	prefix := inst.Prefix
	all := entity.ListAllTables(prefix)
	return all
}
