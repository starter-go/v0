package main

import (
	"os"

	"github.com/starter-go/units"
	webapptemplate "github.com/starter-go/v0/web-app-template"
)

func main() {

	a := os.Args
	m := webapptemplate.ModuleForTest()

	c := &units.Context{
		Arguments: a,
		Module:    m,
		UsePanic:  true,
	}

	units.Run(c)
}
