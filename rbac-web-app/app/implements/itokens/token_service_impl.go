package itokens

import (
	"context"

	"github.com/starter-go/v0/rbac-web-app/app/classes/subjects"
	"github.com/starter-go/v0/rbac-web-app/app/classes/tokens"
	"github.com/starter-go/v0/rbac-web-app/app/web/dto"
)

type TokenServiceAPI interface {
	tokens.Service
}

type TokenServiceImpl struct {

	//starter:component

	_as func(tokens.Service) //starter:as("#")

	JWTser tokens.JWTokenService //starter:inject("#")
}

func (inst *TokenServiceImpl) _impl() TokenServiceAPI {
	return inst
}

func (inst *TokenServiceImpl) GetCurrentToken(cc context.Context) (*dto.Token, error) {

	sub, err := subjects.Current(cc)
	if err != nil {
		return nil, err
	}

	return sub.GetToken()
}

func (inst *TokenServiceImpl) LoadCurrentToken(cc context.Context) (*dto.Token, error) {

	// holder, err := webcontexts.GetHolder(cc)
	// gc := holder.Get()

	ser := inst.JWTser
	jwt, err := ser.ReadJWT(cc)
	if err != nil {
		return nil, err
	}

	return ser.DecodeJWT(jwt)
}

func (inst *TokenServiceImpl) SetCurrentToken(cc context.Context, token *dto.Token) (*dto.Token, error) {

	ser := inst.JWTser
	jwt, err := ser.EncodeJWT(token)
	if err != nil {
		return nil, err
	}

	err = ser.WriteJWT(cc, jwt)
	if err != nil {
		return nil, err
	}

	return token, nil
}
