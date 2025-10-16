package main

import (
	"os"

	"github.com/starter-go/starter"
	securitywebapp "github.com/starter-go/v0/security-web-app"
)

func main() {

	a := os.Args
	m := securitywebapp.ModuleForTest()
	i := starter.Init(a)

	i.MainModule(m)

	i.WithPanic(true).Run()
}
