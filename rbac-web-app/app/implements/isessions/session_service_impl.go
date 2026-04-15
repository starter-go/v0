package isessions

import (
	"context"
	"fmt"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/v0/rbac-web-app/app/classes/sessions"
	"github.com/starter-go/v0/rbac-web-app/app/classes/subjects"
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

}

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

	sub, err := subjects.Current(cc)
	if err != nil {
		return nil, err
	}
	return sub.GetSession()
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
