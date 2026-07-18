package itables

import (
	"context"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/libgorm"
	"github.com/starter-go/v0/rbac-web-app/app/classes/tables"
)

type TableServiceImpl struct {

	//starter:component

	_as func(tables.Service) //starter:as("#")

	Dao tables.DAO //starter:inject("#")

	DataGroupList []libgorm.GroupRegistry //starter:inject(".")

}

// Find implements [tables.Service].
func (inst *TableServiceImpl) Find(cc context.Context, id tables.ID) (*tables.DTO, error) {

	it1, err := inst.Dao.Find(nil, id)
	if err != nil {
		return nil, err
	}

	it2 := new(tables.DTO)
	err = tables.ConvertE2D(it1, it2)
	return it2, err
}

// Insert implements [tables.Service].
func (inst *TableServiceImpl) Insert(cc context.Context, item *tables.DTO) (*tables.DTO, error) {

	it2e := new(tables.Entity)
	err := tables.ConvertD2E(item, it2e)
	if err != nil {
		return nil, err
	}

	it2e.TableURI = inst.innerComputeTableURIWithEntity(it2e)

	it3e, err := inst.Dao.Insert(nil, it2e)
	if err != nil {
		return nil, err
	}

	it4d := new(tables.DTO)
	err = tables.ConvertE2D(it3e, it4d)
	return it4d, err
}

// Query implements [tables.Service].
func (inst *TableServiceImpl) Query(cc context.Context, q *tables.Query) ([]*tables.DTO, error) {
	panic("unimplemented")
}

// ScanCurrentRuntime implements [tables.Service].
func (inst *TableServiceImpl) ScanCurrentRuntime(cc context.Context) ([]*tables.DTO, error) {

	scanner := new(innerTableGroupScanner)
	scanner.init(cc, inst)

	err := scanner.scan(inst.DataGroupList)
	if err != nil {
		return nil, err
	}

	res := scanner.getResultList()
	return res, nil
}

// Setup implements [tables.Service].
func (inst *TableServiceImpl) Setup(cc context.Context) ([]*tables.DTO, error) {
	panic("unimplemented")
}

func (inst *TableServiceImpl) innerComputeTableURIWithEntity(o *tables.Entity) lang.URI {
	s1 := o.GroupURI.String()
	s2 := string(o.Name)
	str := s1 + "/" + s2
	return lang.URI(str)
}

func (inst *TableServiceImpl) _impl() tables.Service {
	return inst
}
