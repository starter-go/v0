package main

import (
	"os"

	"github.com/starter-go/starter"
	"github.com/starter-go/v0/liberr/modules/liberr"
)

func main() {

	a := os.Args
	m := liberr.Module()
	i := starter.Init(a)

	i.MainModule(m)

	i.WithPanic(true).Run()
}
