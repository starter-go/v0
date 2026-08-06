package pack1

import (
	"github.com/starter-go/v0/liberr"
	"github.com/starter-go/v0/liberr/api"
)

const (
	theDemoNS api.Namespace = "github.com/starter-go/v0/liberr/src/test/golang/pack1"

	theP1DemoName1 api.Name = "err_1"
	theP1DemoName2 api.Name = "err_2"
	theP1DemoName3 api.Name = "err_3"
)

////////////////////////////////////////////////////////////////////////////////

type Demo1ErrorSet struct {

	//starter:component

	_as func(liberr.ErrorSet) //starter:as(".")

	holder liberr.ErrorSetHolder
}

// ListErrors implements [api.ErrorSet].
func (inst *Demo1ErrorSet) ListErrors() []*api.Registration {
	panic("unimplemented")
}

// Load implements [api.ErrorSetLoader].
func (inst *Demo1ErrorSet) Load() api.ErrorSet {
	return inst
}

func (inst *Demo1ErrorSet) innerMakeError1(args ...any) error {

	return nil
}

func (inst *Demo1ErrorSet) MakeError1(args ...any) error {

	return nil
}

// args : a,b,c
func (inst *Demo1ErrorSet) MakeError2(args ...any) error {
	h := &inst.holder
	ns := theDemoNS
	return h.From(inst).ErrorNS(ns, theP1DemoName2, args...)
}

func (inst *Demo1ErrorSet) _impl() (liberr.ErrorSet, liberr.ErrorSetLoader) {
	return inst, inst
}

////////////////////////////////////////////////////////////////////////////////
