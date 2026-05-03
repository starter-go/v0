package consoles

import (
	"context"
	"fmt"

	"github.com/starter-go/v0/subjects"
)

type SubjectAdapter struct {
}

// GetHolder implements subjects.Adapter.
func (inst *SubjectAdapter) GetHolder(c context.Context) (*subjects.Holder, error) {

	name := inst.innerGetHolderName()
	val := c.Value(name)
	h, ok := val.(*subjects.Holder)

	if !ok {
		return nil, fmt.Errorf("consoles.SubjectAdapter() : no subjects.Holder in context")
	}

	return h, nil
}

func (inst *SubjectAdapter) Setup(sc *subjects.Context) *subjects.Context {

	cc := sc.CC
	if cc == nil {
		cc = context.Background()
	}

	holder := new(subjects.Holder)
	name4h := inst.innerGetHolderName()
	holder.Context = sc

	subjects.SetupAdapter(func(name4a string) {
		cc = context.WithValue(cc, name4a, inst)
		cc = context.WithValue(cc, name4h, holder)
	})

	if sc.Facade == nil {
		sub := subjects.NewSubject(sc)
		sc.Facade = sub
	}

	sc.CC = cc
	return sc
}

func (inst *SubjectAdapter) innerGetHolderName() string {
	return "github.com/starter-go/v0/subjects/consoles#SubjectAdapter/holder"
}

func (inst *SubjectAdapter) _impl() subjects.Adapter {
	return inst
}
