package users

import (
	"context"
)

type Service interface {

	// query

	Find(ctx context.Context, id ID) (*DTO, error)

	Query(ctx context.Context, q *Query) ([]*DTO, error)

	// edit

	Insert(cc context.Context, item *DTO) (*DTO, error)
}
