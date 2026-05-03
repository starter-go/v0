package main

import (
	"os"

	"github.com/starter-go/starter"
	"github.com/starter-go/v0/subjects/modules/subjects"
)

func main() {

	a := os.Args
	m := subjects.ModuleForTest()
	i := starter.Init(a)

	i.MainModule(m)

	i.WithPanic(true).Run()
}
