package subjects

type ReadFilter interface {
	Read(c *IOC, next ReadFilterChain) error
}

type WriteFilter interface {
	Write(c *IOC, next WriteFilterChain) error
}
