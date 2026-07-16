package iexamples

import (
	"github.com/starter-go/security/random"
	"github.com/starter-go/v0/rbac-web-app/app/classes/examples"
	"github.com/starter-go/v0/rbac-web-app/app/data/database"
	"gorm.io/gorm"
)

type ExampleDaoImpl struct {

	//starter:component

	_as func(examples.DAO) //starter:as("#")

	Agent   database.Agent     //starter:inject("#")
	UUIDSer random.UUIDService //starter:inject("#")

}

// Find implements examples.DAO.
func (inst *ExampleDaoImpl) Find(db *gorm.DB, id examples.ID) (*examples.Entity, error) {
	panic("unimplemented")
}

func (inst *ExampleDaoImpl) _impl() examples.DAO {
	return inst
}
