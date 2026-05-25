package codecs

import (
	"github.com/starter-go/v0/htttest"
	"github.com/starter-go/v0/htttest/api/h3t"
)

type BinaryCodec struct {

	//starter:component

	_as func(htttest.CodecProvider) //starter:as(".")

}

// Registrations implements h3t.CodecProvider.
func (inst *BinaryCodec) Registrations() []*h3t.CodecRegistration {
	r1 := &h3t.CodecRegistration{

		Name:     "binary.codec",
		Type:     "application/octet-stream",
		Enabled:  true,
		Priority: 0,

		Encoder: inst,
		Decoder: inst,
		Codec:   inst,
	}
	return []*h3t.CodecRegistration{r1}
}

// Decode implements h3t.Codec.
func (inst *BinaryCodec) Decode(c *h3t.Content) error {
	panic("unimplemented")
}

// Encode implements h3t.Codec.
func (inst *BinaryCodec) Encode(c *h3t.Content) error {
	panic("unimplemented")
}

func (inst *BinaryCodec) _impl() (htttest.CodecProvider, htttest.Codec) {
	return inst, inst
}
