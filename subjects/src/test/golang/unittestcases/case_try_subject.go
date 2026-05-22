package unittestcases

import (
	"context"
	"encoding/json"

	"github.com/starter-go/base/context2"
	"github.com/starter-go/rbac"
	"github.com/starter-go/units"
	"github.com/starter-go/v0/subjects"
	"github.com/starter-go/vlog"
)

type CaseToTrySubject struct {

	//starter:component

	ChainHolder subjects.FilterChainHolder //starter:inject("#")

	Enabled bool //starter:inject("${unit.test-subjects-1.enabled}")

}

// ListRegistrations implements units.Unit.
func (inst *CaseToTrySubject) ListRegistrations(list []*units.Registration) []*units.Registration {

	u1 := &units.Registration{
		Name:    "CaseToTrySubject",
		Enabled: inst.Enabled,
		Do:      inst.run,
	}

	list = append(list, u1)
	return list
}

func (inst *CaseToTrySubject) _impl() units.Unit {
	return inst
}

func (inst *CaseToTrySubject) prepareContext() (context.Context, error) {

	cc := context.Background()
	cc = context2.SetupConsoleContext(cc)
	c2h, err := context2.GetHolder(cc)

	ada, err := subjects.GetAdapter(c2h.Context())
	if err != nil {
		return nil, err
	}

	holder, err := ada.GetHolder(cc)
	if err != nil {
		return nil, err
	}

	ctx := holder.Context
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

	// do get(1)
	inst.innerTryGetL1(cc)

	// gett, err := sub.DoGet()
	// if err != nil {
	// 	return err
	// }

	// do insert
	vlog.Debug("do: subject.insert()")
	inst.innerTryGetL1(cc)

	sett, err := sub.DoSet()
	if err != nil {
		return err
	}
	sett.SetProperty("foo", "f-value-1")
	sett.SetProperty("bar", "b-value-1")

	err = sub.Create()
	if err != nil {
		return err
	}

	err = sub.Save()
	if err != nil {
		return err
	}

	// do update
	vlog.Debug("do: subject.update()")
	inst.innerTryGetL1(cc)

	sett, err = sub.DoSet()
	if err != nil {
		return err
	}
	sett.SetProperty("foo", "f-value-2")
	sett.SetProperty("ccc", "c-value-1")
	sett.SetProperty("ddd", "d-value-1")

	err = sub.Save()
	if err != nil {
		return err
	}

	err = sub.Reload()
	if err != nil {
		return err
	}

	// do get(2)
	inst.innerTryGetL1(cc)
	inst.innerTryGetL1(cc)
	inst.innerTryGetL1(cc)

	// do exit
	inst.innerTryGetL1(cc)

	err = sub.Exit()
	if err != nil {
		return err
	}

	// do get(3)
	inst.innerTryGetL1(cc)

	return nil
}

func (inst *CaseToTrySubject) innerTryGetL1(cc context.Context) {

	err := inst.innerTryGetL2(cc)
	if err == nil {
		return
	}
	vlog.Warn("CaseToTrySubject.innerTryGetL1() : error = %s", err.Error())
}

func (inst *CaseToTrySubject) innerTryGetL2(cc context.Context) error {

	sub, err := subjects.GetCurrent(cc)
	if err != nil {
		return err
	}

	gett, err := sub.DoGet()
	if err != nil {
		return err
	}
	session := new(rbac.SessionDTO)
	gett.GetSession(session)

	str, err := json.MarshalIndent(session, "\t", "\t")
	if err != nil {
		return err
	}

	vlog.Info("subject.GetSession() : result = %s", str)
	return nil
}
