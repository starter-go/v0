package iadapter

import (
	"fmt"
	"sort"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/v0/libjwt/api/jwt"
)

////////////////////////////////////////////////////////////////////////////////

type MainAdapter struct {

	//starter:component

	_as func(jwt.Adapter) //starter:as("#")

	ProviderList []jwt.AdapterProvider //starter:inject(".")

	AutoTermEnabled bool  //starter:inject("${jwt.auto-term.enabled}")
	AutoTermAge     int64 //starter:inject("${jwt.auto-term.age-in-ms}")

	cache *innerAdapterCache
}

// GetToken implements jwt.Adapter.
func (inst *MainAdapter) GetToken(acc *jwt.Access) error {

	acc.Text = ""
	acc.Token = nil

	ok := false
	var lastErr error

	cache, err := inst.innerGetCache()
	if err != nil {
		return err
	}

	for _, reg := range cache.items {
		ada := reg.Adapter
		err := ada.GetToken(acc)
		if err == nil {
			ok = true
			break
		} else {
			lastErr = err
		}
	}

	if !ok {
		if lastErr == nil {
			return fmt.Errorf("jwt.MainAdapter: no Adapter handle GetToken()")
		} else {
			return lastErr
		}
	}

	err = acc.Decode(false)
	if err != nil {
		return err
	}

	return nil
}

// SetToken implements jwt.Adapter.
func (inst *MainAdapter) SetToken(acc *jwt.Access) error {

	inst.innerTryAutoTerm(acc.Token)

	err := acc.Encode(false)
	if err != nil {
		return err
	}

	cache, err := inst.innerGetCache()
	if err != nil {
		return err
	}

	count := 0

	for _, reg := range cache.items {
		ada := reg.Adapter
		err := ada.SetToken(acc)
		if err != nil {
			return err
		}
		count++
	}

	if count < 1 {
		return fmt.Errorf("jwt.MainAdapter: no Adapter handle SetToken()")
	}

	return nil
}

func (inst *MainAdapter) innerTryAutoTerm(t *jwt.Token) error {

	if !inst.AutoTermEnabled {
		return nil
	}

	if t == nil {
		return fmt.Errorf(": token is nil")
	}

	ts1 := t.NotBefore
	ts2 := t.NotAfter

	if ts1 == 0 && ts2 == 0 {

		now := lang.Now()
		age := inst.AutoTermAge

		ts1 = now - 1
		ts2 = now + lang.Time(age)

		t.NotBefore = ts1
		t.NotAfter = ts2
	}

	return nil
}

func (inst *MainAdapter) innerGetCache() (*innerAdapterCache, error) {

	c := inst.cache
	if c != nil {
		return c, nil
	}

	result, err := inst.innerLoadCache()
	if err != nil {
		return nil, err
	}

	inst.cache = result
	c = result

	return result, nil
}

func (inst *MainAdapter) innerLoadCache() (*innerAdapterCache, error) {
	cache := new(innerAdapterCache)
	err := cache.load(inst.ProviderList)
	return cache, err
}

func (inst *MainAdapter) _impl() jwt.Adapter {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

type innerAdapterCache struct {
	items []*jwt.AdapterRegistration
}

// Len implements sort.Interface.
func (inst *innerAdapterCache) Len() int {
	return len(inst.items)
}

// Less implements sort.Interface.
func (inst *innerAdapterCache) Less(i int, j int) bool {
	n1 := inst.innerGetItemOrder(i)
	n2 := inst.innerGetItemOrder(j)
	return (n1 > n2)
}

// Swap implements sort.Interface.
func (inst *innerAdapterCache) Swap(i int, j int) {
	all := inst.items
	all[i], all[j] = all[j], all[i]
}

func (inst *innerAdapterCache) load(plist []jwt.AdapterProvider) error {

	src := plist
	dst := inst.items

	for _, provider := range src {
		if provider == nil {
			continue
		}
		reg := provider.Registration()
		inst.innerTryLoad(reg)
		if inst.innerIsRegReady(reg) {
			dst = append(dst, reg)
		}
	}

	inst.items = dst
	inst.innerSortItems()

	return nil
}

func (inst *innerAdapterCache) innerTryLoad(reg *jwt.AdapterRegistration) error {

	if reg == nil {
		return fmt.Errorf("jwt.adapter: AdapterRegistration is nil")
	}

	ldr := reg.Loader
	ada := reg.Adapter

	if ada != nil {
		return nil
	}
	if ldr == nil {
		return fmt.Errorf("jwt.main-adapter: adapter loader is nil")
	}

	return ldr.Load(reg)
}

func (inst *innerAdapterCache) innerIsRegReady(reg *jwt.AdapterRegistration) bool {

	if reg == nil {
		return false
	}

	if !reg.Enabled {
		return false
	}

	if reg.Adapter == nil {
		return false
	}

	return true
}

func (inst *innerAdapterCache) innerGetItemOrder(index int) int {
	it := inst.items[index]
	if it == nil {
		return 0
	}
	return it.Priority
}

func (inst *innerAdapterCache) innerSortItems() {
	sort.Sort(inst)
}

////////////////////////////////////////////////////////////////////////////////
