package main

import (
	"os"

	"github.com/starter-go/units"
	"github.com/starter-go/v0/libdao/modules/libdao"
)

func main() {

	a := os.Args
	m := libdao.ModuleForTest()

	c := &units.Context{
		Arguments: a,
		Module:    m,
		UsePanic:  true,
	}

	units.Run(c)
}
