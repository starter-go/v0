package statestores

import (
	"context"
	"net/http"
)

////////////////////////////////////////////////////////////////////////////////

type Method string

const (
	MethodGet    Method = http.MethodGet
	MethodPut    Method = http.MethodPut
	MethodPost   Method = http.MethodPost
	MethodDelete Method = http.MethodDelete
)

////////////////////////////////////////////////////////////////////////////////

type FilterPriority int

const (
	FilterPriorityBase FilterPriority = iota
	FilterPriorityMin

	FilterPriorityDB
	FilterPriorityCache
	FilterPriorityToken

	FilterPriorityMax
)

////////////////////////////////////////////////////////////////////////////////

type Store interface {
	Handle(sc *Context) error
}

type Service interface {
	GetStore(cc context.Context) (Store, error)
}

////////////////////////////////////////////////////////////////////////////////
// EOF
