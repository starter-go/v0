package isessions

import (
	"context"

	"github.com/starter-go/application/properties"
	"github.com/starter-go/v0/subjects/core/classes/sessions"
)

type SessionServiceImpl struct {

	//starter:component

	_as func(sessions.Service) //starter:as("#")

	Dao sessions.DAO //starter:inject("#")

}

// Find implements sessions.Service.
func (inst *SessionServiceImpl) Find(c context.Context, id sessions.ID) (*sessions.DTO, error) {

	o1, err := inst.Dao.Find(nil, id)
	if err != nil {
		return nil, err
	}

	o2 := new(sessions.DTO)
	err = sessions.ConvertE2D(o1, o2)
	if err != nil {
		return nil, err
	}

	return o2, nil
}

// Insert implements sessions.Service.
func (inst *SessionServiceImpl) Insert(c context.Context, item *sessions.DTO) (*sessions.DTO, error) {

	o2 := new(sessions.Entity)
	err := sessions.ConvertD2E(item, o2)
	if err != nil {
		return nil, err
	}

	o3, err := inst.Dao.Insert(nil, o2)
	if err != nil {
		return nil, err
	}

	o4 := new(sessions.DTO)
	err = sessions.ConvertE2D(o3, o4)
	if err != nil {
		return nil, err
	}

	return o4, nil
}

// Query implements sessions.Service.
func (inst *SessionServiceImpl) Query(c context.Context, q *sessions.Query) ([]*sessions.DTO, error) {

	list1, err := inst.Dao.Query(nil, q)
	if err != nil {
		return nil, err
	}
	return sessions.ConvertListE2D(list1, nil)
}

// Remove implements sessions.Service.
func (inst *SessionServiceImpl) Remove(c context.Context, id sessions.ID) error {

	return inst.Dao.Remove(nil, id)
}

// Update implements sessions.Service.
func (inst *SessionServiceImpl) Update(c context.Context, id sessions.ID, o1 *sessions.DTO) (*sessions.DTO, error) {

	db := inst.Dao.PrepareDB(nil)
	o2 := new(sessions.Entity)
	err := sessions.ConvertD2E(o1, o2)
	if err != nil {
		return nil, err
	}

	o4, err := inst.Dao.Update(db, id, func(o3 *sessions.Entity) error {
		return inst.innerUpdateCallback(o2, o3)
	})

	if err != nil {
		return nil, err
	}

	o5 := new(sessions.DTO)
	err = sessions.ConvertE2D(o4, o5)
	if err != nil {
		return nil, err
	}

	return o5, nil
}

func (inst *SessionServiceImpl) innerUpdateCallback(src, dst *sessions.Entity) error {

	if !src.NotAfter.IsZero() {
		dst.NotAfter = src.NotAfter
	}

	if !src.NotBefore.IsZero() {
		dst.NotBefore = src.NotBefore
	}

	if src.Properties != "" {
		pDst := dst.Properties.Map()
		pSrc := src.Properties.Map()
		pDst = inst.innerUpdateProps(pSrc, pDst)
		dst.Properties = pDst.Text()
	}

	return nil
}

func (inst *SessionServiceImpl) innerUpdateProps(src, dst properties.Map) properties.Map {
	for key, value := range src {
		dst.Put(key, value)
	}
	return dst.Trim()
}

func (inst *SessionServiceImpl) _impl() sessions.Service {
	return inst
}
