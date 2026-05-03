package subjects

type ReadFilterChain interface {
	Read(c *Context) error
}

type WriteFilterChain interface {
	Write(c *Context) error
}

////////////////////////////////////////////////////////////////////////////////

// type FilterChain struct {
// 	Reader ReadFilterChain
// 	Writer WriteFilterChain
// }

type FilterChainLoader interface {
	LoadReaderChain() (ReadFilterChain, error)
	LoadWriterChain() (WriteFilterChain, error)
}

type FilterChainHolder interface {
	GetReaderChain() (ReadFilterChain, error)
	GetWriterChain() (WriteFilterChain, error)
}
