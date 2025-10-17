package main4security

import "github.com/starter-go/application"

//starter:configen(version="4")

func ExportComponents(cr application.ComponentRegistry) error {
	return registerComponents(cr)
}
