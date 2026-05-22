package subjects

import "fmt"

type CacheHolder interface {
	GetCache(c *Context) (*Cache, error)
}

////////////////////////////////////////////////////////////////////////////////

type innerCacheHolderImpl struct {
}

// GetCache implements CacheHolder.
func (inst *innerCacheHolderImpl) GetCache(c *Context) (*Cache, error) {

	if c == nil {
		return nil, fmt.Errorf("innerCacheHolderImpl.GetCache() : param 'context' is nil")
	}

	cache := c.Cache
	if inst.innerIsCacheReady(cache) {
		return cache, nil
	}

	// do load

	loader := c.GetCacheLoader(true)
	cache, err := loader.LoadCache(c)
	if err != nil {
		return nil, err
	}
	c.Cache = cache
	return cache, nil
}

func (inst *innerCacheHolderImpl) innerIsCacheReady(cache *Cache) bool {
	if cache == nil {
		return false
	}
	return cache.Ready()
}

func (inst *innerCacheHolderImpl) _impl() CacheHolder {
	return inst
}
