package main

import (
	"os"

	"github.com/starter-go/starter"
	"github.com/starter-go/v0/httest"
)

func main() {

	a := os.Args
	m := httest.Module()
	i := starter.Init(a)

	i.MainModule(m)

	i.WithPanic(true).Run()
}
