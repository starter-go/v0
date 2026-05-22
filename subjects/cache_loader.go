package subjects

import (
	"fmt"

	"github.com/starter-go/application/properties"
)

////////////////////////////////////////////////////////////////////////////////

type CacheLoader interface {
	LoadCache(c *Context) (*Cache, error)
}

////////////////////////////////////////////////////////////////////////////////

type innerCacheLoaderImpl struct {
}

// LoadCache implements CacheLoader.
func (inst *innerCacheLoaderImpl) LoadCache(ctx *Context) (*Cache, error) {

	if ctx == nil {
		return nil, fmt.Errorf("context is nil")
	}

	ctx.Cache = nil
	chain := ctx.GetReader(true)
	ioc := ctx.NewIOC(MethodGet)

	err := chain.Read(ioc)
	if err != nil {
		return nil, err
	}

	have := &ioc.Have
	cache := ioc.Context.Cache
	pt := have.Properties
	if cache == nil {
		cache = new(Cache)
	}
	if pt == nil {
		pt = properties.NewTable(nil)
	}
	cache.Properties = pt
	cache.Loaded = true

	return cache, nil
}

func (inst *innerCacheLoaderImpl) _impl() CacheLoader {
	return inst
}

////////////////////////////////////////////////////////////////////////////////
