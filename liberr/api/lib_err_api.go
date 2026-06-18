package api

type Code int

type Name string

type Namespace string

type Registration struct {
	CMajor Code
	CMinor Code
	Code   Code

	Name Name

	Namespace Namespace

	Format string

	Args string // like: "(a int, b string, c float, ...)"

	Formatter Formatter

	Example error
}

type ErrorSet interface {
	ListErrors() []*Registration
}

type Formatter interface {
	Format(parent error, reg *Registration, args ...any) error
}
