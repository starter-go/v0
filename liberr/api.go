package liberr

import (
	"github.com/starter-go/v0/liberr/api"
)

////////////////////////////////////////////////////////////////////////////////
// types

type Code = api.Code

type Name = api.Name

type Namespace = api.Namespace

type Registration = api.Registration

type ErrorSet = api.ErrorSet

type Formatter = api.Formatter

type ErrorManager = api.ErrorManager

////////////////////////////////////////////////////////////////////////////////
// functions

func SetDefaultFormatter(f Formatter) {
	api.SetDefaultFormatter(f)
}

func DefaultFormatter() Formatter {
	return api.DefaultFormatter()
}

func Format(parent error, reg *Registration, args ...any) error {
	return api.Format(parent, reg, args...)
}

////////////////////////////////////////////////////////////////////////////////
// EOF
