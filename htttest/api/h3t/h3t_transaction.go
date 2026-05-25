package h3t

import (
	"context"
	"net/http"
	"net/url"
)

type Want struct {

	// request

	Request *http.Request

	Method string

	RawURL string

	URL string

	Location *url.URL

	Params map[string]string

	Query map[string]string

	Head map[string]string

	Body *Content
}

type Have struct {

	// response

	Response *http.Response

	Status int

	Message string

	Error error

	Head map[string]string

	Body *Content
}

type Transaction struct {

	// contexts

	CC context.Context

	AC *AgentContext

	Service Service

	CloserPool CloserPool

	// message

	Want Want

	Have Have
}
