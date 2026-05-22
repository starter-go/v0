package iservice

import "github.com/starter-go/v0/libjwt/api/jwt"

type JwtServiceImpl struct {

	//starter:component

	_as func(jwt.Service) //starter:as("#")

	Adapter jwt.Adapter //starter:inject("#")
	Codec   jwt.CODEC   //starter:inject("#")

}

// GetAdapter implements jwt.Service.
func (inst *JwtServiceImpl) GetAdapter() jwt.Adapter {
	return inst.Adapter
}

// GetDecoder implements jwt.Service.
func (inst *JwtServiceImpl) GetDecoder() jwt.Decoder {
	return inst.Codec
}

// GetEncoder implements jwt.Service.
func (inst *JwtServiceImpl) GetEncoder() jwt.Encoder {
	return inst.Codec
}

func (inst *JwtServiceImpl) _impl() jwt.Service {
	return inst
}
