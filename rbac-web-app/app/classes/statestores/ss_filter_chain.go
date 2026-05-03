package statestores

import "sort"

////////////////////////////////////////////////////////////////////////////////

type FilterChain interface {
	Store
}

type FilterChainHolder interface {
	GetChain() (FilterChain, error)
}

type FilterChainLoader interface {
	LoadChain() (FilterChain, error)
}

////////////////////////////////////////////////////////////////////////////////
// FilterChainBuilder

type FilterChainBuilder struct {
	items []*FilterRegistration
}

func (inst *FilterChainBuilder) Build() FilterChain {

	all := inst.sort()
	var chain FilterChain
	chain = new(innerStoreChainEnd)

	for _, it := range all {
		n := new(innerStoreChainNode)
		n.filter = it.Filter
		n.next = chain
		chain = n
	}

	return chain
}

func (inst *FilterChainBuilder) accept(r *FilterRegistration) bool {

	if r == nil {
		return false
	}

	if !r.Enabled {
		return false
	}

	if r.Filter == nil {
		return false
	}

	return true
}

func (inst *FilterChainBuilder) AddRegistration(r *FilterRegistration) {
	if inst.accept(r) {
		inst.items = append(inst.items, r)
	}
}

func (inst *FilterChainBuilder) AddRegistry(r FilterRegistry) {
	if r == nil {
		return
	}
	src := r.ListRegistrations()
	for _, it := range src {
		inst.AddRegistration(it)
	}
}

func (inst *FilterChainBuilder) sort() []*FilterRegistration {
	sort.Sort(inst)
	return inst.items
}

func (inst *FilterChainBuilder) Len() int {
	return len(inst.items)
}
func (inst *FilterChainBuilder) Less(i1, i2 int) bool {
	o1 := inst.items[i1]
	o2 := inst.items[i2]
	return o1.Priority < o2.Priority
}
func (inst *FilterChainBuilder) Swap(i1, i2 int) {
	list := inst.items
	list[i1], list[i2] = list[i2], list[i1]
}

////////////////////////////////////////////////////////////////////////////////
// innerStoreChainNode

type innerStoreChainNode struct {
	filter Filter
	next   FilterChain
}

// Handle implements StoreChain.
func (inst *innerStoreChainNode) Handle(sc *Context) error {

	f := inst.filter
	n := inst.next

	return f.Handle(sc, n)
}

func (inst *innerStoreChainNode) _impl() FilterChain {
	return inst
}

////////////////////////////////////////////////////////////////////////////////
// innerStoreChainEnd

type innerStoreChainEnd struct {
}

// Handle implements StoreChain.
func (inst *innerStoreChainEnd) Handle(sc *Context) error {
	// NOP
	return nil
}

func (inst *innerStoreChainEnd) _impl() FilterChain {
	return inst
}

////////////////////////////////////////////////////////////////////////////////
// EOF
