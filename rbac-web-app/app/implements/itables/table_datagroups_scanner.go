package itables

import (
	"context"
	"fmt"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/libgorm"
	"github.com/starter-go/rbac"
	"github.com/starter-go/v0/rbac-web-app/app/classes/tables"
	"github.com/starter-go/vlog"
)

type TableNamer interface {
	TableName() string
}

type innerTableGroupScanner struct {
	tabserv tables.Service
	ctx     context.Context
	results []*tables.DTO
}

func (inst *innerTableGroupScanner) init(ctx context.Context, ts tables.Service) {
	inst.tabserv = ts
	inst.ctx = ctx
}

func (inst *innerTableGroupScanner) scan(src []libgorm.GroupRegistry) error {

	for _, gr1 := range src {
		tmp := gr1.Groups()
		for _, gr2 := range tmp {
			group := gr2.Group
			ptlist := group.Prototypes()
			for _, pt := range ptlist {
				inst.innerOnTable(gr2, pt)
			}
		}

	}

	return nil
}

func (inst *innerTableGroupScanner) innerOnTable(reg *libgorm.GroupRegistration, pt any) error {

	namer, ok := pt.(TableNamer)
	if !ok {
		return fmt.Errorf("entity.Prototypes no method of TableName() ")
	}

	tabname := namer.TableName()
	gname := reg.Alias
	guri := reg.URI

	const f = "[group name:'%s' uri:'%s' ].[table name:'%s']"
	vlog.Info(f, gname, guri, tabname)

	o1 := new(tables.DTO)
	ser := inst.tabserv
	ctx := inst.ctx

	// fields:  table_name, group_name, group_uri, table_uri

	o1.Name = rbac.TableName(tabname)
	o1.TableURI = ""
	o1.GroupName = reg.Alias
	o1.GroupURI = lang.URI(reg.URI)

	o2, err := ser.Insert(ctx, o1)
	if err != nil {
		vlog.Warn("%s", err.Error())
		o2 = o1
	}

	inst.results = append(inst.results, o2)

	return nil
}

func (inst *innerTableGroupScanner) getResultList() []*tables.DTO {
	return inst.results
}
