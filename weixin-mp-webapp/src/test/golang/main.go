package main

import (
	"os"

	"github.com/starter-go/starter"
	"github.com/starter-go/v0/weixin-mp-webapp/modules/wxmp"
)

func main() {

	a := os.Args
	m := wxmp.ModuleForTest()
	i := starter.Init(a)

	i.MainModule(m)

	i.WithPanic(true).Run()
}
