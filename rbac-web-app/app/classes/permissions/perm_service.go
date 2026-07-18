package permissions

import "context"

type Service interface {
	Find(cc context.Context, id ID) (*DTO, error)

	Query(cc context.Context, q *Query) ([]*DTO, error)

	Insert(cc context.Context, item *DTO) (*DTO, error)

	Setup(cc context.Context) ([]*DTO, error)

	Reload(cc context.Context) error
}
