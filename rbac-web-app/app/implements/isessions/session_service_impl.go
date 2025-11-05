package isessions

import (
	"context"
	"fmt"

	"github.com/starter-go/v0/rbac-web-app/app/classes/sessions"
	"github.com/starter-go/v0/rbac-web-app/app/data/dxo"
	"github.com/starter-go/v0/rbac-web-app/app/data/entity"
	"github.com/starter-go/v0/rbac-web-app/app/web/dto"
)

type SessionServiceAPI interface {
	sessions.Service
}

type SessionServiceImpl struct {

	//starter:component

	_as func(sessions.Service) //starter:as("#")

	Dao sessions.DAO //starter:inject("#")

}

func (inst *SessionServiceImpl) _impl() SessionServiceAPI {
	return inst
}

func (inst *SessionServiceImpl) Find(cc context.Context, id dxo.SessionID) (*dto.Session, error) {

	return nil, fmt.Errorf("no impl ")

}

func (inst *SessionServiceImpl) Insert(cc context.Context, item *dto.Session) (*dto.Session, error) {

	it2 := new(entity.Session)
	it4 := new(dto.Session)
	sessions.ConvertD2E(item, it2)

	it3, err := inst.Dao.Insert(nil, it2)
	if err != nil {
		return nil, err
	}

	sessions.ConvertE2D(it3, it4)
	return it4, nil
}
