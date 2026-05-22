package icodec

import (
	"encoding/json"
	"fmt"

	jwtv5 "github.com/golang-jwt/jwt/v5"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/v0/libjwt/api/jwt"
)

////////////////////////////////////////////////////////////////////////////////

type innerSimpleDecoder struct {
	key []byte
}

// Decode implements jwt.Decoder.
func (inst *innerSimpleDecoder) Decode(txt jwt.Text) (*jwt.Token, error) {

	tokenString := txt.String()
	hmacSampleSecret := inst.key

	token, err := jwtv5.Parse(tokenString, func(token *jwtv5.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtv5.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSampleSecret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwtv5.MapClaims)
	if !ok {
		return nil, fmt.Errorf("bad jwt.claims")
	}

	if !token.Valid {
		return nil, fmt.Errorf("bad jwt.valid")
	}

	tokenJSON := fmt.Sprint(claims["token"])
	t2 := &jwt.Token{}
	err = json.Unmarshal([]byte(tokenJSON), t2)
	if err != nil {
		return nil, err
	}

	err = inst.verify(t2)
	if err != nil {
		return nil, err
	}

	return t2, nil

}

func (inst *innerSimpleDecoder) verify(t *jwt.Token) error {

	if t == nil {
		return fmt.Errorf("jwt.decoder: token is nil")
	}

	t0 := t.NotBefore
	t1 := t.NotAfter
	now := lang.Now()
	timeIsOK := (t0 <= now && now <= t1)

	if !timeIsOK {
		return fmt.Errorf("jwt.decoder: token is expired")
	}

	return nil
}

func (inst *innerSimpleDecoder) _impl() jwt.Decoder {
	return inst
}

////////////////////////////////////////////////////////////////////////////////
