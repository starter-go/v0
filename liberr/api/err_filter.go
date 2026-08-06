package api

import (
	"context"

	"github.com/starter-go/i18n"
)

type Want struct {

	// contexts

	Context context.Context

	LiberrContext *Context

	Service Service

	// params

	Language i18n.Language

	Name Name

	NS Namespace

	Args []any

	// tmp

	Info *HyperErrorInfo

	// results

	Error HyperErrorI18n
}

type Filter interface {
	Error(w *Want, next FilterChain) error
}

type FilterChain interface {
	Error(w *Want) error
}

type FilterRegistration struct {
	Name string

	Enabled bool

	Order int

	Filter Filter
}

type FilterRegistry interface {
	ListRegistrations() []*FilterRegistration
}
