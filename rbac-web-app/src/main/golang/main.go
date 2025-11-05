package main

import (
	"os"

	"github.com/starter-go/starter"
	"github.com/starter-go/v0/rbac-web-app/modules/rbacwebapp"
)

func main() {

	a := os.Args
	m := rbacwebapp.Module()
	i := starter.Init(a)

	i.MainModule(m)

	i.WithPanic(true).Run()
}
