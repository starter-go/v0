package isubjectdb

import (
	"fmt"

	"github.com/starter-go/libgorm"
	"github.com/starter-go/v0/subjects/core/data/subjectdb"
	"gorm.io/gorm"
)

type SubjectDataBaseAgent struct {

	//starter:component

	_as func(subjectdb.Agent) //starter:as("#")

	DSM        libgorm.DataSourceManager //starter:inject("#")
	SourceName string                    //starter:inject("${datagroup.subjects.datasource}")

	cached *gorm.DB
}

// DB implements subjectdb.Agent.
func (inst *SubjectDataBaseAgent) DB(db *gorm.DB) *gorm.DB {

	if db != nil {
		return db
	}

	db = inst.cached
	if db != nil {
		return db
	}

	db, err := inst.openDB()
	if err != nil {
		panic(err)
	}

	inst.cached = db
	return db
}

func (inst *SubjectDataBaseAgent) openDB() (*gorm.DB, error) {

	loader := new(libgorm.DataSourceAgent)
	loader.Init(inst.DSM, inst.SourceName)

	if !loader.Ready() {
		return nil, fmt.Errorf("SubjectDataBaseAgent.openDB() : data-source-agent is not ready")
	}

	var db *gorm.DB
	db = loader.DB(db)

	if db == nil {
		return nil, fmt.Errorf("SubjectDataBaseAgent.openDB() : db is nil")
	}

	return db, nil
}

func (inst *SubjectDataBaseAgent) _impl() subjectdb.Agent {
	return inst
}
