package libkeys

import "github.com/starter-go/v0/keys"

type Loader struct {

	//starter:component

	_as func(keys.ServiceLoader) //starter:as("#")

	Ser keys.Service //starter:inject("#")

}

// Load implements keys.ServiceLoader.
func (inst *Loader) Load() keys.Service {
	return inst.Ser
}

func (inst *Loader) _impl() keys.ServiceLoader {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

// func NewLoader() *Loader {
// 	return new(Loader)
// }
