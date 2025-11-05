package webcontexts

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

const (
	theAccessAttrKey = "github.com/starter-go/v0/rbac-web-app/app/classes/webcontexts#Access-binding"
)

////////////////////////////////////////////////////////////////////////////////
// 接口

type AttrSetter interface {
	SetAttr(name string, value any)
}

type AttrGetter interface {
	GetAttr(name string) any

	GetAttrRequired(name string) (any, error)
}

type Access interface {
	AttrSetter

	AttrGetter
}

////////////////////////////////////////////////////////////////////////////////
// 公开函数

func GetAccess(cc context.Context) (Access, error) {
	const key = theAccessAttrKey
	value := cc.Value(key)
	if value == nil {
		return nil, fmt.Errorf("need invoke webcontexts.SetupAccessForGinContext() first")
	}
	acc := value.(Access)
	return acc, nil
}

func SetupAccessForGinContext(c *gin.Context) {

	const key = theAccessAttrKey
	older := c.Value(key)
	if older != nil {
		acc, ok := older.(Access)
		if ok && (acc != nil) {
			return
		}
	}

	// do setup
	holder := NewHolder(c)
	c.Set(key, holder)
}

////////////////////////////////////////////////////////////////////////////////
// EOF
