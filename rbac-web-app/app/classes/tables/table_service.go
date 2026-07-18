package tables

import "context"

// type Service = rbac.TableService

type Service interface {
	Find(cc context.Context, id ID) (*DTO, error)

	Query(cc context.Context, q *Query) ([]*DTO, error)

	Insert(cc context.Context, item *DTO) (*DTO, error)

	Setup(cc context.Context) ([]*DTO, error)

	ScanCurrentRuntime(cc context.Context) ([]*DTO, error)
}
