package main

import (
	"os"

	"github.com/starter-go/starter"
	webapptemplate "github.com/starter-go/v0/web-app-template"
)

func main() {

	a := os.Args
	m := webapptemplate.ModuleForTest()
	i := starter.Init(a)

	i.MainModule(m)

	i.WithPanic(true).Run()
}
