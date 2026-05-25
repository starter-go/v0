package h3t

type Content struct {
	Type ContentType

	Length int64

	Decoder Decoder

	Encoder Encoder

	Encoded []byte

	Data any
}
