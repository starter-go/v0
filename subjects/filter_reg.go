package subjects

type FilterPriority int

const (
	FilterPriorityMin FilterPriority = iota

	FilterPriorityDB
	FilterPriorityCache2
	FilterPriorityJWT
	FilterPriorityCache1
	FilterPriorityLog

	FilterPriorityMax
)

////////////////////////////////////////////////////////////////////////////////

type FilterRegistration struct {
	Name string

	Enabled bool

	Priority FilterPriority

	Writer WriteFilter

	Reader ReadFilter
}

type FilterRegistry interface {
	GetRegistrationList() []*FilterRegistration
}
