package icodec

import (
	"fmt"
	"testing"

	"github.com/starter-go/v0/libjwt/api/jwt"
)

////////////////////////////////////////////////////////////////////////////////

func TestMainCodecSort(t *testing.T) {

	tt := new(innerMainCodecTester)
	tt.test(t)
}

////////////////////////////////////////////////////////////////////////////////

type innerMainCodecTester struct {
	priority int
}

// Registration implements jwt.CodecProvider.
func (inst *innerMainCodecTester) Registration() *jwt.CodecRegistration {

	name := fmt.Sprintf("testing-jwt-codec-%d", inst.priority)
	reg := &jwt.CodecRegistration{
		Priority: inst.priority,
		Name:     name,
		Enabled:  true,
		Loader:   inst,
		Provider: inst,
	}
	return reg
}

// Load implements jwt.CodecLoader.
func (inst *innerMainCodecTester) Load(target *jwt.CodecRegistration) error {

	target.Codec = inst
	target.Loader = inst
	target.Provider = inst

	return nil
}

// Decode implements jwt.CODEC.
func (inst *innerMainCodecTester) Decode(txt jwt.Text) (*jwt.Token, error) {
	panic("unimplemented")
}

// Encode implements jwt.CODEC.
func (inst *innerMainCodecTester) Encode(token *jwt.Token) (jwt.Text, error) {
	panic("unimplemented")
}

func (inst *innerMainCodecTester) makeProvider(priority int) jwt.CodecProvider {

	pro := new(innerMainCodecTester)
	pro.priority = priority
	return pro
}

func (inst *innerMainCodecTester) makeProviderList() []jwt.CodecProvider {

	list := make([]jwt.CodecProvider, 0)

	list = append(list, inst.makeProvider(11))
	list = append(list, inst.makeProvider(5))
	list = append(list, inst.makeProvider(2))
	list = append(list, inst.makeProvider(3))
	list = append(list, inst.makeProvider(7))
	list = append(list, inst.makeProvider(13))

	return list
}

func (inst *innerMainCodecTester) _impl(p int) (jwt.CodecProvider, jwt.CodecLoader, jwt.CODEC) {
	return inst, inst, inst
}

func (inst *innerMainCodecTester) test(t *testing.T) {

	m := new(MainCodec)
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
