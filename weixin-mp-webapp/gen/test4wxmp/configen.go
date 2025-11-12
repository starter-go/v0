package test4wxmp

import "github.com/starter-go/application"

//starter:configen(version="4")

func ExportComponents(cr application.ComponentRegistry) error {
	return registerComponents(cr)
	// return nil
}
