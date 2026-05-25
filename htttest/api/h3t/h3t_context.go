package h3t

import (
	"net/http"

	"github.com/starter-go/application/attributes"
	"github.com/starter-go/application/properties"
)

type AgentContext struct {
	BaseURL string

	MaxContentLength int64

	Service Service

	Resolver Resolver

	UserAgent UserAgent

	Client *http.Client

	Attributes attributes.Table

	Properties properties.Table

	RequestHeaders properties.Table
}
