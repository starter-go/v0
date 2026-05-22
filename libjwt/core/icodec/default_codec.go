package icodec

import "github.com/starter-go/v0/libjwt/api/jwt"

type DefaultCodec struct {

	//starter:component

	_as func(jwt.CodecProvider) //starter:as(".")

	Name     string //starter:inject("${jwt-codec.default.name}")
	Enabled  bool   //starter:inject("${jwt-codec.default.enabled}")
	Priority int    //starter:inject("${jwt-codec.default.priority}")

	KeyLoader jwt.KeyLoader //starter:inject("#")

	cachedKey *jwt.Key
}

// Decode implements jwt.Decoder.
func (inst *DefaultCodec) Decode(txt jwt.Text) (*jwt.Token, error) {

	key, err := inst.innerGetKey()
	if err != nil {
		return nil, err
	}

	dec := new(innerSimpleDecoder)
	dec.key = key.Bytes()

	return dec.Decode(txt)
}

// Encode implements jwt.Encoder.
func (inst *DefaultCodec) Encode(token *jwt.Token) (jwt.Text, error) {

	key, err := inst.innerGetKey()
	if err != nil {
		return "", err
	}

	encoder := new(innerSimpleEncoder)
	encoder.key = key.Bytes()

	return encoder.Encode(token)
}

func (inst *DefaultCodec) innerGetKey() (*jwt.Key, error) {

	k := inst.cachedKey
	if k != nil {
		return k, nil
	}

	ldr := inst.KeyLoader
	k2, err := ldr.Load()
	if err != nil {
		return nil, err
	}

	inst.cachedKey = k2
	k = k2

	return k, nil
}

// Load implements jwt.CodecLoader.
func (inst *DefaultCodec) Load(target *jwt.CodecRegistration) error {
	target.Codec = inst
	return nil
}

// Registration implements jwt.CodecProvider.
func (inst *DefaultCodec) Registration() *jwt.CodecRegistration {

	r1 := &jwt.CodecRegistration{

		Name:     inst.Name,
		Enabled:  inst.Enabled,
		Priority: inst.Priority,

		Loader: inst,
		Codec:  inst,
	}

	return r1
}

func (inst *DefaultCodec) _impl() (jwt.CodecProvider, jwt.CodecLoader, jwt.Encoder, jwt.Decoder) {
	return inst, inst, inst, inst
}
