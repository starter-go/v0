package core

import (
	"fmt"

	"github.com/starter-go/v0/htttest"
	"github.com/starter-go/v0/htttest/api/h3t"
)

////////////////////////////////////////////////////////////////////////////////

type CodecManagerImpl struct {

	//starter:component

	_as func(htttest.CodecManager) //starter:as("#")

	RawProviderList []htttest.CodecProvider //starter:inject(".")

	cache *innerCodecManagerCache
}

// FindCodec implements h3t.CodecManager.
func (inst *CodecManagerImpl) FindCodec(ctype h3t.ContentType) (h3t.Codec, error) {
	reg, err := inst.innerFindReg(MapTagCodec, ctype)
	if err != nil {
		return nil, err
	}
	codec := reg.Codec
	if codec == nil {
		return nil, fmt.Errorf("CodecManagerImpl: the result codec is nil")
	}
	return codec, nil
}

// FindDecoder implements h3t.CodecManager.
func (inst *CodecManagerImpl) FindDecoder(ctype h3t.ContentType) (h3t.Decoder, error) {
	reg, err := inst.innerFindReg(MapTagDecoder, ctype)
	if err != nil {
		return nil, err
	}
	dec := reg.Decoder
	if dec == nil {
		return nil, fmt.Errorf("CodecManagerImpl: the result decoder is nil")
	}
	return dec, nil
}

// FindEncoder implements h3t.CodecManager.
func (inst *CodecManagerImpl) FindEncoder(ctype h3t.ContentType) (h3t.Encoder, error) {
	reg, err := inst.innerFindReg(MapTagEncoder, ctype)
	if err != nil {
		return nil, err
	}
	enc := reg.Encoder
	if enc == nil {
		return nil, fmt.Errorf("CodecManagerImpl: the result encoder is nil")
	}
	return enc, nil
}

// GetDefaultCodec implements h3t.CodecManager.
func (inst *CodecManagerImpl) GetDefaultCodec() h3t.Codec {
	codec := new(innerNopCodec)
	return codec
}

func (inst *CodecManagerImpl) innerFindReg(tag innerCodecMapTag, ctype h3t.ContentType) (*h3t.CodecRegistration, error) {
	cache, err := inst.innerGetCache()
	if err != nil {
		return nil, err
	}
	return cache.findReg(tag, ctype)
}

func (inst *CodecManagerImpl) innerGetCache() (*innerCodecManagerCache, error) {
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

func (inst *CodecManagerImpl) innerLoadCache() (*innerCodecManagerCache, error) {
	c := new(innerCodecManagerCache)
	c = c.init()
	err := c.load(inst.RawProviderList)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (inst *CodecManagerImpl) _impl() htttest.CodecManager {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

type innerCodecMapKey string

type innerCodecMapTag string

const (
	MapTagCodec   innerCodecMapTag = "codec"
	MapTagEncoder innerCodecMapTag = "encoder"
	MapTagDecoder innerCodecMapTag = "decoder"
)

////////////////////////////////////////////////////////////////////////////////

type innerNopCodec struct {
}

// Decode implements h3t.Codec.
func (inst *innerNopCodec) Decode(c *h3t.Content) error {
	c.Data = ""
	return nil
}

// Encode implements h3t.Codec.
func (inst *innerNopCodec) Encode(c *h3t.Content) error {
	c.Encoded = []byte{}
	return nil
}

func (inst *innerNopCodec) _impl() h3t.Codec {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

type innerCodecManagerCache struct {
	table map[innerCodecMapKey]*h3t.CodecRegistration
}

func (inst *innerCodecManagerCache) init() *innerCodecManagerCache {
	if inst == nil {
		inst = new(innerCodecManagerCache)
	}
	if inst.table == nil {
		inst.table = make(map[innerCodecMapKey]*h3t.CodecRegistration)
	}
	return inst
}

func (inst *innerCodecManagerCache) keyFor(tag innerCodecMapTag, ctype h3t.ContentType) innerCodecMapKey {
	str1 := string(tag) + ":"
	str2 := ctype.Pure().Normalize().String()
	return innerCodecMapKey(str1 + str2)
}

func (inst *innerCodecManagerCache) load(plist []htttest.CodecProvider) error {
	table := inst.init().table
	for _, provider := range plist {
		if provider == nil {
			continue
		}
		tmp := provider.Registrations()
		for _, reg := range tmp {
			inst.innerOnLoadReg(reg, table)
		}
	}
	inst.table = table
	return nil
}

func (inst *innerCodecManagerCache) innerOnLoadReg(reg *h3t.CodecRegistration, dst map[innerCodecMapKey]*h3t.CodecRegistration) error {

	if reg == nil || dst == nil {
		return nil
	}

	if !reg.Enabled {
		return nil
	}

	if reg.Codec != nil {
		key := inst.keyFor(MapTagCodec, reg.Type)
		dst[key] = reg
	}

	if reg.Encoder != nil {
		key := inst.keyFor(MapTagEncoder, reg.Type)
		dst[key] = reg
	}

	if reg.Decoder != nil {
		key := inst.keyFor(MapTagDecoder, reg.Type)
		dst[key] = reg
	}

	return nil
}

func (inst *innerCodecManagerCache) findReg(tag innerCodecMapTag, ctype h3t.ContentType) (*h3t.CodecRegistration, error) {
	key := inst.keyFor(tag, ctype)
	reg := inst.table[key]
	if reg == nil {
		return nil, fmt.Errorf("[Error this:innerCodecManagerCache key:'%s' msg:'no codec with key']", key)
	}
	return reg, nil
}

////////////////////////////////////////////////////////////////////////////////
// EOF
