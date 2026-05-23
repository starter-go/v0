package jwt

import (
	"context"
	"fmt"
)

type Access struct {
	Context context.Context

	Token *Token

	Text Text

	Service Service
}

func (inst *Access) Encode(force bool) error {

	if inst == nil {
		return fmt.Errorf("jwt.Access: this is nil")
	}

	dst := inst.Text
	src := inst.Token
	ser := inst.Service

	if (dst != "") && (!force) {
		return nil
	}

	if src == nil {
		return fmt.Errorf("jwt.Access: source (Token) is nil")
	}

	if ser == nil {
		return fmt.Errorf("jwt.Access: service is nil")
	}

	enc := ser.GetEncoder()
	result, err := enc.Encode(src)
	if err != nil {
		return err
	}

	inst.Text = result
	return nil
}

func (inst *Access) Decode(force bool) error {

	if inst == nil {
		return fmt.Errorf("jwt.Access: this is nil")
	}

	src := inst.Text
	dst := inst.Token
	ser := inst.Service

	if (dst != nil) && (!force) {
		return nil
	}

	if src == "" {
		return fmt.Errorf("jwt.Access: source (Text) is empty")
	}

	if ser == nil {
		return fmt.Errorf("jwt.Access: service is nil")
	}

	dec := ser.GetDecoder()
	result, err := dec.Decode(src)
	if err != nil {
		return err
	}

	inst.Token = result
	return nil
}
