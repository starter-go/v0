package main

import (
	"os"

	"github.com/starter-go/units"
	"github.com/starter-go/v0/keys/modules/keys"
)

func main() {

	a := os.Args
	m := keys.ModuleTest()
	c := new(units.Context)

	c.Arguments = a
	c.Module = m
	c.UsePanic = true

	units.Run(c)

}
