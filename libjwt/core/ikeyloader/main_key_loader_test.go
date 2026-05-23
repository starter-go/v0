package ikeyloader

import (
	"fmt"
	"testing"

	"github.com/starter-go/v0/libjwt/api/jwt"
)

////////////////////////////////////////////////////////////////////////////////

func TestMainKeyLoaderSort(t *testing.T) {

	tt := new(innerMainKeyLoaderTester)
	tt.test(t)
}

////////////////////////////////////////////////////////////////////////////////

type innerMainKeyLoaderTester struct {
	priority int
}

// Load implements jwt.KeyLoader.
func (inst *innerMainKeyLoaderTester) Load() (*jwt.Key, error) {
	panic("unimplemented")
}

// Registration implements jwt.AdapterProvider.
func (inst *innerMainKeyLoaderTester) Registration() *jwt.KeyLoaderRegistration {

	name := fmt.Sprintf("testing-jwt-key-loader-%d", inst.priority)
	reg := &jwt.KeyLoaderRegistration{
		Priority: inst.priority,
		Name:     name,
		Enabled:  true,
		Loader:   inst,
		Provider: inst,
	}
	return reg
}

func (inst *innerMainKeyLoaderTester) makeProvider(priority int) jwt.KeyLoaderProvider {

	pro := new(innerMainKeyLoaderTester)
	pro.priority = priority
	return pro
}

func (inst *innerMainKeyLoaderTester) makeProviderList() []jwt.KeyLoaderProvider {

	list := make([]jwt.KeyLoaderProvider, 0)

	list = append(list, inst.makeProvider(11))
	list = append(list, inst.makeProvider(5))
	list = append(list, inst.makeProvider(2))
	list = append(list, inst.makeProvider(3))
	list = append(list, inst.makeProvider(7))
	list = append(list, inst.makeProvider(13))

	return list
}

func (inst *innerMainKeyLoaderTester) _impl(p int) (jwt.KeyLoader, jwt.KeyLoaderProvider) {
	return inst, inst
}

func (inst *innerMainKeyLoaderTester) test(t *testing.T) {

	m := new(MainKeyLoader)
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
