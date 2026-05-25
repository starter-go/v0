package codecs

import (
	"github.com/starter-go/v0/htttest"
	"github.com/starter-go/v0/htttest/api/h3t"
)

type HtmlCodec struct {

	//starter:component

	_as func(htttest.CodecProvider) //starter:as(".")

}

// Registrations implements h3t.CodecProvider.
func (inst *HtmlCodec) Registrations() []*h3t.CodecRegistration {
	r1 := &h3t.CodecRegistration{

		Name:     "html.codec",
		Type:     "text/html",
		Enabled:  true,
		Priority: 0,

		Encoder: inst,
		Decoder: inst,
		Codec:   inst,
	}
	return []*h3t.CodecRegistration{r1}
}

// Decode implements h3t.Codec.
func (inst *HtmlCodec) Decode(c *h3t.Content) error {
	txt := ""
	bin := c.Encoded
	if bin != nil {
		txt = string(bin)
	}
	c.Data = txt
	return nil
}

// Encode implements h3t.Codec.
func (inst *HtmlCodec) Encode(c *h3t.Content) error {
	src := c.Data
	str, ok := src.(string)
	if ok {
		c.Encoded = []byte(str)
	} else {
		c.Encoded = []byte{}
	}
	return nil
}

func (inst *HtmlCodec) _impl() (htttest.CodecProvider, htttest.Codec) {
	return inst, inst
}
