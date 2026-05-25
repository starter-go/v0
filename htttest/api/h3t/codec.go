package h3t

type Encoder interface {
	Encode(c *Content) error
}

type Decoder interface {
	Decode(c *Content) error
}

type Codec interface {
	Encoder
	Decoder
}

type CodecProvider interface {
	Registrations() []*CodecRegistration
}

type CodecRegistration struct {
	Type     ContentType
	Name     string
	Enabled  bool
	Priority int
	Encoder  Encoder
	Decoder  Decoder
	Codec    Codec
}

type CodecManager interface {
	GetDefaultCodec() Codec

	FindCodec(ctype ContentType) (Codec, error)

	FindEncoder(ctype ContentType) (Encoder, error)

	FindDecoder(ctype ContentType) (Decoder, error)
}
