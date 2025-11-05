package tokens

import (
	"context"

	"github.com/starter-go/v0/rbac-web-app/app/web/dto"
)

type JWT string

func (t JWT) String() string {
	return string(t)
}

////////////////////////////////////////////////////////////////////////////////

type JWTokenCodec interface {
	EncodeJWT(token *dto.Token) (JWT, error)

	DecodeJWT(jwt JWT) (*dto.Token, error)
}

type JWTokenIO interface {
	ReadJWT(cc context.Context) (JWT, error)

	WriteJWT(cc context.Context, jwt JWT) error
}

type JWTokenService interface {
	JWTokenIO

	JWTokenCodec
}
