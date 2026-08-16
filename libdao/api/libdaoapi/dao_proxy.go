package libdaoapi

import (
	"fmt"
	"sort"
	"strings"

	"github.com/starter-go/base/lang"
)

////////////////////////////////////////////////////////////////////////////////

type DAO interface {
	GetRegistration() *DaoRegistration
}

type DaoRegistration struct {
	ID string

	Name string

	Class string

	Enabled bool

	Priority int

	DAO DAO
}

////////////////////////////////////////////////////////////////////////////////

type innerDaoCache[T any] struct {
	to   T
	info *DaoRegistration
}

////////////////////////////////////////////////////////////////////////////////

type innerDaoLoader[T DAO] struct {
	items []*DaoRegistration
	all   []*DaoRegistration
}

// Len implements [sort.Interface].
func (inst *innerDaoLoader[T]) Len() int {
	return len(inst.items)
}

// Less implements [sort.Interface].
func (inst *innerDaoLoader[T]) Less(i int, j int) bool {
	o1 := inst.items[i]
	o2 := inst.items[j]
	return (o1.Priority > o2.Priority)
}

// Swap implements [sort.Interface].
func (inst *innerDaoLoader[T]) Swap(i int, j int) {
	l := inst.items
	l[i], l[j] = l[j], l[i]
}

func (inst *innerDaoLoader[T]) add(item T) {
	info := item.GetRegistration()
	if info == nil {
		return
	}
	if info.Enabled && info.DAO != nil {
		inst.items = append(inst.items, info)
	}
	inst.all = append(inst.all, info)
}

func (inst *innerDaoLoader[T]) sort() sort.Interface {
	sort.Sort(inst)
	return inst
}

func (inst *innerDaoLoader[T]) load(sel string) (T, *DaoRegistration, error) {
	info, err := inst.select1(sel)
	if err != nil {
		var x T
		return x, nil, err
	}
	t := info.DAO.(T)
	return t, info, err
}

func (inst *innerDaoLoader[T]) select1(sel string) (*DaoRegistration, error) {

	inst.sort()
	all := inst.items

	if strings.HasPrefix(sel, "#") {
		// by id
		for _, it := range all {
			id := it.ID
			if sel == ("#" + id) {
				return it, nil
			}
		}
	} else if strings.HasPrefix(sel, ".") {
		// by class
		for _, it := range all {
			if inst.isClassOf(sel, it) {
				return it, nil
			}
		}
	} else if sel != "" {
		// use name
		for _, it := range all {
			name := it.Name
			if sel == name {
				return it, nil
			}
		}
	} else {
		// use first
		for _, it := range all {
			return it, nil
		}
	}

	// return error
	var api T
	cl := lang.ClassOf(&api)
	const msg = "no dao match selector"
	const f = "[error msg:'%s' selector:'%s' dao:'%s' ]"
	return nil, fmt.Errorf(f, msg, sel, cl.FullName())
}

func (inst *innerDaoLoader[T]) isClassOf(sel string, item *DaoRegistration) bool {

	const space = ' '
	const sep = '\n'

	str := strings.ReplaceAll(item.Class, string(space), string(sep))
	classlist := strings.Split(str, string(sep))

	for _, cl := range classlist {
		cl = strings.TrimSpace(cl)
		if cl == "" {
			continue
		}
		if sel == ("." + cl) {
			return true
		}
	}

	return false
}

////////////////////////////////////////////////////////////////////////////////

type DaoHolder[T DAO] struct {
	cache *innerDaoCache[T]
}

func (inst *DaoHolder[T]) Reset() {
	inst.cache = nil
}

func (inst *DaoHolder[T]) Select(sel string, all []T) T {

	target, err := inst.innerGetTarget(sel, all)
	if err != nil {
		panic(err)
	}
	return target.to
}

func (inst *DaoHolder[T]) innerGetTarget(sel string, all []T) (*innerDaoCache[T], error) {

	ca := inst.cache
	if ca != nil {
		return ca, nil
	}

	// load

	ca, err := inst.innerLoadTarget(sel, all)
	if err == nil {
		inst.cache = ca
	}
	return ca, err
}

func (inst *DaoHolder[T]) innerLoadTarget(sel string, all []T) (*innerDaoCache[T], error) {

	loader := new(innerDaoLoader[T])
	for _, item := range all {
		loader.add(item)
	}

	t, info, err := loader.load(sel)
	if err != nil {
		return nil, err
	}

	cache := new(innerDaoCache[T])
	cache.info = info
	cache.to = t
	return cache, nil
}

////////////////////////////////////////////////////////////////////////////////
// EOF
