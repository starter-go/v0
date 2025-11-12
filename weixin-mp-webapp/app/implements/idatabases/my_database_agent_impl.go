package idatabases

import (
	"fmt"

	"github.com/starter-go/libgorm"
	"github.com/starter-go/v0/weixin-mp-webapp/app/data/database"
	"gorm.io/gorm"
)

type MyDatabaseAgentImpl struct {

	//starter:component

	_as func(database.Agent) //starter:as("#")

	DSM    libgorm.DataSourceManager //starter:inject("#")
	DSName string                    //starter:inject("${datagroup.weixin-mp.datasource}")

	inner libgorm.Agent
}

// DB implements database.Agent.
func (inst *MyDatabaseAgentImpl) DB(db *gorm.DB) *gorm.DB {
	if db != nil {
		return db
	}
	a2, err := inst.innerGetTarget()
	if err != nil {
		panic(err)
	}
	return a2.DB(db)
}

func (inst *MyDatabaseAgentImpl) innerGetTarget() (libgorm.Agent, error) {
	agent := inst.inner
	if agent == nil {
		a2, err := inst.innerLoadTarget()
		if err != nil {
			return nil, err
		}
		agent = a2
		inst.inner = a2
	}
	return agent, nil
}

func (inst *MyDatabaseAgentImpl) innerLoadTarget() (libgorm.Agent, error) {

	name := inst.DSName
	dsm := inst.DSM

	a2 := new(libgorm.DataSourceAgent)
	a2.Init(dsm, name)
	db := a2.DB(nil)
	if db == nil {
		return nil, fmt.Errorf("db(*gorm.DB) is nil")
	}
	return a2, nil
}

func (inst *MyDatabaseAgentImpl) _impl() database.Agent {
	return inst
}
