package examples

import "context"

type Service interface {
	Find(c context.Context, id ID) (*DTO, error)

	Query(c context.Context, q *Query) ([]*DTO, error)
}
