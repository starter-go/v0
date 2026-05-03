package subjects

import "sort"

////////////////////////////////////////////////////////////////////////////////

type ReadFilterChainBuilder struct {
	items []*FilterRegistration
}

func (inst *ReadFilterChainBuilder) Build() ReadFilterChain {

	inst.sort()

	var chain ReadFilterChain
	chain = new(innerReaderChainEnd)
	all := inst.items

	for _, it := range all {
		node := new(innerReaderChainNode)
		node.filter = it.Reader
		node.next = chain
		chain = node
	}

	return chain
}

func (inst *ReadFilterChainBuilder) AddRegistry(r FilterRegistry) {
	if r == nil {
		return
	}
	list := r.GetRegistrationList()
	for _, it := range list {
		inst.AddRegistration(it)
	}
}

func (inst *ReadFilterChainBuilder) AddRegistration(r *FilterRegistration) {
	if !inst.accept(r) {
		return
	}
	inst.items = append(inst.items, r)
}

func (inst *ReadFilterChainBuilder) accept(r *FilterRegistration) bool {

	if r == nil {
		return false
	}

	if r.Reader == nil {
		return false
	}

	if !r.Enabled {
		return false
	}

	return true
}

func (inst *ReadFilterChainBuilder) sort() {
	sort.Sort(inst)
}

func (inst *ReadFilterChainBuilder) Len() int {
	return len(inst.items)
}

func (inst *ReadFilterChainBuilder) Less(i1, i2 int) bool {
	o1 := inst.items[i1]
	o2 := inst.items[i2]
	return o1.Priority < o2.Priority
}

func (inst *ReadFilterChainBuilder) Swap(i1, i2 int) {
	list := inst.items
	list[i1], list[i2] = list[i2], list[i1]
}

////////////////////////////////////////////////////////////////////////////////

type innerReaderChainNode struct {
	filter ReadFilter
	next   ReadFilterChain
}

// Read implements FilterChain.
func (inst *innerReaderChainNode) Read(c *Context) error {

	n := inst.next
	f := inst.filter
	return f.Read(c, n)
}

func (inst *innerReaderChainNode) _impl() ReadFilterChain {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

type innerReaderChainEnd struct {
}

// Read implements FilterChain.
func (inst *innerReaderChainEnd) Read(c *Context) error {
	return nil
}

func (inst *innerReaderChainEnd) _impl() ReadFilterChain {
	return inst
}

////////////////////////////////////////////////////////////////////////////////
// EOF
