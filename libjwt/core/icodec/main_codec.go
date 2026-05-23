package icodec

import (
	"fmt"
	"sort"

	"github.com/starter-go/v0/libjwt/api/jwt"
)

////////////////////////////////////////////////////////////////////////////////

type MainCodec struct {

	//starter:component

	_as func(jwt.CODEC) //starter:as("#")

	ProviderList []jwt.CodecProvider //starter:inject(".")

	cache *innerCodecCache
}

// Decode implements jwt.CODEC.
func (inst *MainCodec) Decode(txt jwt.Text) (*jwt.Token, error) {

	cache, err := inst.innerGetCache()
	if err != nil {
		return nil, err
	}

	codec, err := cache.getFirstCodec()
	if err != nil {
		return nil, err
	}

	return codec.Decode(txt)
}

// Encode implements jwt.CODEC.
func (inst *MainCodec) Encode(token *jwt.Token) (jwt.Text, error) {

	cache, err := inst.innerGetCache()
	if err != nil {
		return "", err
	}

	codec, err := cache.getFirstCodec()
	if err != nil {
		return "", err
	}

	return codec.Encode(token)
}

func (inst *MainCodec) innerGetCache() (*innerCodecCache, error) {
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

func (inst *MainCodec) innerLoadCache() (*innerCodecCache, error) {
	cache := new(innerCodecCache)
	err := cache.load(inst.ProviderList)
	return cache, err
}

func (inst *MainCodec) _impl() jwt.CODEC {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

type innerCodecCache struct {
	items []*jwt.CodecRegistration
	first *jwt.CodecRegistration
}

func (inst *innerCodecCache) load(plist []jwt.CodecProvider) error {

	dst := make([]*jwt.CodecRegistration, 0)

	for _, provider := range plist {
		if provider == nil {
			continue
		}
		reg := provider.Registration()
		if inst.innerIsReady(reg) {
			err := inst.innerLoadItem(reg)
			if err != nil {
				return err
			}
			dst = append(dst, reg)
		}
	}

	inst.items = dst
	inst.innerSortItems()

	first, err := inst.getFirstCodec()
	if err != nil {
		return err
	}
	if first == nil {
		return fmt.Errorf("jwt.main-codec: no any codec")
	}

	return nil
}

func (inst *innerCodecCache) innerLoadItem(item *jwt.CodecRegistration) error {

	if item == nil {
		return fmt.Errorf("jwt-codec-cache: param 'jwt.CodecRegistration' is nil")
	}

	codec := item.Codec

	if codec == nil {
		loader := item.Loader
		if loader == nil {
			return fmt.Errorf("jwt-codec-cache: loader is nil")
		}
		err := loader.Load(item)
		if err != nil {
			return err
		}
		codec = item.Codec
	}

	if codec == nil {
		return fmt.Errorf("jwt-codec-cache: codec is nil")
	}

	return nil
}

func (inst *innerCodecCache) innerIsReady(item *jwt.CodecRegistration) bool {

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

func (inst *innerCodecCache) innerLoadFirst() (*jwt.CodecRegistration, error) {
	src := inst.items
	for _, item := range src {
		if item != nil {
			return item, nil
		}
	}
	return nil, fmt.Errorf("no jwt.codec provider (list of jwt.codec is empty)")
}

func (inst *innerCodecCache) innerSortItems() {
	if inst == nil {
		return
	}
	sort.Sort(inst)
}

func (inst *innerCodecCache) innerGetOrderNumberOf(index int) int {
	item := inst.items[index]
	if item == nil {
		return 0
	}
	return item.Priority
}

func (inst *innerCodecCache) Len() int {
	return len(inst.items)
}
func (inst *innerCodecCache) Less(i1, i2 int) bool {
	n1 := inst.innerGetOrderNumberOf(i1)
	n2 := inst.innerGetOrderNumberOf(i2)
	return (n1 > n2)
}
func (inst *innerCodecCache) Swap(i1, i2 int) {
	list := inst.items
	list[i1], list[i2] = list[i2], list[i1]
}

func (inst *innerCodecCache) getFirstCodec() (jwt.CODEC, error) {

	reg, err := inst.getFirst()
	if err != nil {
		return nil, err
	}

	codec := reg.Codec
	if codec == nil {
		err = fmt.Errorf("jwt.main-codec: codec is nil")
		return nil, err
	}

	return codec, nil
}

func (inst *innerCodecCache) getFirst() (*jwt.CodecRegistration, error) {

	the1st := inst.first
	if the1st != nil {
		return the1st, nil
	}

	// load

	reg, err := inst.innerLoadFirst()
	if err != nil {
		return nil, err
	}

	codec := reg.Codec
	if codec == nil {
		return nil, fmt.Errorf("innerCodecCache: the first codec is nil")
	}

	inst.first = reg
	return reg, nil
}

////////////////////////////////////////////////////////////////////////////////
