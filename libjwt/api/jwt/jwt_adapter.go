package jwt

type Adapter interface {
	SetToken(acc *Access) error

	GetToken(acc *Access) error
}

type AdapterRegistration struct {
	Name string

	Enabled bool

	Priority int

	Adapter  Adapter
	Loader   AdapterLoader
	Provider AdapterProvider
}

type AdapterProvider interface {
	Registration() *AdapterRegistration
}

// AdapterLoader  负责加载 Adapter 实例
type AdapterLoader interface {
	Load(target *AdapterRegistration) error
}
