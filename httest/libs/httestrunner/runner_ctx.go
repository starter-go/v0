package httestrunner

import (
	"github.com/starter-go/application"
	"github.com/starter-go/v0/httest/libs/httestcore"
)

const theContextBindingName = "ht-test-context"

type Context struct {
	Args []string

	Module application.Module

	Properties map[string]string

	RequestBuilder *httestcore.RequestBuilder
}
