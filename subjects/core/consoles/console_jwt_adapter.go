package consoles

import (
	"context"
	"fmt"

	"github.com/starter-go/base/context2"
	"github.com/starter-go/v0/libjwt/api/jwt"
)

type ConsoleJwtAdapter struct {

	//starter:component

	_as func(jwt.AdapterProvider) //starter:as(".")

	MyName     string //starter:inject("${jwt-adapter.console.name}")
	MyEnabled  bool   //starter:inject("${jwt-adapter.console.enabled}")
	MyPriority int    //starter:inject("${jwt-adapter.console.priority}")

}

// Registration implements jwt.AdapterProvider.
func (inst *ConsoleJwtAdapter) Registration() *jwt.AdapterRegistration {
	reg := &jwt.AdapterRegistration{
		Enabled:  inst.MyEnabled,
		Name:     inst.MyName,
		Priority: inst.MyPriority,
		Loader:   inst,
		Provider: inst,
	}
	return reg
}

// Load implements jwt.AdapterLoader.
func (inst *ConsoleJwtAdapter) Load(target *jwt.AdapterRegistration) error {
	target.Adapter = inst
	target.Loader = inst
	target.Provider = inst
	return nil
}

// GetToken implements jwt.Adapter.
func (inst *ConsoleJwtAdapter) GetToken(acc *jwt.Access) error {

	ctx := acc.Context
	txt, err := inst.innerGet(ctx)
	if err != nil {
		return err
	}

	acc.Text = txt
	err = acc.Decode(false)

	return err
}

// SetToken implements jwt.Adapter.
func (inst *ConsoleJwtAdapter) SetToken(acc *jwt.Access) error {

	err := acc.Encode(false)
	if err != nil {
		return err
	}

	txt := acc.Text
	ctx := acc.Context
	err = inst.innerSet(ctx, txt)

	return err
}

func (inst *ConsoleJwtAdapter) innerGetKey() string {
	return "github.com/starter-go/v0/subjects/core/consoles/ConsoleJwtAdapter#jwt-text-binding"
}

func (inst *ConsoleJwtAdapter) innerGet(c context.Context) (jwt.Text, error) {

	key := inst.innerGetKey()
	value := c.Value(key)
	if value == nil {
		return "", fmt.Errorf("ConsoleJwtAdapter.get: value is nil")
	}

	txt, ok := value.(jwt.Text)
	if !ok {
		return "", fmt.Errorf("ConsoleJwtAdapter.get: value is not a jwt.Text")
	}

	return txt, nil
}

func (inst *ConsoleJwtAdapter) innerSet(c context.Context, txt jwt.Text) error {

	key := inst.innerGetKey()
	valist, err := context2.GetValues(c)
	if err != nil {
		return err
	}
	valist.SetValue(key, txt)
	return nil
}

func (inst *ConsoleJwtAdapter) _impl() (jwt.AdapterProvider, jwt.AdapterLoader, jwt.Adapter) {
	return inst, inst, inst
}
