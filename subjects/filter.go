package subjects

type ReadFilter interface {
	Read(c *Context, next ReadFilterChain) error
}

type WriteFilter interface {
	Write(c *Context, next WriteFilterChain) error
}
