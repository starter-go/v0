package ikeyloader

import (
	"encoding/base64"
	"encoding/hex"
	"strings"

	"github.com/starter-go/v0/libjwt/api/jwt"
)

type PropertyKeyLoader struct {

	//starter:component

	_as func(jwt.KeyLoaderProvider) //starter:as(".")

	MyEnabled    bool   //starter:inject("${jwt-key-loader.property.enabled}")
	MyName       string //starter:inject("${jwt-key-loader.property.name}")
	MyPriority   int    //starter:inject("${jwt-key-loader.property.priority}")
	MySecretData string //starter:inject("${jwt-key-loader.property.secret}")

	cache []byte
}

// Load implements jwt.KeyLoader.
func (inst *PropertyKeyLoader) Load() (*jwt.Key, error) {

	data, err := inst.innerGetData()
	if err != nil {
		return nil, err
	}
	k := jwt.NewKey(data)
	return k, nil
}

func (inst *PropertyKeyLoader) innerGetData() ([]byte, error) {
	data := inst.cache
	if data == nil {
		raw, err := inst.innerLoadData()
		if err != nil {
			return nil, err
		}
		data = raw
		inst.cache = raw
	}
	return data, nil
}

func (inst *PropertyKeyLoader) innerLoadData() ([]byte, error) {
	const (
		prefixHex = "hex:"
		prefixB64 = "base64:"
	)
	str := inst.MySecretData
	if strings.HasPrefix(str, prefixHex) {
		return inst.innerLoadDataHex(str)

	} else if strings.HasPrefix(str, prefixB64) {
		return inst.innerLoadDataBase64(str)
	}
	return inst.innerLoadDataStr(str)
}

func (inst *PropertyKeyLoader) innerGetDataString(uri string) string {
	const (
		n   = 2
		sep = ":"
	)
	list := strings.SplitN(uri, sep, n)
	if len(list) == n {
		str1 := list[1]
		return strings.TrimSpace(str1)
	}
	return ""
}

func (inst *PropertyKeyLoader) innerLoadDataBase64(uri string) ([]byte, error) {
	str := inst.innerGetDataString(uri)
	return base64.RawStdEncoding.DecodeString(str)
}

func (inst *PropertyKeyLoader) innerLoadDataHex(uri string) ([]byte, error) {
	str := inst.innerGetDataString(uri)
	return hex.DecodeString(str)
}

func (inst *PropertyKeyLoader) innerLoadDataStr(uri string) ([]byte, error) {
	bin := []byte(uri)
	return bin, nil
}

// Registration implements jwt.KeyLoader.
func (inst *PropertyKeyLoader) Registration() *jwt.KeyLoaderRegistration {
	return &jwt.KeyLoaderRegistration{
		Enabled:  inst.MyEnabled,
		Name:     inst.MyName,
		Priority: inst.MyPriority,
		Loader:   inst,
	}
}

func (inst *PropertyKeyLoader) _impl() (jwt.KeyLoader, jwt.KeyLoaderProvider) {
	return inst, inst
}
