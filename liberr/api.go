package liberr

import "github.com/starter-go/v0/liberr/api"

////////////////////////////////////////////////////////////////////////////////
// types

type Code = api.Code

type Name = api.Name

type Namespace = api.Namespace

type Registration = api.Registration

type ErrorSet = api.ErrorSet

type ErrorSetHolder = api.ErrorSetHolder

type ErrorSetLoader = api.ErrorSetLoader

type Formatter = api.Formatter

type ErrorManager = api.ErrorManager

type Service = api.Service

// hyper-error

type HyperError = api.HyperError

type HyperErrorInfo = api.HyperErrorInfo

////////////////////////////////////////////////////////////////////////////////
// functions

func SetDefaultFormatter(f Formatter) {
	api.SetDefaultFormatter(f)
}

func DefaultFormatter() Formatter {
	return api.DefaultFormatter()
}

func NewHyperError(info *HyperErrorInfo) HyperError {
	return api.NewHyperError(info)
}

////////////////////////////////////////////////////////////////////////////////
// EOF
