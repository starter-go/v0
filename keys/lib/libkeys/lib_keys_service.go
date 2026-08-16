package libkeys

import "github.com/starter-go/v0/keys"

type LibKeysService struct {

	//starter:component

	_as func(keys.Service) //starter:as("#")

}

// GetDriverManager implements keys.Service.
func (inst *LibKeysService) GetDriverManager() keys.DriverManager {
	panic("unimplemented")
}

// GetDriverRegistry implements keys.Service.
func (inst *LibKeysService) GetDriverRegistry() keys.DriverRegistry {
	panic("unimplemented")
}

func (inst *LibKeysService) _impl() keys.Service {
	return inst
}
