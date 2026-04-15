package main

import (
	"net/http"

	"github.com/starter-go/v0/httest"
)

func main() {

	r := new(httest.Runner)
	r.SetProperty("foo11", "bar22")
	r.SetHeader("x-timestamp", "12345678")
	r.SetMethod(http.MethodGet).SetURL("/a/b/c/test")

	r.Run()

}
