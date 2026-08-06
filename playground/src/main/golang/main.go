package main

import (
	"os"

	"github.com/starter-go/starter"
	"github.com/starter-go/v0/playground"
)

func main() {

	a := os.Args
	m := playground.Module()
	i := starter.Init(a)

	i.MainModule(m)

	i.WithPanic(true).Run()
}
