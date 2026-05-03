package sessions

import "context"

type Service interface {

	// fetch

	Find(c context.Context, id ID) (*DTO, error)

	Query(c context.Context, q *Query) ([]*DTO, error)

	// edit

	Insert(c context.Context, item *DTO) (*DTO, error)

	Update(c context.Context, id ID, item *DTO) (*DTO, error)

	Remove(c context.Context, id ID) error
}
