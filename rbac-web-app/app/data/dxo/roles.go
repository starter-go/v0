package dxo

import (
	"strings"

	"github.com/starter-go/rbac"
)

type RoleName = rbac.RoleName

type RoleNameList = rbac.RoleNameList

////////////////////////////////////////////////////////////////////////////////

type RoleNameSlice []rbac.RoleName

func (sl RoleNameSlice) Format() RoleNameList {

	const sep = ','
	builder := new(strings.Builder)

	for _, it1 := range sl {
		it2 := it1.String()
		it2 = strings.TrimSpace(it2)
		if it2 == "" {
			continue
		}
		if builder.Len() > 0 {
			builder.WriteRune(sep)
		}
		builder.WriteString(it2)
	}

	str := builder.String()
	return RoleNameList(str)
}

////////////////////////////////////////////////////////////////////////////////
