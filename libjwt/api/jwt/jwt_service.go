package jwt

type Service interface {
	GetAdapter() Adapter

	GetDecoder() Decoder

	GetEncoder() Encoder
}
