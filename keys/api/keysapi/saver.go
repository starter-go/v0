package keysapi

type Saver interface {
	Save(key Entity) (*PEM, error)
}
