package main

import (
	"os"

	"github.com/starter-go/starter"
	"github.com/starter-go/v0/security/modules/security"
)

func main() {

	a := os.Args
	m := security.Module()
	i := starter.Init(a)

	i.MainModule(m)

	i.WithPanic(true).Run()
}
