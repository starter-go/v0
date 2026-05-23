package ikeyloader

import (
	"fmt"
	"sort"

	"github.com/starter-go/v0/libjwt/api/jwt"
)

type MainKeyLoader struct {

	//starter:component

	_as func(jwt.KeyLoader) //starter:as("#")

	ProviderList []jwt.KeyLoaderProvider //starter:inject(".")

	LengthInBytes int //starter:inject("${jwt-key-loader.main.length}")

	cache *innerKeyLoaderCache
}

// Load implements jwt.KeyLoader.
func (inst *MainKeyLoader) Load() (*jwt.Key, error) {

	cache, err := inst.innerGetCache()
	if err != nil {
		return nil, err
	}

	loader, err := cache.getFirstLoader()
	if err != nil {
		return nil, err
	}

	key, err := loader.Load()
	if err != nil {
		return nil, err
	}

	key = inst.innerNormalizeKeyLength(key)

	return key, err
}

// Registration implements jwt.KeyLoader.
func (inst *MainKeyLoader) Registration() *jwt.KeyLoaderRegistration {
	return new(jwt.KeyLoaderRegistration)
}

func (inst *MainKeyLoader) innerGetCache() (*innerKeyLoaderCache, error) {
	c := inst.cache
	if c == nil {
		c2, err := inst.innerLoadCache()
		if err != nil {
			return nil, err
		}
		c = c2
		inst.cache = c2
	}
	return c, nil
}

func (inst *MainKeyLoader) innerLoadCache() (*innerKeyLoaderCache, error) {
	cache := new(innerKeyLoaderCache)
	err := cache.init(inst.ProviderList)
	return cache, err
}

func (inst *MainKeyLoader) innerNormalizeKeyLength(key1 *jwt.Key) *jwt.Key {

	const (
		theDefaultLen = 16
	)

	raw := key1.Bytes()
	lengthWant := inst.LengthInBytes
	lengthHave := len(raw)

	if lengthWant < 1 {
		lengthWant = theDefaultLen
	}

	if lengthHave == lengthWant {
		return key1
	}

	dst := make([]byte, lengthWant)
	copy(dst, raw)
	return jwt.NewKey(dst)
}

func (inst *MainKeyLoader) _impl() (jwt.KeyLoader, jwt.KeyLoaderProvider) {
	return inst, inst
}

////////////////////////////////////////////////////////////////////////////////

type innerKeyLoaderCache struct {
	items []*jwt.KeyLoaderRegistration
	first *jwt.KeyLoaderRegistration
}

// Len implements sort.Interface.
func (inst *innerKeyLoaderCache) Len() int {
	return len(inst.items)
}

// Less implements sort.Interface.
func (inst *innerKeyLoaderCache) Less(i int, j int) bool {
	n1 := inst.innerGetItemOrderNum(i)
	n2 := inst.innerGetItemOrderNum(j)
	return (n1 > n2)
}

// Swap implements sort.Interface.
func (inst *innerKeyLoaderCache) Swap(i int, j int) {
	all := inst.items
	all[i], all[j] = all[j], all[i]
}

func (inst *innerKeyLoaderCache) init(plist []jwt.KeyLoaderProvider) error {

	dst := make([]*jwt.KeyLoaderRegistration, 0)

	for _, provider := range plist {
		if provider == nil {
			continue
		}
		info := provider.Registration()
		if inst.innerIsItemReady(info) {
			dst = append(dst, info)
		}
	}

	inst.items = dst
	inst.innerSortItems()

	_, err := inst.getFirstLoader()
	return err
}

func (inst *innerKeyLoaderCache) innerGetItemOrderNum(index int) int {
	it := inst.items[index]
	if it == nil {
		return 0
	}
	return it.Priority
}

func (inst *innerKeyLoaderCache) innerSortItems() {
	sort.Sort(inst)
}

func (inst *innerKeyLoaderCache) innerIsItemReady(item *jwt.KeyLoaderRegistration) bool {

	if item == nil {
		return false
	}

	if !item.Enabled {
		return false
	}

	if item.Loader == nil {
		return false
	}

	return true
}

func (inst *innerKeyLoaderCache) getFirstLoader() (jwt.KeyLoader, error) {

	reg1, err := inst.getFirstReg()
	if err != nil {
		return nil, err
	}

	ldr := reg1.Loader
	if ldr == nil {
		return nil, fmt.Errorf("jwt.key-loader: the first key-loader is nil")
	}

	return ldr, nil
}

func (inst *innerKeyLoaderCache) getFirstReg() (*jwt.KeyLoaderRegistration, error) {

	first := inst.first
	if first != nil {
		return first, nil
	}

	// do load
	src := inst.items
	for _, it := range src {
		if it == nil {
			continue
		}
		inst.first = it
		return it, nil
	}

	return nil, fmt.Errorf("jwt.key-loader: no key-loader (list is empty)")
}
