package idatabase

import (
	"github.com/starter-go/libgorm"
	"github.com/starter-go/v0/rbac-web-app/app/data/database"
	"gorm.io/gorm"
)

type MyDatabaseAgentAPI interface {
	database.Agent
}

type MyDatabaseAgentImpl struct {

	//starter:component

	_as func(database.Agent) //starter:as("#")

	DSManager libgorm.DataSourceManager //starter:inject("#")

	cached libgorm.Agent
}

func (inst *MyDatabaseAgentImpl) _impl() MyDatabaseAgentAPI {
	return inst
}

func (inst *MyDatabaseAgentImpl) getInner() libgorm.Agent {
	in := inst.cached
	if in == nil {
		in = inst.loadInner()
		inst.cached = in
	}
	return in
}

func (inst *MyDatabaseAgentImpl) loadInner() libgorm.Agent {
	const dsName = "main"
	man := inst.DSManager
	dsa := new(libgorm.DataSourceAgent)
	dsa.Init(man, dsName)
	return dsa
}

func (inst *MyDatabaseAgentImpl) DB(db *gorm.DB) *gorm.DB {
	return inst.getInner().DB(db)
}
