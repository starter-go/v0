package main

import (
	"os"

	"github.com/starter-go/starter"
	"github.com/starter-go/v0/libdao/modules/libdao"
)

func main() {

	a := os.Args
	m := libdao.Module()
	i := starter.Init(a)

	i.MainModule(m)

	i.WithPanic(true).Run()
}
