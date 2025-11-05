package webcontexts

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Holder struct {
	context *gin.Context
}

func (inst *Holder) _impl() Access {
	return inst
}

func (inst *Holder) Get() *gin.Context {
	return inst.context
}

func (inst *Holder) GetAttr(name string) any {
	return inst.context.Value(name)
}

func (inst *Holder) GetAttrRequired(name string) (any, error) {
	value := inst.GetAttr(name)
	if value == nil {
		return nil, fmt.Errorf("no (attr) value with name [%s] in the gin.Context", name)
	}
	return value, nil
}

func (inst *Holder) SetAttr(name string, value any) {
	if name == "" || value == nil {
		return
	}
	inst.context.Set(name, value)
}

func GetHolder(cc context.Context) (*Holder, error) {
	acc, err := GetAccess(cc)
	if err != nil {
		return nil, err
	}
	h := acc.(*Holder)
	return h, nil
}

func NewHolder(c *gin.Context) *Holder {
	return &Holder{
		context: c,
	}
}
