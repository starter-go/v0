package tokens

import "context"

type TokenService interface {
	GetToken(c context.Context) (*TokenDTO, error)
	SetToken(c context.Context, item *TokenDTO) (*TokenDTO, error)
}

type JWTService interface {
	GetJWT(c context.Context) (JWT, error)
	SetJWT(c context.Context, value JWT) (JWT, error)

	Decode(c context.Context, jwt JWT) (*TokenDTO, error)
	Encode(c context.Context, token *TokenDTO) (JWT, error)
}
