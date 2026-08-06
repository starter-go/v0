package pack2

import (
	"github.com/starter-go/v0/liberr"
	"github.com/starter-go/v0/liberr/api"
)

const (
	theP2DemoNS api.Namespace = "github.com/starter-go/v0/liberr/src/test/golang/pack2"

	theP2DemoName1 api.Name = "err_1"
	theP2DemoName2 api.Name = "err_2"
	theP2DemoName3 api.Name = "err_3"
)

////////////////////////////////////////////////////////////////////////////////

type Demo2ErrorSet struct {

	//starter:component

	_as func(liberr.ErrorSet) //starter:as(".")

	holder liberr.ErrorSetHolder
}

// ListErrors implements [api.ErrorSet].
func (inst *Demo2ErrorSet) ListErrors() []*api.Registration {
	panic("unimplemented")
}

// Load implements [api.ErrorSetLoader].
func (inst *Demo2ErrorSet) Load() api.ErrorSet {
	return inst
}

func (inst *Demo2ErrorSet) innerMakeError1(args ...any) error {

	return nil
}

func (inst *Demo2ErrorSet) MakeError1(args ...any) error {

	return nil
}

// args : a,b,c
func (inst *Demo2ErrorSet) MakeError2(args ...any) error {
	h := &inst.holder
	ns := theP2DemoNS
	return h.From(inst).ErrorNS(ns, theP2DemoName2, args...)
}

func (inst *Demo2ErrorSet) _impl() (liberr.ErrorSet, liberr.ErrorSetLoader) {
	return inst, inst
}

////////////////////////////////////////////////////////////////////////////////
