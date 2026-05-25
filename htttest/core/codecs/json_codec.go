package codecs

import (
	"encoding/json"

	"github.com/starter-go/v0/htttest"
	"github.com/starter-go/v0/htttest/api/h3t"
)

type JsonCodec struct {

	//starter:component

	_as func(htttest.CodecProvider) //starter:as(".")

}

// Registrations implements h3t.CodecProvider.
func (inst *JsonCodec) Registrations() []*h3t.CodecRegistration {

	ctype := inst.innerGetMyContentType()

	r1 := &h3t.CodecRegistration{

		Name:     "json.codec",
		Type:     ctype,
		Enabled:  true,
		Priority: 0,

		Encoder: inst,
		Decoder: inst,
		Codec:   inst,
	}
	return []*h3t.CodecRegistration{r1}
}

func (inst *JsonCodec) innerGetMyContentType() htttest.ContentType {
	return "application/json"
}

// Decode implements h3t.Codec.
func (inst *JsonCodec) Decode(c *h3t.Content) error {

	bin := c.Encoded
	data := c.Data
	err := json.Unmarshal(bin, data)
	if err != nil {
		return err
	}

	size := len(bin)
	ctype := inst.innerGetMyContentType()

	c.Type = ctype
	c.Length = int64(size)

	return nil
}

// Encode implements h3t.Codec.
func (inst *JsonCodec) Encode(c *h3t.Content) error {

	data := c.Data
	bin, err := json.Marshal(data)

	if err != nil {
		return err
	}

	size := len(bin)
	ctype := inst.innerGetMyContentType()

	c.Type = ctype
	c.Length = int64(size)
	c.Encoded = bin

	return nil
}

func (inst *JsonCodec) _impl() (htttest.CodecProvider, htttest.Codec) {
	return inst, inst
}
