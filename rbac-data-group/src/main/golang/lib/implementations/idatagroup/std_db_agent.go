package idatagroup

import (
	"github.com/starter-go/libgorm"
	"github.com/starter-go/v0/rbac-data-group/src/main/golang/api/daos"
	"gorm.io/gorm"
)

////////////////////////////////////////////////////////////////////////////////

type StdRbacDataAgent struct {

	//starter:component

	_as func(daos.IDatabaseAgent) //starter:as("#")

	// service

	DSM libgorm.DataSourceManager //starter:inject("#")

	// config

	ConfigSource string //starter:inject("${datagroup.std-rbac-dg.datasource}")

	// cache

	cache *gorm.DB
}

func (inst *StdRbacDataAgent) load() (*gorm.DB, error) {
	source := inst.ConfigSource
	ds, err := inst.DSM.GetDataSource(source)
	if err != nil {
		return nil, err
	}
	return ds.DB()
}

// DB implements [daos.IDatabaseAgent].
func (inst *StdRbacDataAgent) DB(db *gorm.DB) *gorm.DB {
	c := inst.cache
	if c == nil {
		c2, err := inst.load()
		if err != nil {
			panic(err)
		}
		c = c2
		inst.cache = c2
	}
	return c
}

func (inst *StdRbacDataAgent) _impl() daos.IDatabaseAgent {
	return inst
}
