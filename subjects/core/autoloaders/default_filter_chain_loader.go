package autoloaders

import (
	"github.com/starter-go/v0/subjects"
)

////////////////////////////////////////////////////////////////////////////////

type DefaultFilterChainLoader struct {

	//starter:component

	_as func(subjects.FilterChainLoader) //starter:as("#")

	RegList []subjects.FilterRegistry //starter:inject(".")

}

// LoadReaderChain implements subjects.FilterChainLoader.
func (inst *DefaultFilterChainLoader) LoadReaderChain() (subjects.ReadFilterChain, error) {

	b := new(subjects.ReadFilterChainBuilder)
	src := inst.RegList

	for _, it := range src {
		b.AddRegistry(it)
	}

	chain := b.Build()
	return chain, nil

}

// LoadWriterChain implements subjects.FilterChainLoader.
func (inst *DefaultFilterChainLoader) LoadWriterChain() (subjects.WriteFilterChain, error) {

	b := new(subjects.WriteFilterChainBuilder)
	src := inst.RegList

	for _, it := range src {
		b.AddRegistry(it)
	}

	chain := b.Build()
	return chain, nil

}

func (inst *DefaultFilterChainLoader) _impl() subjects.FilterChainLoader {
	return inst
}

////////////////////////////////////////////////////////////////////////////////
