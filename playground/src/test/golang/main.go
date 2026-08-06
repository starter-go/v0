package main

import (
	"os"

	"github.com/starter-go/units"
	"github.com/starter-go/v0/playground"
)

func main() {

	a := os.Args
	m := playground.ModuleForTest()

	c := &units.Context{
		Arguments: a,
		Module:    m,
		UsePanic:  true,
	}

	units.Run(c)
}
