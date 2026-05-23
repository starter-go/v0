package jwt

import "bytes"

type Key struct {
	secret []byte
}

type KeyLoader interface {
	Load() (*Key, error)
}

type KeyLoaderProvider interface {
	Registration() *KeyLoaderRegistration
}

type KeyLoaderRegistration struct {
	Name     string
	Enabled  bool
	Priority int
	Loader   KeyLoader
	Provider KeyLoaderProvider
}

////////////////////////////////////////////////////////////////////////////////

func NewKey(data []byte) *Key {
	if data == nil {
		data = []byte{}
	}
	k := new(Key)
	k.secret = data
	return k
}

func (k *Key) Bytes() []byte {
	src := k.secret
	size := len(src)
	if size < 1 {
		return []byte{}
	}
	return bytes.Clone(src)
}

////////////////////////////////////////////////////////////////////////////////
