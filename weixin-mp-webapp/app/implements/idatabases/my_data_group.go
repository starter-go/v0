package idatabases

import (
	"github.com/starter-go/libgorm"
	"github.com/starter-go/v0/weixin-mp-webapp/app/data/entity"
)

type MyDataGroup struct {

	//starter:component

	_as func(libgorm.GroupRegistry) //starter:as(".")

	Alias   string //starter:inject("${datagroup.weixin-mp.alias}")
	URI     string //starter:inject("${datagroup.weixin-mp.uri}")
	Prefix  string //starter:inject("${datagroup.weixin-mp.table-name-prefix}")
	Source  string //starter:inject("${datagroup.weixin-mp.datasource}")
	Enabled bool   //starter:inject("${datagroup.weixin-mp.enabled}")
}

// Prototypes implements libgorm.Group.
func (inst *MyDataGroup) Prototypes() []any {
	prefix := inst.Prefix
	all := entity.ListAllTables(prefix)
	return all
}

// Groups implements libgorm.GroupRegistry.
func (inst *MyDataGroup) Groups() []*libgorm.GroupRegistration {
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

func (inst *MyDataGroup) _impl() (libgorm.Group, libgorm.GroupRegistry) {
	return inst, inst
}
