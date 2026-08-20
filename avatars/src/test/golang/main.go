package main

import (
	"os"

	"github.com/starter-go/units"
	"github.com/starter-go/v0/avatars/modules/avatars"
)

func main() {

	a := os.Args
	m := avatars.ModuleForTest()

	c := &units.Context{
		Arguments: a,
		Module:    m,
		UsePanic:  true,
	}

	units.Run(c)
}
