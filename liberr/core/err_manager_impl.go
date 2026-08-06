package core

import (
	"fmt"

	"github.com/starter-go/application"
	"github.com/starter-go/v0/liberr"
	"github.com/starter-go/v0/liberr/api"
	"github.com/starter-go/vlog"
)

////////////////////////////////////////////////////////////////////////////////

type ErrorManagerImpl struct {

	//starter:component

	_as func(liberr.ErrorManager) //starter:as("#")

	PropLogOnStart bool //starter:inject("${liberr.log-on-start}")
	PropReCode     bool //starter:inject("${liberr.re-code}")
	PropCodeMask   int  //starter:inject("${liberr.code-mask}")

	ErrSets []liberr.ErrorSet //starter:inject(".")

	cache *innerErrorManagerCache
}

func (inst *ErrorManagerImpl) _impl() (application.Lifecycle, liberr.ErrorManager) {
	return inst, inst
}

// ListErrors implements api.ErrorManager.
func (inst *ErrorManagerImpl) ListErrors() []*api.Registration {
	c := inst.innerGetCache()
	return c.copyItems()
}

func (inst *ErrorManagerImpl) Life() *application.Life {
	l := new(application.Life)
	l.Order = -99

	l.OnCreate = inst.onCreate
	l.OnStart = inst.onStart

	return l
}

func (inst *ErrorManagerImpl) onCreate() error {
	inst.innerGetCache()
	return nil
}

func (inst *ErrorManagerImpl) onStart() error {
	return inst.innerLog()
}

func (inst *ErrorManagerImpl) innerLog() error {

	if inst.PropLogOnStart {
		vlog.Info("list of Liberr.Errors:")
	} else {
		return nil
	}

	c := inst.innerGetCache()
	all := c.all

	for idx, it := range all {
		vlog.Info("    [Error(%d) code:'%d' ns:'%s' name:'%s']", idx, it.Code, it.Namespace, it.Name)
	}

	return nil
}

func (inst *ErrorManagerImpl) innerGetCache() *innerErrorManagerCache {
	c := inst.cache
	if c == nil {
		c2 := new(innerErrorManagerCache)
		c2.man = inst
		err := c2.load()
		if err != nil {
			panic(err)
		}
		inst.cache = c
	}
	return c
}

////////////////////////////////////////////////////////////////////////////////

type innerErrorManagerCache struct {
	man *ErrorManagerImpl
	all []*api.Registration
}

func (inst *innerErrorManagerCache) load() error {

	inst.all = nil

	src := inst.man.ErrSets
	dst := inst.all

	for _, it := range src {
		tmp := it.ListErrors()
		dst = append(dst, tmp...)
	}

	inst.filterEmpty(dst)
	inst.rebuildCodeForItems(dst)

	inst.all = dst
	return nil
}

func (inst *innerErrorManagerCache) rebuildCodeForItems(items []*api.Registration) {

}

func (inst *innerErrorManagerCache) rebuildCodeForItem(item *api.Registration) {

}

func (inst *innerErrorManagerCache) filterEmpty(src []*api.Registration) []*api.Registration {
	dst := make([]*api.Registration, 0)
	for _, reg := range src {
		if reg == nil {
			continue
		}

		ff := reg.GetFormatter()
		example := reg.Example

		if ff == nil {
			ff = liberr.DefaultFormatter()
			reg.Formatter = ff
		}

		if example == nil {
			args := inst.makeMockArgs(reg)
			err := ff.Format(args...)
			reg.Example = err
		}

		dst = append(dst, reg)
	}
	return dst
}

func (inst *innerErrorManagerCache) copyItems() []*api.Registration {
	src := inst.all
	dst := make([]*api.Registration, len(src))
	copy(dst, src)
	return dst
}

func (inst *innerErrorManagerCache) makeMockArgs(reg *api.Registration) []any {
	count := 0
	bin := []byte(reg.Format)
	for _, b := range bin {
		if b == '%' {
			count++
		}
	}
	args := make([]any, count)
	for i := 0; i < count; i++ {
		args[i] = fmt.Sprintf("{a%d}", i)
	}
	return args
}

////////////////////////////////////////////////////////////////////////////////
// EOF
