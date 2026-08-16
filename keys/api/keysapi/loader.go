package keysapi

type Loader interface {
	Load(p *PEM) (Entity, error)
}
