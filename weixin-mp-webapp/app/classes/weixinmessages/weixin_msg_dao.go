package weixinmessages

import (
	"github.com/starter-go/v0/weixin-mp-webapp/app/data/dxo"
	"github.com/starter-go/v0/weixin-mp-webapp/app/data/entity"
	"gorm.io/gorm"
)

type DAO interface {

	// query

	Find(db *gorm.DB, id dxo.ExampleID) (*entity.Example, error)

	Query(db *gorm.DB, q *Query) ([]*entity.Example, error)

	// edit

	Insert(db *gorm.DB, item *entity.Example) (*entity.Example, error)

	Update(db *gorm.DB, id dxo.ExampleID, fn func(item *entity.Example) error) (*entity.Example, error)

	Remove(db *gorm.DB, id dxo.ExampleID) error
}
