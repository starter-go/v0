package iadapter

import (
	"fmt"
	"testing"

	"github.com/starter-go/v0/libjwt/api/jwt"
)

////////////////////////////////////////////////////////////////////////////////

func TestMainAdapterSort(t *testing.T) {
	tt := new(innerMainAdapterTester)
	tt.test(t)
}

////////////////////////////////////////////////////////////////////////////////

type innerMainAdapterTester struct {
	priority int
}

// Load implements jwt.AdapterLoader.
func (inst *innerMainAdapterTester) Load(target *jwt.AdapterRegistration) error {

	target.Adapter = inst
	target.Loader = inst
	target.Provider = inst

	return nil
}

// GetToken implements jwt.Adapter.
func (inst *innerMainAdapterTester) GetToken(acc *jwt.Access) error {
	panic("unimplemented")
}

// SetToken implements jwt.Adapter.
func (inst *innerMainAdapterTester) SetToken(acc *jwt.Access) error {
	panic("unimplemented")
}

// Registration implements jwt.AdapterProvider.
func (inst *innerMainAdapterTester) Registration() *jwt.AdapterRegistration {

	name := fmt.Sprintf("testing-jwt-adapter-%d", inst.priority)
	reg := &jwt.AdapterRegistration{
		Priority: inst.priority,
		Name:     name,
		Enabled:  true,
		Loader:   inst,
		Provider: inst,
	}
	return reg
}

func (inst *innerMainAdapterTester) makeProvider(priority int) jwt.AdapterProvider {

	pro := new(innerMainAdapterTester)
	pro.priority = priority
	return pro
}

func (inst *innerMainAdapterTester) makeProviderList() []jwt.AdapterProvider {

	list := make([]jwt.AdapterProvider, 0)

	list = append(list, inst.makeProvider(11))
	list = append(list, inst.makeProvider(5))
	list = append(list, inst.makeProvider(2))
	list = append(list, inst.makeProvider(3))
	list = append(list, inst.makeProvider(7))
	list = append(list, inst.makeProvider(13))

	return list
}

func (inst *innerMainAdapterTester) _impl(p int) (jwt.AdapterProvider, jwt.AdapterLoader, jwt.Adapter) {
	return inst, inst, inst
}

func (inst *innerMainAdapterTester) test(t *testing.T) {

	m := new(MainAdapter)
	m.ProviderList = inst.makeProviderList()
	src := m.ProviderList

	t.Log("raw:")

	for index, provider := range src {
		reg := provider.Registration()
		t.Logf("[reg index:%d  priority:%d  name:'%s' ]", index, reg.Priority, reg.Name)
	}

	cache, err := m.innerGetCache()
	if err != nil {
		t.Error(err)
		return
	}
	items := cache.items

	t.Log("result:")

	for index, it := range items {
		t.Logf("[reg index:%d  priority:%d  name:'%s' ]", index, it.Priority, it.Name)
	}

}

////////////////////////////////////////////////////////////////////////////////
// EOF
