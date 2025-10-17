package subjects

import "context"

type Service interface {
	GetContext(cc context.Context) (*Context, error)

	GetSubject(cc context.Context) (*Subject, error)
}
