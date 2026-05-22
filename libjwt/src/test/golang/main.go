package main

import (
	"os"

	"github.com/starter-go/units"
	"github.com/starter-go/v0/libjwt"
)

func main() {

	a := os.Args
	m := libjwt.ModuleForTest()

	ctx := &units.Context{
		Arguments: a,
		Module:    m,
		UsePanic:  true,
		Selector:  "",
	}

	units.Run(ctx)

}
