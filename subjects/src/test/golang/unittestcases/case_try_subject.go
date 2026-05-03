package unittestcases

import (
	"context"

	"github.com/starter-go/application"
	"github.com/starter-go/v0/subjects"
	"github.com/starter-go/v0/subjects/core/consoles"
)

type CaseToTrySubject struct {

	//starter:component

	ChainHolder subjects.FilterChainHolder //starter:inject("#")

	Enabled bool //starter:inject("${test-case.subjects.enabled}")

}

func (inst *CaseToTrySubject) Life() *application.Life {

	l := new(application.Life)

	if inst.Enabled {
		l.OnLoop = inst.run
	}

	return l
}

func (inst *CaseToTrySubject) prepareContext() (context.Context, error) {

	ada := new(consoles.SubjectAdapter)
	ctx := new(subjects.Context)
	ch := inst.ChainHolder

	wtr, err1 := ch.GetWriterChain()
	rdr, err2 := ch.GetReaderChain()
	if err1 != nil {
		return nil, err1
	}
	if err2 != nil {
		return nil, err2
	}

	ctx.Writer = wtr
	ctx.Reader = rdr
	ctx.ChainHolder = ch
	ctx = ada.Setup(ctx)

	return ctx.CC, nil
}

func (inst *CaseToTrySubject) run() error {

	cc, err := inst.prepareContext()
	if err != nil {
		return err
	}

	sub, err := subjects.GetCurrent(cc)
	if err != nil {
		return err
	}

	sub.GetProperties()

	return nil
}
