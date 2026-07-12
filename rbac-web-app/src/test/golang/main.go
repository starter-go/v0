package main

import (
	"os"

	"github.com/starter-go/units"
	"github.com/starter-go/v0/rbac-web-app/modules/rbacwebapp"
)

func main() {

	a := os.Args
	m := rbacwebapp.ModuleForTest()

	c := &units.Context{
		Arguments: a,
		Module:    m,
		UsePanic:  true,
	}

	units.Run(c)
}
