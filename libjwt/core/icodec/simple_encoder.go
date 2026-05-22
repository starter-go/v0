package icodec

import (
	"encoding/json"

	jwtv5 "github.com/golang-jwt/jwt/v5"

	"github.com/starter-go/v0/libjwt/api/jwt"
)

////////////////////////////////////////////////////////////////////////////////

type innerSimpleEncoder struct {
	key []byte
}

// Encode implements jwt.Encoder.
func (inst *innerSimpleEncoder) Encode(token *jwt.Token) (jwt.Text, error) {

	key := inst.key
	t := jwtv5.New(jwtv5.SigningMethodHS256)
	t1 := token

	if t1 == nil {
		t1 = &jwt.Token{}
	}

	tokenJSON, err := json.Marshal(t1)
	if err != nil {
		return "", err
	}

	claims := make(jwtv5.MapClaims)
	claims["token"] = string(tokenJSON)
	t.Claims = claims

	s, err := t.SignedString(key)
	if err != nil {
		return "", err
	}

	return jwt.Text(s), err

}

func (inst *innerSimpleEncoder) _impl() jwt.Encoder {
	return inst
}

////////////////////////////////////////////////////////////////////////////////
