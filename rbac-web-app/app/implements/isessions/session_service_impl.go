package isessions

import (
	"context"
	"fmt"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/v0/rbac-web-app/app/classes/sessions"
	"github.com/starter-go/v0/rbac-web-app/app/classes/statestores"
	"github.com/starter-go/v0/rbac-web-app/app/classes/tokens"
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

	TokenService tokens.Service //starter:inject("#")

	StoreChainHolder statestores.FilterChainHolder //starter:inject("#")

}

// FindByUUID implements SessionServiceAPI.
func (inst *SessionServiceImpl) FindByUUID(cc context.Context, uuid lang.UUID) (*dto.Session, error) {

	o1, err := inst.Dao.FindByUUID(nil, uuid)
	if err != nil {
		return nil, err
	}

	o2 := new(dto.Session)
	err = sessions.ConvertE2D(o1, o2)
	return o2, err
}

// Remove implements SessionServiceAPI.
func (inst *SessionServiceImpl) Remove(cc context.Context, id dxo.SessionID) error {

	return inst.Dao.Remove(nil, id)
}

// Update implements SessionServiceAPI.
func (inst *SessionServiceImpl) Update(cc context.Context, id dxo.SessionID, item *dto.Session) (*dto.Session, error) {

	src := new(entity.Session)
	err := sessions.ConvertD2E(item, src)
	if err != nil {
		return nil, err
	}

	o3, err := inst.Dao.Update(nil, id, func(dst *entity.Session) error {

		dst.DisplayName = src.DisplayName
		dst.Roles = src.Roles

		return nil
	})

	o4 := new(dto.Session)
	err = sessions.ConvertE2D(o3, o4)
	return o4, err
}

// GetStore implements SessionServiceAPI.
// func (inst *SessionServiceImpl) GetStore(cc context.Context) (statestores.Store, error) {
// 	ch, err := inst.StoreChainHolder.GetChain()
// 	return ch, err
// }

// GetCurrent implements SessionServiceAPI.
func (inst *SessionServiceImpl) GetCurrent(cc context.Context) (*dto.Session, error) {

	token, err := inst.TokenService.GetCurrentToken(cc)
	if err != nil {
		return nil, err
	}

	session, err := inst.Find(cc, token.SessionID)
	if err != nil {
		return nil, err
	}

	now := lang.Now()

	if session.UUID != token.SessionUUID {
		return nil, fmt.Errorf("bad token")
	}
	if !inst.isSessionAlive(session, now) {
		return nil, fmt.Errorf("bad session")
	}
	if !inst.isTokenAlive(token, now) {
		return nil, fmt.Errorf("bad token")
	}

	return session, nil
}

func (inst *SessionServiceImpl) isTokenAlive(item *dto.Token, now lang.Time) bool {

	if item == nil {
		return false
	}

	// if !item.Alive { 		return false 	}

	return ((item.AliveFrom <= now) && (now <= item.AliveTo))
}

func (inst *SessionServiceImpl) isSessionAlive(item *dto.Session, now lang.Time) bool {

	if item == nil {
		return false
	}

	if !item.Alive {
		return false
	}

	return ((item.AliveFrom <= now) && (now <= item.AliveTo))
}

func (inst *SessionServiceImpl) _impl() SessionServiceAPI {
	return inst
}

func (inst *SessionServiceImpl) Find(cc context.Context, id dxo.SessionID) (*dto.Session, error) {

	o1, err := inst.Dao.Find(nil, id)
	if err != nil {
		return nil, err
	}

	o2 := new(dto.Session)
	err = sessions.ConvertE2D(o1, o2)
	return o2, err
}

func (inst *SessionServiceImpl) Insert(cc context.Context, item *dto.Session) (*dto.Session, error) {

	it2 := new(entity.Session)
	err := sessions.ConvertD2E(item, it2)
	if err != nil {
		return nil, err
	}

	it3, err := inst.Dao.Insert(nil, it2)
	if err != nil {
		return nil, err
	}

	it4 := new(dto.Session)
	err = sessions.ConvertE2D(it3, it4)
	return it4, err
}
