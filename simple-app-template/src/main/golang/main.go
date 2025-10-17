package main

import (
	"os"

	"github.com/starter-go/starter"
	simpleapptemplate "github.com/starter-go/v0/simple-app-template"
)

func main() {

	a := os.Args
	m := simpleapptemplate.Module()
	i := starter.Init(a)

	i.MainModule(m)

	i.WithPanic(true).Run()
}
