package unittestcases

import (
	"fmt"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/units"
	"github.com/starter-go/v0/libjwt/api/jwt"
	"github.com/starter-go/vlog"
)

type CaseTryJwtCodec struct {

	//starter:component

	Service jwt.Service //starter:inject("#")

}

// ListRegistrations implements units.Unit.
func (inst *CaseTryJwtCodec) ListRegistrations(list []*units.Registration) []*units.Registration {

	const name = "try-jwt-codec"

	ur1 := &units.Registration{

		ID:          name,
		Name:        name,
		Class:       "github.com/starter-go/v0/libjwt/src/test/golang/unittestcases#CaseTryJwtCodec",
		Description: "",

		Enabled:  true,
		Priority: 0,
		OnError:  units.OnErrorError,

		Do: inst.run,
	}

	list = append(list, ur1)
	return list
}

func (inst *CaseTryJwtCodec) run() error {

	ser := inst.Service
	dec := ser.GetDecoder()
	enc := ser.GetEncoder()
	pt := make(map[string]string)
	now := lang.Now()
	const span = 3600 * 1000

	pt["foo"] = "bar"

	token1 := &jwt.Token{

		SessionID:   1234567890,
		SessionUUID: "f1385427-0572-720f-ccee-82164dbd7f5c",

		Properties: pt,

		NotBefore: now - span,
		NotAfter:  now + span,
	}

	txt1, err := enc.Encode(token1)
	if err != nil {
		return err
	}

	token2, err := dec.Decode(txt1)
	if err != nil {
		return err
	}

	txt2, err := enc.Encode(token2)
	if err != nil {
		return err
	}

	if txt1 != txt2 {
		return fmt.Errorf("txt1 != txt2")
	}

	/// log results

	vlog.Debug("jwt.text-1 = %v", txt1)
	vlog.Debug("jwt.text-2 = %v", txt2)
	vlog.Debug("jwt.token  = %v", token2)

	return nil
}

func (inst *CaseTryJwtCodec) _impl() units.Unit {
	return inst
}
