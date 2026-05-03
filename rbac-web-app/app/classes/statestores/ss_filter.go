package statestores

// filters

type Filter interface {
	Handle(sc *Context, next FilterChain) error
}
