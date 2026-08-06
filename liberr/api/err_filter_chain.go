package api

////////////////////////////////////////////////////////////////////////////////

type FilterChainBuilder struct {
}

func (inst *FilterChainBuilder) Build() FilterChain {

}

////////////////////////////////////////////////////////////////////////////////

type innerFilterChainNode struct {
	next FilterChain
	fi   Filter
}

// Error implements [FilterChain].
func (inst *innerFilterChainNode) Error(w *Want) error {
	panic("unimplemented")
}

func (inst *innerFilterChainNode) _impl() FilterChain {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

type innerFilterChainEnd struct {
}

// Error implements [FilterChain].
func (inst *innerFilterChainEnd) Error(w *Want) error {
	panic("unimplemented")
}

func (inst *innerFilterChainEnd) _impl() FilterChain {
	return inst
}

////////////////////////////////////////////////////////////////////////////////
// EOF
