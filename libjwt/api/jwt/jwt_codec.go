package jwt

type Encoder interface {
	Encode(token *Token) (Text, error)
}

type Decoder interface {
	Decode(txt Text) (*Token, error)
}

type CODEC interface {
	Encoder
	Decoder
}

type CodecRegistration struct {
	Name string

	Enabled bool

	Priority int

	Codec    CODEC
	Loader   CodecLoader
	Provider CodecProvider
}

type CodecProvider interface {
	Registration() *CodecRegistration
}

// CodecLoader  负责加载 CODEC 实例
type CodecLoader interface {
	Load(target *CodecRegistration) error
}
