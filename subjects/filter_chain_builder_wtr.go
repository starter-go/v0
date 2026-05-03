package subjects

import "sort"

////////////////////////////////////////////////////////////////////////////////

type WriteFilterChainBuilder struct {
	items []*FilterRegistration
}

func (inst *WriteFilterChainBuilder) Build() WriteFilterChain {

	inst.sort()

	var chain WriteFilterChain
	chain = new(innerWriterChainEnd)
	all := inst.items

	for _, it := range all {
		node := new(innerWriterChainNode)
		node.filter = it.Writer
		node.next = chain
		chain = node
	}

	return chain
}

func (inst *WriteFilterChainBuilder) AddRegistry(r FilterRegistry) {
	if r == nil {
		return
	}
	list := r.GetRegistrationList()
	for _, it := range list {
		inst.AddRegistration(it)
	}
}

func (inst *WriteFilterChainBuilder) AddRegistration(r *FilterRegistration) {
	if !inst.accept(r) {
		return
	}
	inst.items = append(inst.items, r)
}

func (inst *WriteFilterChainBuilder) accept(r *FilterRegistration) bool {

	if r == nil {
		return false
	}

	if r.Writer == nil {
		return false
	}

	if !r.Enabled {
		return false
	}

	return true
}

func (inst *WriteFilterChainBuilder) sort() {
	sort.Sort(inst)
}

func (inst *WriteFilterChainBuilder) Len() int {
	return len(inst.items)
}

func (inst *WriteFilterChainBuilder) Less(i1, i2 int) bool {
	o1 := inst.items[i1]
	o2 := inst.items[i2]
	return o1.Priority < o2.Priority
}

func (inst *WriteFilterChainBuilder) Swap(i1, i2 int) {
	list := inst.items
	list[i1], list[i2] = list[i2], list[i1]
}

////////////////////////////////////////////////////////////////////////////////

type innerWriterChainNode struct {
	filter WriteFilter
	next   WriteFilterChain
}

// Write implements FilterChain.
func (inst *innerWriterChainNode) Write(c *Context) error {

	n := inst.next
	f := inst.filter
	return f.Write(c, n)
}

func (inst *innerWriterChainNode) _impl() WriteFilterChain {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

type innerWriterChainEnd struct {
}

// Write implements FilterChain.
func (inst *innerWriterChainEnd) Write(c *Context) error {
	return nil
}

func (inst *innerWriterChainEnd) _impl() WriteFilterChain {
	return inst
}

////////////////////////////////////////////////////////////////////////////////
// EOF
