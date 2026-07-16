package entity

import (
	"fmt"

	"github.com/starter-go/rbac"
	"gorm.io/gorm"
)

type Finder struct {
	db   *gorm.DB
	page *rbac.Pagination

	all bool

	list any
	item any
	want any

	// results
	count int64
	res   *gorm.DB
}

func (inst *Finder) SetDB(db *gorm.DB) *Finder {
	inst.db = db
	return inst
}

func (inst *Finder) SetList(l any) *Finder {
	inst.list = l
	return inst
}

func (inst *Finder) SetItem(it any) *Finder {
	inst.item = it
	return inst
}

func (inst *Finder) SetPagination(p *rbac.Pagination) *Finder {
	inst.page = p
	return inst
}

func (inst *Finder) SetWant(w any) *Finder {
	inst.want = w
	return inst
}

func (inst *Finder) SetAll(all bool) *Finder {
	inst.all = all
	return inst
}

func (inst *Finder) Find() error {

	db := inst.db
	list := inst.list
	item := inst.item
	page := inst.page
	want := inst.want

	if db == nil {
		return fmt.Errorf("entity.Finder.Find() : param 'db' is nil")
	}
	if item == nil {
		return fmt.Errorf("entity.Finder.Find() : param 'item' is nil")
	}
	if list == nil {
		return fmt.Errorf("entity.Finder.Find() : param 'list' is nil")
	}

	// model
	db = db.Model(item)

	// where (want)
	if want != nil {
		db = db.Where(want)
	}

	// count
	var count int64
	db.Count(&count)
	inst.count = count

	// page
	if page == nil {
		page = new(rbac.Pagination)
		page.Size = 10
		inst.page = page
	}
	db = db.Limit(int(page.Limit()))
	db = db.Offset(int(page.Offset()))
	page.Total = count

	//find
	res := db.Find(list)
	inst.res = res
	return res.Error
}
