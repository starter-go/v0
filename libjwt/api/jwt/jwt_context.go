package jwt

import "context"

type Access struct {
	Context context.Context

	Token *Token

	Text Text

	Service Service
}
