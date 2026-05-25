package h3t

import "net/url"

type Resolver interface {
	Resolve(c *Transaction) (*url.URL, error)
}
