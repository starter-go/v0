package main

import (
	"os"

	"github.com/starter-go/units"
	"github.com/starter-go/v0/htttest/modules/htttest"
)

func main() {

	a := os.Args
	m := htttest.Module()

	ctx := &units.Context{
		Arguments: a,
		Module:    m,
		UsePanic:  true,
	}

	units.Run(ctx)
}
