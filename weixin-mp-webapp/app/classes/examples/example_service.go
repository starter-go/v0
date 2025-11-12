package examples

import (
	"context"

	"github.com/starter-go/v0/weixin-mp-webapp/app/data/dxo"
	"github.com/starter-go/v0/weixin-mp-webapp/app/web/dto"
)

type Service interface {

	// query

	Find(cc context.Context, id dxo.ExampleID) (*dto.Example, error)

	Query(cc context.Context, q *Query) ([]*dto.Example, error)

	// edit

	Insert(cc context.Context, item *dto.Example) (*dto.Example, error)

	Update(cc context.Context, id dxo.ExampleID, item *dto.Example) (*dto.Example, error)

	Remove(cc context.Context, id dxo.ExampleID) error
}
