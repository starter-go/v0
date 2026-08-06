package main

import (
	"os"

	"github.com/starter-go/units"
	"github.com/starter-go/v0/liberr/modules/liberr"
)

func main() {

	a := os.Args
	m := liberr.ModuleForTest()

	c := &units.Context{
		Arguments: a,
		Module:    m,
		UsePanic:  true,
	}

	units.Run(c)
}
