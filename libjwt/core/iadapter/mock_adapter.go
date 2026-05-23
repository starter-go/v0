package iadapter

import "github.com/starter-go/v0/libjwt/api/jwt"

////////////////////////////////////////////////////////////////////////////////

type MockAdapter struct {

	//starter:component

	_as func(jwt.AdapterProvider) //starter:as(".")

	MyName     string //starter:inject("${jwt-adapter.mock.name}")
	MyPriority int    //starter:inject("${jwt-adapter.mock.priority}")
	MyEnabled  bool   //starter:inject("${jwt-adapter.mock.enabled}")

	tmp jwt.Text
}

// Registration implements jwt.AdapterProvider.
func (inst *MockAdapter) Registration() *jwt.AdapterRegistration {

	reg := &jwt.AdapterRegistration{
		Enabled:  inst.MyEnabled,
		Priority: inst.MyPriority,
		Name:     inst.MyName,
		Loader:   inst,
	}

	return reg
}

// Load implements jwt.AdapterLoader.
func (inst *MockAdapter) Load(target *jwt.AdapterRegistration) error {
	target.Loader = inst
	target.Adapter = inst
	return nil
}

// GetToken implements jwt.Adapter.
func (inst *MockAdapter) GetToken(acc *jwt.Access) error {

	txt := inst.tmp
	dec := acc.Service.GetDecoder()

	token, err := dec.Decode(txt)
	if err != nil {
		return err
	}

	acc.Token = token
	acc.Text = txt

	return nil
}

// SetToken implements jwt.Adapter.
func (inst *MockAdapter) SetToken(acc *jwt.Access) error {

	token := acc.Token
	enc := acc.Service.GetEncoder()

	txt, err := enc.Encode(token)
	if err != nil {
		return err
	}

	acc.Text = txt
	inst.tmp = txt
	return nil
}

func (inst *MockAdapter) _impl() (jwt.AdapterProvider, jwt.AdapterLoader, jwt.Adapter) {
	return inst, inst, inst
}

////////////////////////////////////////////////////////////////////////////////
