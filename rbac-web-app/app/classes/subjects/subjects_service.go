package subjects

import "context"

type Service interface {
	InitContext(c1 context.Context, c2 *Context) error

	InitSubject(c1 context.Context, sub Subject) error
}
