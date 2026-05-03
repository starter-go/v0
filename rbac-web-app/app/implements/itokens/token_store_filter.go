package itokens

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/v0/rbac-web-app/app/classes/statestores"
	"github.com/starter-go/v0/rbac-web-app/app/classes/tokens"
	"github.com/starter-go/v0/rbac-web-app/app/web/dto"
)

type TokenStoreFilter struct {

	//starter:component

	_as func(statestores.FilterRegistry) //starter:as(".")

	Tokens tokens.Service //starter:inject("#")

}

// Handle implements sessions.StoreFilter.
func (inst *TokenStoreFilter) Handle(sc *statestores.Context, next statestores.FilterChain) error {

	method := sc.Method

	switch method {
	case statestores.MethodGet:
		return inst.handleMethodGet(sc, next)

	case statestores.MethodPut:
		return inst.handleMethodPut(sc, next)

	case statestores.MethodPost:
		return inst.handleMethodPost(sc, next)

	case statestores.MethodDelete:
		return inst.handleMethodDel(sc, next)
	}
	return next.Handle(sc)
}

func (inst *TokenStoreFilter) handleMethodGet(sc *statestores.Context, next statestores.FilterChain) error {

	ser := inst.Tokens
	ctx := sc.CC

	token, err := ser.GetCurrentToken(ctx)
	if err != nil {
		return err
	}

	sc.SessionID = token.SessionID
	sc.SessionUUID = token.SessionUUID

	return next.Handle(sc)
}

func (inst *TokenStoreFilter) handleMethodPost(sc *statestores.Context, next statestores.FilterChain) error {

	err := next.Handle(sc)
	if err != nil {
		return err
	}

	ser := inst.Tokens
	ctx := sc.CC
	token := new(dto.Token)

	now := lang.Now()
	ttl := sc.TokenMaxAge.Milliseconds()

	token.SessionID = sc.SessionID
	token.SessionUUID = sc.SessionUUID
	token.CreatedAt = now
	token.AliveFrom = now
	token.AliveTo = now + lang.Time(ttl)

	_, err = ser.SetCurrentToken(ctx, token)
	return err
}

func (inst *TokenStoreFilter) handleMethodPut(sc *statestores.Context, next statestores.FilterChain) error {

	err := next.Handle(sc)
	if err != nil {
		return err
	}

	ser := inst.Tokens
	ctx := sc.CC
	token := new(dto.Token)

	now := lang.Now()
	ttl := sc.TokenMaxAge.Milliseconds()

	token.SessionID = sc.SessionID
	token.SessionUUID = sc.SessionUUID
	token.CreatedAt = now
	token.AliveFrom = now
	token.AliveTo = now + lang.Time(ttl)

	_, err = ser.SetCurrentToken(ctx, token)
	return err
}

func (inst *TokenStoreFilter) handleMethodDel(sc *statestores.Context, next statestores.FilterChain) error {

	err := next.Handle(sc)
	if err != nil {
		return err
	}

	ser := inst.Tokens
	ctx := sc.CC
	token := new(dto.Token)

	token.SessionID = sc.SessionID
	token.SessionUUID = sc.SessionUUID

	_, err = ser.SetCurrentToken(ctx, token)
	return err
}

// ListRegistrations implements sessions.StoreFilterRegistry.
func (inst *TokenStoreFilter) ListRegistrations() []*statestores.FilterRegistration {

	reg := &statestores.FilterRegistration{
		Name:     "TokenStoreFilter",
		Enabled:  true,
		Priority: statestores.FilterPriorityToken,
		Filter:   inst,
	}

	return []*statestores.FilterRegistration{reg}
}

func (inst *TokenStoreFilter) _impl() (statestores.FilterRegistry, statestores.Filter) {
	return inst, inst
}
