package unittestcases

import (
	"context"

	"github.com/starter-go/units"
	"github.com/starter-go/v0/subjects/core/classes/sessions"
	"github.com/starter-go/vlog"
)

type TrySessionService struct {

	//starter:component

	Service sessions.Service //starter:inject("#")

	Enabled bool //starter:inject("${test-case.session-service.enabled}")

	id sessions.ID
}

// ListRegistrations implements units.Unit.
func (inst *TrySessionService) ListRegistrations(list []*units.Registration) []*units.Registration {

	ur := &units.Registration{
		Name:    "TrySessionService",
		Enabled: true,
		Do:      inst.run,
	}
	list = append(list, ur)
	return list
}

func (inst *TrySessionService) _impl() units.Unit {
	return inst
}

func (inst *TrySessionService) run() error {

	vlog.Debug("TrySessionService.run()")

	steps := make([]func(c context.Context) error, 0)
	ctx := inst.innerPrepareContext()

	// steps

	steps = append(steps, inst.innerDoInsert)
	steps = append(steps, inst.innerDoFind)

	steps = append(steps, inst.innerDoUpdate)
	steps = append(steps, inst.innerDoQuery)
	steps = append(steps, inst.innerDoDelete)

	// loop

	for idx, step := range steps {

		vlog.Debug(" ... innerDoStep[%d]", idx)

		err := step(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *TrySessionService) innerPrepareContext() context.Context {

	ctx := context.Background()
	return ctx
}

func (inst *TrySessionService) innerDoQuery(ctx context.Context) error {

	vlog.Info("innerDoQuery () ...")

	ser := inst.Service
	q := &sessions.Query{}

	list, err := ser.Query(ctx, q)
	if err != nil {
		return err
	}

	for _, it := range list {
		vlog.Debug("item = %v", it)
	}

	return nil
}

func (inst *TrySessionService) innerDoFind(ctx context.Context) error {

	vlog.Info("innerDoFind () ...")

	ser := inst.Service
	id := inst.id

	o1, err := ser.Find(ctx, id)
	if err != nil {
		return err
	}

	vlog.Debug("o1 = %v", o1)

	return nil
}

func (inst *TrySessionService) innerDoInsert(ctx context.Context) error {

	vlog.Info("innerDoInsert () ...")

	ser := inst.Service
	o1 := new(sessions.DTO)

	pm := o1.Properties
	pm = pm.Init()
	pm.Put("aa", "1")
	pm.Put("bb", "2")
	pm.Put("cc", "3")
	o1.Properties = pm

	o2, err := ser.Insert(ctx, o1)
	if err == nil && o2 != nil {
		inst.id = o2.ID
	}

	return err
}

func (inst *TrySessionService) innerDoUpdate(ctx context.Context) error {

	vlog.Info("innerDoUpdate () ...")

	ser := inst.Service
	id := inst.id

	o1, err := ser.Find(ctx, id)
	if err != nil {
		return err
	}

	pm := o1.Properties
	pm = pm.Init()
	// pm.Put("aa", "1")
	// pm.Put("bb", "2")
	pm.Put("cc", "3.14")
	pm.Put("x", "77")
	pm.Put("y", "88")
	pm.Put("z", "99")
	o1.Properties = pm

	o2, err := ser.Update(ctx, id, o1)
	if err != nil {
		return err
	}

	vlog.Debug("o1 = %v", o1)
	vlog.Debug("o2 = %v", o2)

	return err
}

func (inst *TrySessionService) innerDoDelete(ctx context.Context) error {

	vlog.Info("innerDoDelete () ...")

	ser := inst.Service
	id := inst.id

	return ser.Remove(ctx, id)
}
