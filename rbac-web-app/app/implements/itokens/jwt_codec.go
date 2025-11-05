package itokens

import (
	"github.com/starter-go/v0/rbac-web-app/app/classes/tokens"

	jwt "github.com/golang-jwt/jwt/v5"
)

type innerJWTCodec struct {
	key []byte
}

func (inst *innerJWTCodec) encode(attrs map[string]string) (tokens.JWT, error) {
	// Create a new JWT token with HMAC signing method
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the claims from the attrs map
	claims := make(jwt.MapClaims)
	for k, v := range attrs {
		claims[k] = v
	}
	token.Claims = claims

	// Sign the token with the key
	tokenString, err := token.SignedString(inst.key)
	if err != nil {
		return "", err
	}

	return tokens.JWT(tokenString), nil
}

func (inst *innerJWTCodec) decode(str tokens.JWT) (attrs map[string]string, err error) {
	// Parse the token
	token, err := jwt.Parse(str.String(), func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return inst.key, nil
	})

	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, jwt.ErrTokenMalformed
	}

	// Extract claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		attrs = make(map[string]string)
		for k, v := range claims {
			if s, ok := v.(string); ok {
				attrs[k] = s
			}
		}
		return attrs, nil
	}

	return nil, jwt.ErrTokenMalformed
}
