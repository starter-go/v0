package roles

import "context"

type Service interface {
	Find(c context.Context, id ID) (*DTO, error)

	Query(c context.Context, q *Query) ([]*DTO, error)

	Insert(c context.Context, item *DTO) (*DTO, error)
}
