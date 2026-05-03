package autoloaders

import (
	"github.com/starter-go/v0/subjects"
)

////////////////////////////////////////////////////////////////////////////////

type DefaultFilterChainHolder struct {

	//starter:component

	_as func(subjects.FilterChainHolder) //starter:as("#")

	Loader subjects.FilterChainLoader //starter:inject("#")

	chainR subjects.ReadFilterChain
	chainW subjects.WriteFilterChain
}

// GetReaderChain implements subjects.FilterChainHolder.
func (inst *DefaultFilterChainHolder) GetReaderChain() (subjects.ReadFilterChain, error) {

	ch := inst.chainR

	if ch == nil {
		tmp, err := inst.Loader.LoadReaderChain()
		if err != nil {
			return nil, err
		}
		ch = tmp
		inst.chainR = tmp
	}

	return ch, nil
}

// GetWriterChain implements subjects.FilterChainHolder.
func (inst *DefaultFilterChainHolder) GetWriterChain() (subjects.WriteFilterChain, error) {

	ch := inst.chainW

	if ch == nil {
		tmp, err := inst.Loader.LoadWriterChain()
		if err != nil {
			return nil, err
		}
		ch = tmp
		inst.chainW = tmp
	}

	return ch, nil
}

func (inst *DefaultFilterChainHolder) _impl() subjects.FilterChainHolder {
	return inst
}

////////////////////////////////////////////////////////////////////////////////
