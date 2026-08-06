package main

import (
	"os"

	"github.com/starter-go/starter"
	"github.com/starter-go/v0/rbac-data-group/modules/rbacdg"
)

func main() {

	a := os.Args
	m := rbacdg.Module()
	i := starter.Init(a)

	i.MainModule(m)

	i.WithPanic(true).Run()
}
