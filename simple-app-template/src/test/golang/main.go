package main

import (
	"os"

	"github.com/starter-go/units"
	simpleapptemplate "github.com/starter-go/v0/simple-app-template"
)

func main() {

	a := os.Args
	m := simpleapptemplate.ModuleForTest()

	c := &units.Context{
		Arguments: a,
		Module:    m,
		UsePanic:  true,
	}

	units.Run(c)
}
