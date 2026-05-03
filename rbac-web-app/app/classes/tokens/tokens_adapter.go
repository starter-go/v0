package tokens

import (
	"context"
)

type Adapter interface {
	Put(ctx context.Context, t JWT) error

	Get(ctx context.Context) (JWT, error)
}
