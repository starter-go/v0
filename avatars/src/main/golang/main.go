package main

import (
	"os"

	"github.com/starter-go/starter"
	"github.com/starter-go/v0/avatars/modules/avatars"
)

func main() {

	a := os.Args
	m := avatars.Module()
	i := starter.Init(a)

	i.MainModule(m)

	i.WithPanic(true).Run()
}
