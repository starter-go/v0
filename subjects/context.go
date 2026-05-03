package subjects

import (
	"context"
)

type Context struct {
	CC context.Context

	Method Method

	Cache *Cache // for reader

	Buffer *Buffer // for writer

	Facade Subject

	Reader ReadFilterChain

	Writer WriteFilterChain

	ChainHolder FilterChainHolder
}

func (inst *Context) Clone() *Context {

	c2 := new(Context)
	*c2 = *inst
	c2.Facade = NewSubject(c2)
	return c2
}
