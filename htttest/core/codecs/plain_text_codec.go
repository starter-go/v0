package codecs

import (
	"github.com/starter-go/v0/htttest"
	"github.com/starter-go/v0/htttest/api/h3t"
)

type PlainTextCodec struct {

	//starter:component

	_as func(htttest.CodecProvider) //starter:as(".")

}

// Registrations implements h3t.CodecProvider.
func (inst *PlainTextCodec) Registrations() []*h3t.CodecRegistration {
	r1 := &h3t.CodecRegistration{

		Name:     "plain-text.codec",
		Type:     "text/plain",
		Enabled:  true,
		Priority: 0,

		Encoder: inst,
		Decoder: inst,
		Codec:   inst,
	}
	return []*h3t.CodecRegistration{r1}
}

// Decode implements h3t.Codec.
func (inst *PlainTextCodec) Decode(c *h3t.Content) error {
	str := ""
	src := c.Encoded
	if src != nil {
		str = string(src)
	}
	c.Data = str
	return nil
}

// Encode implements h3t.Codec.
func (inst *PlainTextCodec) Encode(c *h3t.Content) error {
	src := c.Data
	str, ok := src.(string)
	if ok {
		c.Encoded = []byte(str)
	} else {
		c.Encoded = []byte{}
	}
	return nil
}

func (inst *PlainTextCodec) _impl() (htttest.CodecProvider, htttest.Codec) {
	return inst, inst
}
